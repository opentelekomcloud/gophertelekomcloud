package v1

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	nwv1 "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/networking/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPoliciesList(t *testing.T) {
	if os.Getenv("RUN_CSBS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewCsbsV1Client()
	if err != nil {
		t.Fatalf("Unable to create a CSBSv1 client: %s", err)
	}

	listOpts := policies.ListOpts{}
	backupPolicies, err := policies.List(client, listOpts)
	if err != nil {
		t.Fatalf("Unable fetch CSBSv1 policies pages: %s", err)
	}

	for _, policy := range backupPolicies {
		tools.PrintResource(t, policy)
	}
}

func TestPoliciesLifeCycle(t *testing.T) {
	if os.Getenv("RUN_CSBS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewCsbsV1Client()
	if err != nil {
		t.Fatalf("Unable to create CSBSv1 client: %s", err)
	}

	subnet := nwv1.CreateNetwork(t, prefix, az)
	defer nwv1.DeleteNetwork(t, subnet)
	server := createComputeInstance(t, subnet.ID)

	defer deleteComputeInstance(t, server.ID)

	// Create CSBSv1 policy
	policy, err := createCSBSPolicy(t, client, server.ID)
	th.AssertNoErr(t, err)
	defer deleteCSBSPolicy(t, client, policy.ID)
	tools.PrintResource(t, policy)

	err = updateCSBSPolicy(client, policy.ID)
	th.AssertNoErr(t, err)

	policyUpdate, err := policies.Get(client, policy.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policyUpdate)
}

const (
	image  = "Standard_Debian_10_latest"
	flavor = "s2.large.2"
	az     = "eu-de-01"
	prefix = "csbs-acc-"
)

func createCSBSPolicy(t *testing.T, client *golangsdk.ServiceClient, serverId string) (*policies.BackupPolicy, error) {
	if os.Getenv("RUN_CSBS") == "" {
		t.Skip("unstable test")
	}
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

	policy, err := policies.Create(client, createOpts)
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
	if os.Getenv("RUN_CSBS") == "" {
		t.Skip("unstable test")
	}
	t.Logf("Attempting to delete CSBSv1: %s", policyId)

	err := policies.Delete(client, policyId)
	if err != nil {
		t.Fatalf("Unable to delete CSBSv1 policy: %s", err)
	}
	err = waitForCSBSPolicyDelete(client, 600, policyId)
	if err != nil {
		t.Fatalf("Wait for CSBSv1 policy delete fails: %s", err)
	}

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

func createComputeInstance(t *testing.T, subnetID string) *servers.Server {
	if os.Getenv("RUN_CSBS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	serverName := tools.RandomString(prefix, 5)
	opts := servers.CreateOpts{
		Name:             serverName,
		SecurityGroups:   []string{"default"},
		FlavorName:       flavor,
		ImageName:        image,
		AvailabilityZone: az,
		ServiceClient:    client,
		Networks: []servers.Network{
			{
				UUID: subnetID,
			},
		},
	}

	server, err := servers.Create(client, opts)
	th.AssertNoErr(t, err)
	err = waitForComputeInstanceAvailable(client, 600, server.ID)
	th.AssertNoErr(t, err)

	return server
}

func deleteComputeInstance(t *testing.T, instanceId string) {
	if os.Getenv("RUN_CSBS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	err = servers.Delete(client, instanceId)
	th.AssertNoErr(t, err)
	err = waitForComputeInstanceDelete(client, 600, instanceId)
	th.AssertNoErr(t, err)
}

func waitForComputeInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		server, err := servers.Get(client, instanceId)
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
		server, err := servers.Get(client, instanceId)
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
