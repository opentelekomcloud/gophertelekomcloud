package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v2/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyLifecycle(t *testing.T) {
	client, err := clients.NewAutoscalingV2Client()
	th.AssertNoErr(t, err)

	asPolicyCreateName := tools.RandomString("as-policy-create-", 3)
	asGroupCreateName := tools.RandomString("as-group-create-", 3)
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if networkID == "" || vpcID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but AS Policy test requires")
	}
	groupID := autoscaling.CreateAutoScalingGroup(t, client, networkID, vpcID, asGroupCreateName)
	defer func() {
		autoscaling.DeleteAutoScalingGroup(t, client, groupID)
	}()

	createOpts := policies.CreateOpts{
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
	policyID, err := policies.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Policy: %s", policyID)
	defer func() {
		t.Logf("Attempting to delete AutoScaling Policy")
		err := policies.Delete(client, policyID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted AutoScaling Policy: %s", policyID)
	}()

	policy, err := policies.Get(client, policyID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, asPolicyCreateName, policy.PolicyName)
	th.AssertEquals(t, 15, policy.PolicyAction.Percentage)
	th.AssertEquals(t, "ADD", policy.PolicyAction.Operation)

	t.Logf("Attempting to update AutoScaling policy")
	asPolicyUpdateName := tools.RandomString("as-policy-update-", 3)

	updateOpts := policies.UpdateOpts{
		PolicyName: asPolicyUpdateName,
		Action: policies.ActionOpts{
			Percentage: 30,
		},
	}

	policyID, err = policies.Update(client, policy.PolicyID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated AutoScaling Policy")

	policy, err = policies.Get(client, policyID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policy)
	th.AssertEquals(t, asPolicyUpdateName, policy.PolicyName)
	th.AssertEquals(t, 30, policy.PolicyAction.Percentage)
}
