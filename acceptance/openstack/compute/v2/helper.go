package v2

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/flavors"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateServer(t *testing.T, client *golangsdk.ServiceClient) (*servers.Server, error) {
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

	// TODO: API discarded and not working.
	imageID, err := images.IDFromName(client, "Standard_Debian_10_latest")
	th.AssertNoErr(t, err)

	flavorID, err := flavors.IDFromName(client, "s2.large.2")
	th.AssertNoErr(t, err)

	ecs, err := servers.Create(client, servers.CreateOpts{
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
	}).Extract()
	th.AssertNoErr(t, err)

	err = servers.WaitForStatus(client, ecs.ID, "ACTIVE", 1200)
	th.AssertNoErr(t, err)
	t.Logf("Created ECSv2: %s", ecs.ID)
	return ecs, err
}

func DeleteServer(t *testing.T, client *golangsdk.ServiceClient, ecs *servers.Server) {
	t.Logf("Attempting to delete ECSv2: %s", ecs.ID)

	_, err := servers.Delete(client, ecs.ID).ExtractJobResponse()
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(1200, func() (bool, error) {
		_, err := servers.Get(client, ecs.ID).Extract()
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
}
