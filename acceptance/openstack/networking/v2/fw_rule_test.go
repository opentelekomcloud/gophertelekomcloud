package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/fwaas_v2/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFwRuleLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create BandwidthV2")
	ruleName := tools.RandomString("band-create", 3)
	createOpts := rules.CreateOpts{
		Name:            ruleName,
		Protocol:        "tcp",
		Action:          "deny",
		DestinationPort: "23",
		Enabled:         pointerto.Bool(true),
	}

	rule, err := rules.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete FwaasV2: %s", rule.ID)
		err := rules.Delete(client, rule.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted FwaasV2: %s", rule.ID)
	})
	th.AssertEquals(t, createOpts.Name, rule.Name)
	th.AssertEquals(t, createOpts.DestinationPort, rule.DestinationPort)
	t.Logf("Created FwaasV2: %s", rule.ID)

	t.Logf("Attempting to update FwaasV2: %s", rule.ID)
	updateOpts := rules.UpdateOpts{
		Protocol: pointerto.String("icmp"),
	}
	updatedRule, err := rules.Update(client, rule.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated FwaasV2: %s", rule.ID)
	th.AssertEquals(t, updatedRule.Protocol, *updateOpts.Protocol)

	t.Logf("Attempting to create Ipv6")
	ipv6RuleName := tools.RandomString("ipv6-create", 3)
	createIpv6Opts := rules.CreateOpts{
		Protocol:             "tcp",
		Action:               "deny",
		Name:                 ipv6RuleName,
		IPVersion:            6,
		DestinationIPAddress: "2001:db8::",
		Enabled:              pointerto.Bool(true),
	}

	rule6, err := rules.Create(client, createIpv6Opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, rule6.IPVersion, 6)
}
