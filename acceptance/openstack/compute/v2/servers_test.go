package v2

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/flavors"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServerList(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	listOpts := servers.ListOpts{}
	allServerPages, err := servers.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	serversList, err := servers.ExtractServers(allServerPages)
	th.AssertNoErr(t, err)

	for _, server := range serversList {
		tools.PrintResource(t, server)
	}
}

func TestServerLifecycle(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create ECSv2")
	ecsName := tools.RandomString("create-ecs-", 3)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if networkID == "" {
		t.Skip("OS_NETWORK_ID env var is missing but ECS test requires using existing network")
	}

	imageID, err := images.IDFromName(client, "Standard_Debian_10_latest")
	th.AssertNoErr(t, err)

	flavorID, err := flavors.IDFromName(client, "s2.large.2")
	th.AssertNoErr(t, err)

	createOpts := servers.CreateOpts{
		Name:      ecsName,
		ImageRef:  imageID,
		FlavorRef: flavorID,
		SecurityGroups: []string{
			openstack.DefaultSecurityGroup(t),
		},
		AvailabilityZone: az,
		Networks: []servers.Network{
			{
				UUID: networkID,
			},
		},
	}

	ecs, err := servers.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = servers.WaitForStatus(client, ecs.ID, "ACTIVE", 1200)
	th.AssertNoErr(t, err)
	t.Logf("Created ECSv2: %s", ecs.ID)

	ecs, err = servers.Get(client, ecs.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ecsName, ecs.Name)

	nicInfo, err := servers.GetNICs(client, ecs.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, nicInfo)

	defer func() {
		t.Logf("Attempting to delete ECSv2: %s", ecs.ID)

		err := servers.Delete(client, ecs.ID)
		th.AssertNoErr(t, err)

		err = golangsdk.WaitFor(1200, func() (bool, error) {
			_, err := servers.Get(client, ecs.ID)
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault400); ok {
					time.Sleep(10 * time.Second)
					return false, nil
				}
				return false, err
			}
			return true, nil
		})
		th.AssertNoErr(t, err)

		t.Logf("ECSv2 instance deleted: %s", ecs.ID)
	}()

	t.Logf("Attempting to update ECSv2: %s", ecs.ID)

	ecsName = tools.RandomString("update-ecs-", 3)
	updateOpts := servers.UpdateOpts{
		Name: ecsName,
	}

	_, err = servers.Update(client, ecs.ID, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("ECSv2 successfully updated: %s", ecs.ID)
	th.AssertNoErr(t, err)

	newECS, err := servers.Get(client, ecs.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ecsName, newECS.Name)
}
