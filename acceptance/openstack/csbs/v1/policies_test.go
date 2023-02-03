package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPoliciesList(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	th.AssertNoErr(t, err)

	listOpts := policies.ListOpts{}
	backupPolicies, err := policies.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, policy := range backupPolicies {
		tools.PrintResource(t, policy)
	}
}

func TestPoliciesLifeCycle(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	th.AssertNoErr(t, err)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	t.Cleanup(func() { openstack.DeleteCloudServer(t, computeClient, ecs.ID) })

	// Create CSBSv1 policy
	policy := createCSBSPolicy(t, client, ecs.ID)
	tools.PrintResource(t, policy)

	err = updateCSBSPolicy(client, policy.ID)
	th.AssertNoErr(t, err)

	policyUpdate, err := policies.Get(client, policy.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policyUpdate)
}

func createCSBSPolicy(t *testing.T, client *golangsdk.ServiceClient, serverId string) *policies.BackupPolicy {
	t.Logf("Attempting to create CSBSv1 policy")

	createOpts := policies.CreateOpts{
		Description: tools.RandomString("description-init-", 10),
		Name:        tools.RandomString("policy-init-", 5),
		Parameters: policies.PolicyParam{
			Common: map[string]string{},
		},
		Resources: []policies.Resource{
			{
				Id:   serverId,
				Type: "OS::Nova::Server",
				Name: "resource1",
			},
		},
		ProviderId: "fc4d5750-22e7-4798-8a46-f48f62c4c1da",
		ScheduledOperations: []policies.ScheduledOperation{
			{
				Enabled:       false,
				OperationType: "backup",
				OperationDefinition: policies.OperationDefinition{
					MaxBackups: pointerto.Int(2),
				},
				Trigger: policies.Trigger{
					Properties: policies.TriggerProperties{
						Pattern: "BEGIN:VCALENDAR\\r\\nBEGIN:VEVENT\\r\\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\\r\\nEND:VEVENT\\r\\nEND:VCALENDAR\\r\\n",
					},
				},
			},
		},
	}

	policy, err := policies.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { deleteCSBSPolicy(t, client, policy.ID) })

	err = waitForCSBSPolicyActive(client, 600, policy.ID)
	th.AssertNoErr(t, err)

	t.Logf("Created CSBSv1 Policy: %s", policy.ID)
	return policy
}

func deleteCSBSPolicy(t *testing.T, client *golangsdk.ServiceClient, policyId string) {
	t.Logf("Attempting to delete CSBSv1: %s", policyId)

	err := policies.Delete(client, policyId)
	th.AssertNoErr(t, err)

	err = waitForCSBSPolicyDelete(client, 600, policyId)
	th.AssertNoErr(t, err)

	t.Logf("Deleted CSBSv1 Policy: %s", policyId)
}

func updateCSBSPolicy(client *golangsdk.ServiceClient, policyId string) error {
	policyNameUpdate := tools.RandomString("policy-update-", 5)
	policyDescriptionUpdate := tools.RandomString("description-update-", 10)
	updateOpts := policies.UpdateOpts{
		Description: policyDescriptionUpdate,
		Name:        policyNameUpdate,
	}

	_, err := policies.Update(client, policyId, updateOpts)
	if err != nil {
		return err
	}
	return nil
}

func waitForCSBSPolicyActive(client *golangsdk.ServiceClient, secs int, policyId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		policy, err := policies.Get(client, policyId)
		if err != nil {
			return false, err
		}

		if policy.Status == "suspended" {
			return true, nil
		}
		return false, nil
	})
}

func waitForCSBSPolicyDelete(client *golangsdk.ServiceClient, secs int, policyId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := policies.Get(client, policyId)
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			return true, nil
		}
		return false, nil
	})
}
