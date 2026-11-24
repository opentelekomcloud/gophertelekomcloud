package dns

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/resolverrules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestResolverRulesLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but DNS resolver rule test requires using existing network")
	}

	client, err := clients.NewDNSV2Client()
	th.AssertNoErr(t, err)

	// Cleanups run LIFO: the rule is deleted before the endpoint.
	ep := createEndpoint(t, client, "outbound", subnetID)

	createOpts := resolverrules.CreateOpts{
		Name:        tools.RandomString("acc-dns-rule-", 3),
		DomainName:  strings.ToLower(tools.RandomString("acc", 5)) + ".example.com",
		EndpointID:  ep.ID,
		IPAddresses: []resolverrules.IPInfo{{IP: "1.1.1.1"}},
	}
	rule, err := resolverrules.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created DNS resolver rule: %s", rule.ID)

	t.Cleanup(func() {
		th.AssertNoErr(t, resolverrules.Delete(client, rule.ID))
		th.AssertNoErr(t, waitForDeleted(func() error {
			_, err := resolverrules.Get(client, rule.ID)
			return err
		}))
		t.Logf("Deleted DNS resolver rule: %s", rule.ID)
	})

	th.AssertNoErr(t, waitForRule(client, rule.ID, func(r *resolverrules.ResolverRule) bool {
		return r.Status == "ACTIVE"
	}))

	got, err := resolverrules.Get(client, rule.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)
	th.AssertEquals(t, createOpts.Name, got.Name)
	th.AssertEquals(t, ep.ID, got.EndpointID)
	th.AssertEquals(t, 1, len(got.IPAddresses))
	th.AssertEquals(t, "1.1.1.1", got.IPAddresses[0].IP)

	list, err := resolverrules.List(client, resolverrules.ListOpts{EndpointID: ep.ID})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(list.ResolverRules))
	th.AssertEquals(t, rule.ID, list.ResolverRules[0].ID)

	updateOpts := resolverrules.UpdateOpts{
		Name:        createOpts.Name + "-upd",
		IPAddresses: []resolverrules.IPInfo{{IP: "1.1.1.1"}, {IP: "8.8.8.8"}},
	}
	updated, err := resolverrules.Update(client, rule.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, updated.Name)
	th.AssertNoErr(t, waitForRule(client, rule.ID, func(r *resolverrules.ResolverRule) bool {
		return r.Status == "ACTIVE" && len(r.IPAddresses) == 2
	}))

	router, err := resolverrules.AssociateVPC(client, rule.ID, resolverrules.RouterOpts{RouterID: vpcID})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpcID, router.RouterID)
	th.AssertNoErr(t, waitForRule(client, rule.ID, func(r *resolverrules.ResolverRule) bool {
		return routerStatus(r, vpcID) == "ACTIVE"
	}))

	_, err = resolverrules.DisassociateVPC(client, rule.ID, resolverrules.RouterOpts{RouterID: vpcID})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForRule(client, rule.ID, func(r *resolverrules.ResolverRule) bool {
		return routerStatus(r, vpcID) == ""
	}))
}

// waitForRule waits until done returns true for the current state of the rule.
func waitForRule(client *golangsdk.ServiceClient, ruleID string, done func(*resolverrules.ResolverRule) bool) error {
	return golangsdk.WaitFor(300, func() (bool, error) {
		r, err := resolverrules.Get(client, ruleID)
		if err != nil {
			return false, err
		}
		return done(r), nil
	})
}

// routerStatus returns the status of the VPC association, or "" if the VPC isn't associated.
func routerStatus(r *resolverrules.ResolverRule, vpcID string) string {
	for _, rt := range r.Routers {
		if rt.RouterID == vpcID {
			return rt.Status
		}
	}
	return ""
}
