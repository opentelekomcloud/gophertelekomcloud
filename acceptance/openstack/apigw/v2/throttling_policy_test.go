package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	policy "github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/tr_policy"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestThrottlingPolicyLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test_policy_", 5)

	createOpts := policy.CreateOpts{
		GatewayID:      gatewayID,
		Name:           name,
		ApiCallLimits:  pointerto.Int(200),
		TimeInterval:   pointerto.Int(10000),
		TimeUnit:       "SECOND",
		Description:    "test throttling policy",
		AppCallLimits:  pointerto.Int(100),
		UserCallLimits: pointerto.Int(100),
	}

	createResp, err := policy.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, policy.Delete(client, gatewayID, createResp.ID))
	})

	updateOpts := policy.UpdateOpts{
		GatewayID:      gatewayID,
		ThrottleID:     createResp.ID,
		Name:           name + "_updated",
		ApiCallLimits:  pointerto.Int(199),
		TimeInterval:   pointerto.Int(999),
		TimeUnit:       "MINUTE",
		Description:    "test throttling policy updated",
		AppCallLimits:  pointerto.Int(50),
		UserCallLimits: pointerto.Int(50),
	}
	updateResp, err := policy.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	getResp, err := policy.Get(client, gatewayID, updateResp.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)
}

func TestThrottlingPolicyList(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := policy.List(client, policy.ListOpts{
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}
