package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGroupLifecycle(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}
	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createResp := CreateGroup(client, t, gatewayId)

	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, gatewayId, createResp.ID))
	})

	updateOpts := group.UpdateOpts{
		GroupID:     createResp.ID,
		Name:        createResp.Name + "-updated",
		Description: "test updated",
		GatewayID:   gatewayId,
	}

	_, err = group.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	getResp, err := group.Get(client, gatewayId, createResp.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)
}

func TestGroupList(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := group.List(client, group.ListOpts{
		GatewayID: gatewayId,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}

func CreateGroup(client *golangsdk.ServiceClient, t *testing.T, id string) *group.GroupResp {
	name := tools.RandomString("apigw_group-", 3)

	createOpts := group.CreateOpts{
		Name:        name,
		Description: "test",
		GatewayID:   id,
	}

	createResp, err := group.Create(client, createOpts)
	th.AssertNoErr(t, err)
	return createResp
}
