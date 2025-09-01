package v3

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/sqlfilter"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTaurusSqlFilterLifecycle(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	createResp := createTaurusInstance(t, client, vpcID, subnetID, secGroupID)
	instanceID := createResp.Instance.Id

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db")
		_, err = instance.Delete(client, instanceID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Waiting for instance to become available")
	err = waitForInstanceAvailable(client, 1200, instanceID)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create taurus db read replica")

	createReplicaJobs, err := instance.CreateReplica(client, createResp.Instance.Id,
		instance.CreateReplicaOpts{
			Priorities: []int{1, 2},
		},
	)
	th.AssertNoErr(t, err)

	replicaJobs := strings.Split(*createReplicaJobs, ",")

	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, replicaJobs[0]))
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, replicaJobs[1]))

	getResp, err := instance.Get(client, createResp.Instance.Id)
	th.AssertNoErr(t, err)

	nodeId := getResp.Nodes[0].Id

	t.Logf("Testing GetSqlFilterSwitch - should be OFF by default")
	switchResponse, err := sqlfilter.GetSqlFilterSwitch(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, switchResponse.SwitchStatus, "OFF")

	t.Logf("Testing SwitchSqlFilter - enabling")
	enableSqlFilterOpts := sqlfilter.SqlFilterSwitchOpts{
		SwitchStatus: "ON",
		InstanceId:   instanceID,
	}
	jobID, err := sqlfilter.SwitchSqlFilter(client, enableSqlFilterOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for enable sql filter job to complete")
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 300, *jobID))

	t.Cleanup(func() {
		t.Logf("Attempting to disable sql filter")
		disableOpts := sqlfilter.SqlFilterSwitchOpts{
			SwitchStatus: "OFF",
			InstanceId:   instanceID,
		}
		jobID, err = sqlfilter.SwitchSqlFilter(client, disableOpts)
		th.AssertNoErr(t, err)

		t.Logf("Waiting for disable sql filter job to complete")
		th.AssertNoErr(t, job.WaitForJobSuccess(client, 300, *jobID))
	})

	t.Logf("Testing GetSqlFilterSwitch - should be ON now")
	switchResponse, err = sqlfilter.GetSqlFilterSwitch(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, switchResponse.SwitchStatus, "ON")

	t.Logf("Testing UpdateSqlFilterRules")
	updateRulesOpts := sqlfilter.UpdateSqlFilterRulesOpts{
		SqlFilterRules: []sqlfilter.NodeSqlFilterRuleInfo{
			{
				NodeId: nodeId,
				Rules: []sqlfilter.NodeSqlFilterRule{
					{
						SqlType: "SELECT",
						Patterns: []sqlfilter.NodeSqlFilterRulePattern{
							{
								Pattern:        "select~from~test_table",
								MaxConcurrency: 10,
							},
							{
								Pattern:        "select~from~users~where~id",
								MaxConcurrency: 5,
							},
						},
					},
					{
						SqlType: "UPDATE",
						Patterns: []sqlfilter.NodeSqlFilterRulePattern{
							{
								Pattern:        "update~test_table~set",
								MaxConcurrency: 3,
							},
						},
					},
				},
			},
		},
	}
	jobID, err = sqlfilter.UpdateSqlFilterRules(client, instanceID, updateRulesOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for update sql filter rules job to complete")
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 300, *jobID))

	t.Logf("Testing GetSqlFilterRules - all rules")
	getRulesOpts := sqlfilter.GetSqlFilterRulesOpts{
		NodeId: nodeId,
	}
	rulesResponse, err := sqlfilter.GetSqlFilterRules(client, instanceID, getRulesOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, rulesResponse.NodeId, nodeId)
	th.AssertEquals(t, len(rulesResponse.SqlFilterRules) > 0, true)

	t.Logf("Testing GetSqlFilterRules - SELECT only")
	getRulesOptsSelect := sqlfilter.GetSqlFilterRulesOpts{
		NodeId: nodeId,
		Type:   "SELECT",
	}
	selectRulesResponse, err := sqlfilter.GetSqlFilterRules(client, instanceID, getRulesOptsSelect)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, selectRulesResponse.NodeId, nodeId)

	hasSelectRule := false
	for _, rule := range selectRulesResponse.SqlFilterRules {
		if rule.SqlType == "SELECT" {
			hasSelectRule = true
			th.AssertEquals(t, len(rule.Patterns) > 0, true)
			break
		}
	}
	th.AssertEquals(t, hasSelectRule, true)

	t.Logf("Testing DeleteSqlFilterRules")
	deleteRulesOpts := sqlfilter.DeleteSqlFilterRulesOpts{
		SqlFilterRules: []sqlfilter.DeleteNodeSqlFilterRuleInfo{
			{
				NodeId: nodeId,
				Rules: []sqlfilter.DeleteNodeSqlFilterRule{
					{
						SqlType:  "SELECT",
						Patterns: []string{"select~from~test_table"},
					},
					{
						SqlType:  "UPDATE",
						Patterns: []string{"update~test_table~set"},
					},
				},
			},
		},
	}
	jobID, err = sqlfilter.DeleteSqlFilterRules(client, instanceID, deleteRulesOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for delete sql filter rules job to complete")
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 300, *jobID))

	t.Logf("Verifying rules were deleted")
	finalRulesResponse, err := sqlfilter.GetSqlFilterRules(client, instanceID, getRulesOpts)
	th.AssertNoErr(t, err)

	for _, rule := range finalRulesResponse.SqlFilterRules {
		for _, pattern := range rule.Patterns {
			th.AssertEquals(t, pattern.Pattern != "select~from~test_table", true)
			th.AssertEquals(t, pattern.Pattern != "update~test_table~set", true)
		}
	}
}
