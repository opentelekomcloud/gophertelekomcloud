package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/channel"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/flavors"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVpcChannelLifecycle(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}
	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create APIGW VPC Channel for instance: %s", gatewayId)
	createResp := CreateChannel(client, t, gatewayId)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete APIGW VPC Channel: %s", createResp.ID)
		th.AssertNoErr(t, channel.Delete(client, gatewayId, createResp.ID))
	})

	t.Logf("Attempting to List APIGW VPC Channels")
	channels, err := channel.List(client, channel.ListOpts{
		GatewayID: gatewayId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResp.Name, channels[0].Name)

	t.Logf("Attempting to Update APIGW VPC Channel %s", createResp.ID)
	updateOpts := channel.CreateOpts{
		GatewayID:   gatewayId,
		Name:        createResp.Name,
		Port:        80,
		LbAlgorithm: 2,
		MemberType:  "ecs",
		Type:        2,
		VpcHealthConfig: &channel.VpcHealthConfig{
			Protocol:           "HTTPS",
			HealthyThreshold:   10,
			UnhealthyThreshold: 10,
			Interval:           300,
			Timeout:            30,
			Path:               "/gopher-up/",
			Method:             "HEAD",
			HttpCode:           "201,202,303-404",
		},
	}
	channelUpdate, err := channel.Update(client, createResp.ID, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Obtain updated APIGW VPC Channel %s", channelUpdate.ID)
	ch, err := channel.Get(client, gatewayId, channelUpdate.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 300, ch.VpcHealthConfig.Interval)
	th.AssertEquals(t, 30, ch.VpcHealthConfig.Timeout)
	th.AssertEquals(t, 10, ch.VpcHealthConfig.UnhealthyThreshold)
	th.AssertEquals(t, 10, ch.VpcHealthConfig.HealthyThreshold)
	th.AssertEquals(t, createResp.Name, ch.Name)

	// API work is unstable
	// t.Logf("Attempting to Obtain members of APIGW VPC Channel %s", channelUpdate.ID)
	// members, err := channel.ListMembers(client, channel.ListMembersOpts{
	// 	GatewayID: gatewayId,
	// 	ChannelID: ch.ID,
	// })
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, 1, len(members))
	//
	// t.Logf("Attempting to Delete member from APIGW VPC Channel %s", channelUpdate.ID)
	// err = channel.DeleteMember(client, gatewayId, ch.ID, members[0].ID)
	// th.AssertNoErr(t, err)
}

func TestVpcChannelList(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := channel.List(client, channel.ListOpts{
		GatewayID: gatewayId,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}

func CreateChannel(client *golangsdk.ServiceClient, t *testing.T, id string) *channel.ChannelResp {
	name := tools.RandomString("apigw_channel-", 3)

	t.Logf("Attempting to Create member for VPC Channel")
	ecsClient, err := clients.NewComputeV2Client()
	ecs := CreateServer(ecsClient, t)

	t.Cleanup(func() {
		th.AssertNoErr(t, servers.Delete(ecsClient, ecs.ID).ExtractErr())
	})

	createOpts := channel.CreateOpts{
		GatewayID:   id,
		Name:        name,
		Port:        80,
		LbAlgorithm: 1,
		MemberType:  "ecs",
		Type:        2,
		Members: []channel.Members{
			{
				EcsId:   ecs.ID,
				EcsName: ecs.Name,
			},
		},
		VpcHealthConfig: &channel.VpcHealthConfig{
			Protocol:           "TCP",
			HealthyThreshold:   1,
			UnhealthyThreshold: 1,
			Interval:           1,
			Timeout:            1,
			Path:               "/gopher/",
			Method:             "HEAD",
			HttpCode:           "201,202,303-404",
		},
	}

	createResp, err := channel.Create(client, createOpts)
	th.AssertNoErr(t, err)
	return createResp
}

func CreateServer(client *golangsdk.ServiceClient, t *testing.T) *servers.Server {
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if networkID == "" {
		t.Skip("OS_NETWORK_ID env var is missing but ECS test requires using existing network")
	}
	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	name := tools.RandomString("apigw_channel-member-", 3)
	imageV2Client, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	image, err := images.ListImages(imageV2Client, images.ListImagesOpts{
		Name: "Standard_Debian_10_latest",
	})
	th.AssertNoErr(t, err)

	flavorID, err := flavors.IDFromName(client, "s2.large.2")
	th.AssertNoErr(t, err)

	createOpts := servers.CreateOpts{
		Name:      name,
		ImageRef:  image[0].Id,
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

	server, err := servers.Get(client, ecs.ID).Extract()
	th.AssertNoErr(t, err)

	return server
}
