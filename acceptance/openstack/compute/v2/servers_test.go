package v2

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/availabilityzones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/diskconfig"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/extendedstatus"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/floatingips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/networks"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/flavors"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/cloudimages"
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

	imsClient, err := clients.NewImageServiceV2Client()
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
	imageName := "Standard_Debian_10_latest"
	listOpts := &cloudimages.ListOpts{
		Name: imageName,
	}
	allPages, err := cloudimages.List(imsClient, listOpts).AllPages()
	th.AssertNoErr(t, err)

	extractImages, err := cloudimages.ExtractImages(allPages)
	th.AssertNoErr(t, err)

	if len(extractImages) < 1 {
		t.Fatal("[ERROR] cannot find the image")
	}

	flavorID, err := flavors.IDFromName(client, "s2.large.2")
	th.AssertNoErr(t, err)

	createOpts := servers.CreateOpts{
		Name:      ecsName,
		ImageRef:  extractImages[0].ID,
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

	ecs, err := servers.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = servers.WaitForStatus(client, ecs.ID, "ACTIVE", 1200)
	th.AssertNoErr(t, err)
	t.Logf("Created ECSv2: %s", ecs.ID)

	opts := servers.ListOpts{
		Name: ecsName,
	}
	allServerPages, err := servers.List(client, opts).AllPages()
	th.AssertNoErr(t, err)
	type ServerWithExt struct {
		servers.Server
		availabilityzones.ServerAvailabilityZoneExt
		extendedstatus.ServerExtendedStatusExt
		diskconfig.ServerDiskConfigExt
		floatingips.FloatingIP
		networks.Network
		secgroups.SecurityGroup
		volumeattach.VolumeAttachment
	}
	var allServers []ServerWithExt
	err = servers.ExtractServersInto(allServerPages, &allServers)
	th.AssertNoErr(t, err)

	if len(allServers) < 1 {
		t.Fatal("[ERROR] cannot find the server")
	}
	th.AssertEquals(t, ecsName, allServers[0].Name)
	th.AssertEquals(t, az, allServers[0].AvailabilityZone)

	ecs, err = servers.Get(client, ecs.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ecsName, ecs.Name)

	nicInfo, err := servers.GetNICs(client, ecs.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, nicInfo)

	defer func() {
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
	}()

	t.Logf("Attempting to update ECSv2: %s", ecs.ID)

	ecsName = tools.RandomString("update-ecs-", 3)
	updateOpts := servers.UpdateOpts{
		Name: ecsName,
	}

	_, err = servers.Update(client, ecs.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("ECSv2 successfully updated: %s", ecs.ID)
	th.AssertNoErr(t, err)

	newECS, err := servers.Get(client, ecs.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ecsName, newECS.Name)
}
