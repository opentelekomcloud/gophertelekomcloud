package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/policies"
)

func TestPoliciesList(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	if err != nil {
		t.Fatalf("Unable to create a CSBSv1 client: %s", err)
	}

	listOpts := policies.ListOpts{}
	backupPolicies, err := policies.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable fetch CSBSv1 policies pages: %s", err)
	}

	policiesExtracted, err := policies.ExtractBackupPolicies(backupPolicies)
	if err != nil {
		t.Fatalf("Unable to extract policies: %s", err)
	}

	for _, policy := range policiesExtracted {
		tools.PrintResource(t, policy)
	}
}

func TestPoliciesLifeCycle(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	if err != nil {
		t.Fatalf("Unable to create CSBSv1 client: %s", err)
	}

	computeClient, err := clients.NewComputeV2Client()
	if err != nil {
		t.Fatalf("Unable to create ComputeV2 client: %s", err)
	}

	server, err := createComputeInstance(computeClient)
	if err != nil {
		t.Fatalf("Error creating compute instance: %s", err)
	}

	defer func() {
		err = deleteComputeInstance(computeClient, server.ID)
		if err != nil {
			t.Fatalf("Error deleting compute instance: %s", err)
		}
	}()

	// Create CSBSv1 policy
	policy, err := createCSBSPolicy(t, client, server.ID)
	if err != nil {
		t.Fatalf("Unable to create CSBSv1 policy: %s", err)
	}
	defer deleteCSBSPolicy(t, client, policy.ID)
	tools.PrintResource(t, policy)

	err = updateCSBSPolicy(t, client, policy.ID)
	if err != nil {
		t.Fatalf("Unable to update CSBSv1 policy: %s", err)
	}

	policyUpdate, err := policies.Get(client, policy.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get updated CSBSv1 policy: %s", err)
	}
	tools.PrintResource(t, policyUpdate)
}

func createCSBSPolicy(t *testing.T, client *golangsdk.ServiceClient, serverId string) (*policies.CreateBackupPolicy, error) {
	t.Logf("Attempting to create CSBSv1 policy")

	// These values were got from HelpCenter request example
	triggerPattern := "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"

	providerId := "fc4d5750-22e7-4798-8a46-f48f62c4c1da"

	policyName := tools.RandomString("policy-init-", 5)
	policyDescription := tools.RandomString("description-init-", 10)

	createOpts := policies.CreateOpts{
		Description: policyDescription,
		Name:        policyName,
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
		ProviderId: providerId,
		ScheduledOperations: []policies.ScheduledOperation{
			{
				Enabled:             false,
				OperationType:       "backup",
				OperationDefinition: policies.OperationDefinition{},
				Trigger: policies.Trigger{
					Properties: policies.TriggerProperties{
						Pattern: triggerPattern,
					},
				},
			},
		},
	}

	policy, err := policies.Create(client, createOpts).Extract()
	if err != nil {
		return nil, err
	}
	err = waitForCSBSPolicyActive(client, 600, policy.ID)
	if err != nil {
		return nil, err
	}
	t.Logf("Created CSBSv1 Policy: %s", policy.ID)

	return policy, nil
}

func deleteCSBSPolicy(t *testing.T, client *golangsdk.ServiceClient, policyId string) {
	t.Logf("Attempting to delete CSBSv1: %s", policyId)

	err := policies.Delete(client, policyId).Err
	if err != nil {
		t.Fatalf("Unable to delete CSBSv1 policy: %s", err)
	}
	err = waitForCSBSPolicyDelete(client, 600, policyId)
	if err != nil {
		t.Fatalf("Wait for CSBSv1 policy delete fails: %s", err)
	}

	t.Logf("Deleted CSBSv1 Policy: %s", policyId)
}

func updateCSBSPolicy(t *testing.T, client *golangsdk.ServiceClient, policyId string) error {
	policyNameUpdate := tools.RandomString("policy-update-", 5)
	policyDescriptionUpdate := tools.RandomString("description-update-", 10)
	updateOpts := policies.UpdateOpts{
		Description: policyDescriptionUpdate,
		Name:        policyNameUpdate,
		Parameters: policies.PolicyParam{
			Common: map[string]string{},
		},
	}

	err := policies.Update(client, policyId, updateOpts).Err
	if err != nil {
		return err
	}
	return nil
}

func waitForCSBSPolicyActive(client *golangsdk.ServiceClient, secs int, policyId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		policy, err := policies.Get(client, policyId).Extract()
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
		_, err := policies.Get(client, policyId).Extract()
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			return true, nil
		}
		return false, nil
	})
}

func createComputeInstance(client *golangsdk.ServiceClient) (*servers.Server, error) {
	computeName := tools.RandomString("csbs-acc-", 5)
	createOpts := servers.CreateOpts{
		Name:             computeName,
		SecurityGroups:   []string{"default"},
		FlavorName:       clients.OS_FLAVOR_NAME,
		ImageRef:         clients.OS_IMAGE_ID,
		AvailabilityZone: clients.OS_AVAILABILITY_ZONE,
		ServiceClient:    client,
		Networks: []servers.Network{
			{
				UUID: clients.OS_NETWORK_ID,
			},
		},
	}

	server, err := servers.Create(client, createOpts).Extract()
	if err != nil {
		return nil, err
	}
	err = waitForComputeInstanceAvailable(client, 600, server.ID)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func deleteComputeInstance(client *golangsdk.ServiceClient, instanceId string) error {
	err := servers.Delete(client, instanceId).ExtractErr()
	if err != nil {
		return err
	}
	err = waitForComputeInstanceDelete(client, 600, instanceId)
	if err != nil {
		return err
	}
	return nil
}

func waitForComputeInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		server, err := servers.Get(client, instanceId).Extract()
		if err != nil {
			return false, err
		}
		if server.Status == "ACTIVE" {
			return true, nil
		}
		return false, nil
	})
}

func waitForComputeInstanceDelete(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		server, err := servers.Get(client, instanceId).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		if server.Status == "ERROR" {
			return false, err
		}
		return false, nil
	})
}
