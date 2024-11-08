package evpn

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	cgw "github.com/opentelekomcloud/gophertelekomcloud/openstack/evpn/v5/customer-gateway"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCustomerGatewaysList(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)
	gateways, err := cgw.List(client, cgw.ListOpts{})
	th.AssertNoErr(t, err)
	for _, gw := range gateways {
		tools.PrintResource(t, gw)
	}
}

func TestCustomerGatewayVPCLifecycle(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("acc_evpn_cgw_", 5)
	cOpts := cgw.CreateOpts{
		Name:    name,
		IdType:  "ip",
		IdValue: "10.1.2.10",
		BgpAsn:  pointerto.Int(65000),
		Tags: []tags.ResourceTag{
			{
				Key:   "NewKey",
				Value: "NewValue",
			},
		},
	}
	t.Logf("Attempting to CREATE Enterprise VPN Customer Gateway: %s", name)
	gw, err := cgw.Create(client, cOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, gw.Name)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE Enterprise VPN Customer Gateway: %s", gw.ID)
		th.AssertNoErr(t, cgw.Delete(client, gw.ID))
	})
	t.Logf("Attempting to GET Enterprise VPN Customer Gateway: %s", gw.ID)
	gwGet, err := cgw.Get(client, gw.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 65000, gwGet.BgpAsn)
	th.AssertEquals(t, "10.1.2.10", gwGet.IdValue)
	th.AssertEquals(t, "ip", gwGet.IdType)

	t.Logf("Attempting to UPDATE Enterprise VPN Customer Gateway: %s", gw.ID)
	gwUpdate, err := cgw.Update(client, cgw.UpdateOpts{
		GatewayID: gw.ID,
		Name:      name + "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"updated", gwUpdate.Name)

	t.Logf("Attempting to LIST Enterprise VPN Customer Gateways: %s", gw.ID)
	gws, err := cgw.List(client, cgw.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(gws))
}
