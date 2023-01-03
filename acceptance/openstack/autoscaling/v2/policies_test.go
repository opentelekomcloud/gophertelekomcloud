package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v2/logs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v2/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyLifecycle(t *testing.T) {
	v1client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)
	client, err := clients.NewAutoscalingV2Client()
	th.AssertNoErr(t, err)

	asPolicyCreateName := tools.RandomString("as-policy-create-", 3)
	asGroupCreateName := tools.RandomString("as-group-create-", 3)
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if networkID == "" || vpcID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but AS Policy test requires")
	}

	groupID := autoscaling.CreateAutoScalingGroup(t, v1client, networkID, vpcID, asGroupCreateName)
	t.Cleanup(func() {
		autoscaling.DeleteAutoScalingGroup(t, v1client, groupID)
	})

	createOpts := policies.PolicyOpts{
		PolicyName:   asPolicyCreateName,
		PolicyType:   "RECURRENCE",
		ResourceID:   groupID,
		ResourceType: "SCALING_GROUP",
		SchedulePolicy: policies.SchedulePolicyOpts{
			LaunchTime:      "10:30",
			RecurrenceType:  "Weekly",
			RecurrenceValue: "1,3,5",
			EndTime:         "2040-12-31T10:30Z",
		},
		PolicyAction: policies.ActionOpts{
			Operation:  "ADD",
			Percentage: 15,
		},
	}

	t.Logf("Attempting to create AutoScaling Policy")
	policyID, err := policies.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Policy: %s", policyID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete AutoScaling Policy")
		err := policies.Delete(v1client, policyID)
		th.AssertNoErr(t, err)
		t.Logf("Deleted AutoScaling Policy: %s", policyID)
	})

	policy, err := policies.Get(client, policyID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, asPolicyCreateName, policy.PolicyName)
	th.AssertEquals(t, 15, policy.PolicyAction.Percentage)
	th.AssertEquals(t, "ADD", policy.PolicyAction.Operation)

	t.Logf("Attempting to update AutoScaling policy")
	asPolicyUpdateName := tools.RandomString("as-policy-update-", 3)

	updateOpts := policies.PolicyOpts{
		PolicyName:     asPolicyUpdateName,
		PolicyType:     "RECURRENCE",
		ResourceID:     groupID,
		ResourceType:   "SCALING_GROUP",
		SchedulePolicy: createOpts.SchedulePolicy,
		PolicyAction: policies.ActionOpts{
			Percentage: 30,
		},
		CoolDownTime: 0,
	}

	policyID, err = policies.Update(client, policy.PolicyID, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated AutoScaling Policy")

	policy, err = policies.Get(client, policyID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policy)
	th.AssertEquals(t, asPolicyUpdateName, policy.PolicyName)
	th.AssertEquals(t, 30, policy.PolicyAction.Percentage)

	activityLogs, err := logs.ListScalingActivityLogs(client, logs.ListScalingActivityLogsOpts{ScalingGroupId: groupID})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, activityLogs)
}
