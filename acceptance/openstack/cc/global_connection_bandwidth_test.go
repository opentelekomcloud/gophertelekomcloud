package cc

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	gcb "github.com/opentelekomcloud/gophertelekomcloud/openstack/cc/v3/global_connection_bandwidth"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGlobalConnectionBandwidthLifeCycle(t *testing.T) {
	if os.Getenv("RUN_CC_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("acctest_cc_gcb-", 4)

	created, err := gcb.Create(client, gcb.CreateOpts{
		Name:        name,
		Description: "created by gophertelekomcloud acceptance test",
		Bordercross: pointerto.Bool(false),
		Type:        "Region",
		ChargeMode:  "bwd",
		Size:        5,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, created.Name)

	t.Cleanup(func() {
		err = gcb.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})

	got, err := gcb.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.ID, got.ID)

	updated, err := gcb.Update(client, gcb.UpdateOpts{
		ID:   created.ID,
		Name: name + "-updated",
		Size: 10,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"-updated", updated.Name)
}

func TestGlobalConnectionBandwidthList(t *testing.T) {
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	resp, err := gcb.List(client, gcb.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(resp.GlobalConnectionBandwidths), resp.PageInfo.CurrentCount)
}

func TestGlobalConnectionBandwidthConfigs(t *testing.T) {
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	resp, err := gcb.GetConfigs(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(resp.Configs.ChargeMode) > 0)
}

func TestGlobalConnectionBandwidthSites(t *testing.T) {
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	resp, err := gcb.ListSites(client, gcb.ListSitesOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(resp.SiteInfos), resp.PageInfo.CurrentCount)
	for _, s := range resp.SiteInfos {
		th.AssertEquals(t, true, s.SiteCode != "")
	}
}

func TestGlobalConnectionBandwidthSupportBindings(t *testing.T) {
	client, err := clients.NewCCClient()
	th.AssertNoErr(t, err)

	_, err = gcb.ListSupportBindings(client, gcb.ListSupportBindingsOpts{
		BindingService: "CC",
	})
	th.AssertNoErr(t, err)
}
