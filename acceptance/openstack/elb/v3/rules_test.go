package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/policies"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRuleWorkflow(t *testing.T) {
	t.Parallel()

	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	lbID := createLoadBalancer(t, client)
	t.Cleanup(func() { deleteLoadbalancer(t, client, lbID) })

	listenerID := createListener(t, client, lbID)
	t.Cleanup(func() { deleteListener(t, client, listenerID) })

	poolID := createPool(t, client, lbID)
	t.Cleanup(func() { deletePool(t, client, poolID) })

	policyID := createPolicy(t, client, listenerID, poolID)
	t.Cleanup(func() { deletePolicy(t, client, policyID) })

	opts := rules.CreateOpts{
		Type:        "PATH",
		CompareType: "REGEX",
		Value:       "^.+$",
	}
	created, err := rules.Create(client, policyID, opts)
	th.AssertNoErr(t, err)
	id := created.Id
	t.Logf("Rule %s added to the policy %s", id, policyID)
	th.CheckEquals(t, opts.Type, created.Type)

	t.Cleanup(func() {
		th.AssertNoErr(t, rules.Delete(client, policyID, id))
		t.Log("Rule removed from policy")
	})

	got, err := rules.Get(client, policyID, id)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, created, got)

	pages, err := rules.List(client, policyID, rules.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	rulesSlice, err := rules.ExtractRules(pages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, 1, len(rulesSlice))
	th.CheckDeepEquals(t, *got, rulesSlice[0])

	updateOpts := rules.UpdateOpts{
		Value: "^.*$",
	}
	updated, err := rules.Update(client, policyID, id, updateOpts)
	th.AssertNoErr(t, err)

	got2, err := rules.Get(client, policyID, id)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updated, got2)
}

func TestRuleWorkflowConditions(t *testing.T) {
	t.Parallel()

	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	lbID := createLoadBalancer(t, client)
	t.Cleanup(func() {
		deleteLoadbalancer(t, client, lbID)
	})

	listener, err := listeners.Create(client, listeners.CreateOpts{
		LoadbalancerID:  lbID,
		Protocol:        "HTTP",
		ProtocolPort:    80,
		EnhanceL7policy: pointerto.Bool(true),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() { deleteListener(t, client, listener.ID) })

	poolID := createPool(t, client, lbID)
	t.Cleanup(func() { deletePool(t, client, poolID) })

	policyID := createPolicy(t, client, listener.ID, poolID)
	t.Cleanup(func() { deletePolicy(t, client, policyID) })
	condition := policies.RuleCondition{
		Key:   "",
		Value: "/",
	}
	opts := rules.CreateOpts{
		Type:        "PATH",
		CompareType: "STARTS_WITH",
		Value:       "/bbb.html",
		Conditions:  []policies.RuleCondition{condition},
	}
	created, err := rules.Create(client, policyID, opts)
	th.AssertNoErr(t, err)
	id := created.Id
	t.Logf("Rule %s added to the policy %s", id, policyID)
	th.CheckEquals(t, opts.Type, created.Type)

	t.Cleanup(func() {
		th.AssertNoErr(t, rules.Delete(client, policyID, id))
		t.Log("Rule removed from policy")
	})

	got, err := rules.Get(client, policyID, id)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, created, got)

	pages, err := rules.List(client, policyID, rules.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	rulesSlice, err := rules.ExtractRules(pages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, 1, len(rulesSlice))
	th.CheckDeepEquals(t, *got, rulesSlice[0])
	conditionUpdate := policies.RuleCondition{
		Key:   "",
		Value: "/home",
	}
	updateOpts := rules.UpdateOpts{
		CompareType: "EQUAL_TO",
		Conditions:  []policies.RuleCondition{conditionUpdate},
	}
	updated, err := rules.Update(client, policyID, id, updateOpts)
	th.AssertNoErr(t, err)

	got2, err := rules.Get(client, policyID, id)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updated, got2)
}

func createPolicy(t *testing.T, client *golangsdk.ServiceClient, listenerID, poolID string) string {
	createOpts := policies.CreateOpts{
		Action:         "REDIRECT_TO_POOL",
		ListenerId:     listenerID,
		RedirectPoolId: poolID,
		Description:    "Go SDK test policy",
		Name:           tools.RandomString("sdk-pol-", 5),
		Position:       pointerto.Int(37),
	}
	created, err := policies.Create(client, createOpts)
	th.AssertNoErr(t, err)
	id := created.Id
	t.Logf("Policy created: %s", id)
	return id
}

func deletePolicy(t *testing.T, client *golangsdk.ServiceClient, policyID string) {
	th.AssertNoErr(t, policies.Delete(client, policyID))
	t.Logf("Policy %s deleted", policyID)
}
