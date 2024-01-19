package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	policy "github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/tr_policy"
	specials "github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/tr_specials"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSpecialThrottleLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createPolicyResp := CreatePolicy(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, policy.Delete(client, gatewayID, createPolicyResp.ID))
	})

	createOpts := specials.CreateOpts{
		GatewayID:  gatewayID,
		ThrottleID: createPolicyResp.ID,
		CallLimits: 100,
		ObjectType: "USER",
		ObjectID:   "356de8eb7a8742168586e5daf5339965",
	}

	createResp, err := specials.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, specials.Delete(client, gatewayID, createPolicyResp.ID, createResp.ID))
	})

	tools.PrintResource(t, createResp)

	updateOpts := specials.UpdateOpts{
		GatewayID:       gatewayID,
		ThrottleID:      createPolicyResp.ID,
		SpecialPolicyID: createResp.ID,
		CallLimits:      90,
	}

	updateResp, err := specials.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updateResp)
}

func TestSpecialThrottleList(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createPolicyResp := CreatePolicy(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, policy.Delete(client, gatewayID, createPolicyResp.ID))
	})

	listResp, err := specials.List(client, specials.ListOpts{
		GatewayID:  gatewayID,
		ThrottleID: createPolicyResp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}
