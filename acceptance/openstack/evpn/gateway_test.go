package evpn

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evpn/v5/gateway"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGatewaysList(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)
	gateways, err := gateway.List(client)
	th.AssertNoErr(t, err)
	for _, gw := range gateways {
		tools.PrintResource(t, gw)
	}
}

func TestGatewaysAZsList(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)
	azs, err := gateway.ListAZs(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, azs)
}
