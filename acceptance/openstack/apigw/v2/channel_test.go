package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/channel"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
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
	ecs := openstack.CreateServer(t, ecsClient,
		tools.RandomString("hss_group-member-", 3),
		"Standard_Debian_10_latest",
		"s2.large.2",
	)
	th.AssertNoErr(t, err)

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
