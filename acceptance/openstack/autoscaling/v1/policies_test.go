package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyLifecycle(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if networkID == "" || vpcID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but AS Policy test requires")
	}

	groupID := autoscaling.CreateAutoScalingGroup(t, client, networkID, vpcID, tools.RandomString("as-group-create-", 3))
	t.Cleanup(func() {
		autoscaling.DeleteAutoScalingGroup(t, client, groupID)
	})

	asPolicyCreateName := tools.RandomString("as-policy-create-", 3)
	t.Logf("Attempting to create AutoScaling Policy")
	policyID, err := policies.Create(client, policies.CreateOpts{
		Name: asPolicyCreateName,
		Type: "RECURRENCE",
		ID:   groupID,
		SchedulePolicy: policies.SchedulePolicyOpts{
			LaunchTime:      "10:30",
			RecurrenceType:  "Weekly",
			RecurrenceValue: "1,3,5",
			EndTime:         "2040-12-31T10:30Z",
		},
		Action: policies.Action{
			Operation:          "ADD",
			InstancePercentage: 15,
		},
	})
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Policy: %s", policyID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete AutoScaling Policy")
		err := policies.Delete(client, policyID)
		th.AssertNoErr(t, err)
		t.Logf("Deleted AutoScaling Policy: %s", policyID)
	})

	policy, err := policies.Get(client, policyID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, asPolicyCreateName, policy.Name)
	th.AssertEquals(t, 15, policy.Action.InstancePercentage)
	th.AssertEquals(t, "ADD", policy.Action.Operation)

	t.Logf("Attempting to update AutoScaling policy")
	asPolicyUpdateName := tools.RandomString("as-policy-update-", 3)

	updateOpts := policies.UpdateOpts{
		Name: asPolicyUpdateName,
		Type: "RECURRENCE",
		SchedulePolicy: policies.CreateOpts{
			Name: asPolicyCreateName,
			Type: "RECURRENCE",
			ID:   groupID,
			SchedulePolicy: policies.SchedulePolicyOpts{
				LaunchTime:      "10:30",
				RecurrenceType:  "Weekly",
				RecurrenceValue: "1,3,5",
				EndTime:         "2040-12-31T10:30Z",
			},
			Action: policies.Action{
				Operation:          "ADD",
				InstancePercentage: 15,
			},
		}.SchedulePolicy,
		Action: policies.Action{
			InstancePercentage: 30,
		},
		CoolDownTime: 0,
	}

	policyID, err = policies.Update(client, policyID, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated AutoScaling Policy")

	policy, err = policies.Get(client, policyID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policy)
	th.AssertEquals(t, asPolicyUpdateName, policy.Name)
	th.AssertEquals(t, 30, policy.Action.InstancePercentage)
}
