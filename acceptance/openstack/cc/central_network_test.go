package cc

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/capability"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/central_network"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/quota"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCentralNetworkLifeCycle(t *testing.T) {
	t.Skip("CC acceptance tests are skipped")
	if os.Getenv("RUN_CC_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("acctest_cc_cn-", 4)

	created, err := central_network.Create(client, central_network.CreateOpts{
		Name:        name,
		Description: "created by gophertelekomcloud acceptance test",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, created.Name)

	t.Cleanup(func() {
		err = central_network.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})

	got, err := central_network.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.ID, got.ID)

	updated, err := central_network.Update(client, central_network.UpdateOpts{
		CentralNetworkId: created.ID,
		Name:             name + "-updated",
		Description:      pointerto.String("updated description"),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"-updated", updated.Name)
}

func TestCentralNetworkList(t *testing.T) {
	t.Skip("CC acceptance tests are skipped")
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	resp, err := central_network.List(client, central_network.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(resp.CentralNetworks), resp.PageInfo.CurrentCount)
}

func TestCentralNetworkQuotaList(t *testing.T) {
	t.Skip("CC acceptance tests are skipped")
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	resp, err := quota.List(client, quota.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(resp.Quotas) > 0)
	for _, q := range resp.Quotas {
		th.AssertEquals(t, true, q.QuotaKey != "")
	}
}

func TestCentralNetworkCapabilityList(t *testing.T) {
	t.Skip("CC acceptance tests are skipped")
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	resp, err := capability.List(client, capability.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(resp.Capabilities) > 0)
	for _, c := range resp.Capabilities {
		th.AssertEquals(t, true, c.Capability != "")
	}
}
