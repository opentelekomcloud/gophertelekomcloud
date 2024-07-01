package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	hss "github.com/opentelekomcloud/gophertelekomcloud/openstack/hss/v5/host"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServerGroupList(t *testing.T) {
	client, err := clients.NewHssClient()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, client)
	listResp, err := hss.List(client, hss.ListOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestServerList(t *testing.T) {
	client, err := clients.NewHssClient()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, client)
	listResp, err := hss.ListHost(client, hss.ListHostOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}
