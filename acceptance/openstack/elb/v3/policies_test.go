package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyWorkflow(t *testing.T) {
	t.Parallel()

	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	lbID := createLoadBalancer(t, client)
	t.Cleanup(func() { deleteLoadbalancer(t, client, lbID) })

	listenerID := createListener(t, client, lbID)
	t.Cleanup(func() { deleteListener(t, client, listenerID) })

	poolID := createPool(t, client, lbID)
	t.Cleanup(func() { deletePool(t, client, poolID) })

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
	t.Logf("Created L7 Policy")
	id := created.Id

	t.Cleanup(func() {
		th.AssertNoErr(t, policies.Delete(client, id))
		t.Log("Deleted L7 Policy")
	})

	got, err := policies.Get(client, id)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, created, got)

	listOpts := policies.ListOpts{
		ListenerId: []string{listenerID},
	}
	pages, err := policies.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	policySlice, err := policies.ExtractPolicies(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(policySlice))
	th.AssertEquals(t, id, policySlice[0].Id)

	updateOpts := policies.UpdateOpts{
		Name: tools.RandomString("updated-", 5),
	}
	updated, err := policies.Update(client, id, updateOpts)
	th.AssertNoErr(t, err)
	t.Log("Updated l7 Policy")
	th.AssertEquals(t, created.Action, updated.Action)
	th.AssertEquals(t, id, updated.Id)

	got2, _ := policies.Get(client, id)
	th.AssertDeepEquals(t, updated, got2)
}

func TestPolicyWorkflowFixedResponse(t *testing.T) {
	t.Parallel()

	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	lbID := createLoadBalancer(t, client)
	t.Cleanup(func() { deleteLoadbalancer(t, client, lbID) })

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

	createOpts := policies.CreateOpts{
		Action:     "FIXED_RESPONSE",
		ListenerId: listener.ID,
		FixedResponseConfig: policies.FixedResponseOptions{
			StatusCode:  "200",
			ContentType: "text/plain",
			MessageBody: "Fixed Response",
		},
		Description: "Go SDK test policy",
		Name:        tools.RandomString("sdk-pol-", 5),
		Position:    pointerto.Int(37),
	}
	created, err := policies.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created L7 Policy")
	id := created.Id

	t.Cleanup(func() {
		th.AssertNoErr(t, policies.Delete(client, id))
		t.Log("Deleted L7 Policy")
	})

	got, err := policies.Get(client, id)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, created, got)

	listOpts := policies.ListOpts{
		ListenerId: []string{listener.ID},
	}
	pages, err := policies.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	policySlice, err := policies.ExtractPolicies(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(policySlice))
	th.AssertEquals(t, id, policySlice[0].Id)

	updateOpts := policies.UpdateOpts{
		Name: tools.RandomString("updated-", 5),
		FixedResponseConfig: policies.FixedResponseOptions{
			StatusCode:  "200",
			ContentType: "text/plain",
			MessageBody: "Fixed Response Update",
		},
	}
	updated, err := policies.Update(client, id, updateOpts)
	th.AssertNoErr(t, err)
	t.Log("Updated l7 Policy")
	th.AssertEquals(t, created.Action, updated.Action)
	th.AssertEquals(t, id, updated.Id)

	got2, _ := policies.Get(client, id)
	th.AssertDeepEquals(t, updated, got2)
}

func TestPolicyWorkflowUlrRedirect(t *testing.T) {
	t.Parallel()

	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	lbID := createLoadBalancer(t, client)
	t.Cleanup(func() { deleteLoadbalancer(t, client, lbID) })

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

	createOpts := policies.CreateOpts{
		Action:      "REDIRECT_TO_URL",
		ListenerId:  listener.ID,
		RedirectUrl: "https://www.bing.com:443",
		RedirectUrlConfig: policies.RedirectUrlOptions{
			StatusCode: "302",
			Protocol:   "${protocol}",
			Port:       "${port}",
			Path:       "${path}",
			Query:      "${query}&name=my_name",
			Host:       "${host}",
		},
		Description: "Go SDK test policy",
		Name:        tools.RandomString("sdk-pol-", 5),
		Position:    pointerto.Int(37),
	}
	created, err := policies.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created L7 Policy")
	id := created.Id

	t.Cleanup(func() {
		th.AssertNoErr(t, policies.Delete(client, id))
		t.Log("Deleted L7 Policy")
	})

	got, err := policies.Get(client, id)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, created, got)

	listOpts := policies.ListOpts{
		ListenerId: []string{listener.ID},
	}
	pages, err := policies.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	policySlice, err := policies.ExtractPolicies(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(policySlice))
	th.AssertEquals(t, id, policySlice[0].Id)

	updateOpts := policies.UpdateOpts{
		Name: tools.RandomString("updated-", 5),
		RedirectUrlConfig: policies.RedirectUrlOptions{
			StatusCode: "308",
			Protocol:   "${protocol}",
			Port:       "${port}",
			Path:       "${path}",
			Query:      "${query}&name=my_name",
			Host:       "${host}",
		},
	}
	updated, err := policies.Update(client, id, updateOpts)
	th.AssertNoErr(t, err)
	t.Log("Updated l7 Policy")
	th.AssertEquals(t, created.Action, updated.Action)
	th.AssertEquals(t, id, updated.Id)

	got2, _ := policies.Get(client, id)
	th.AssertDeepEquals(t, updated, got2)
}
