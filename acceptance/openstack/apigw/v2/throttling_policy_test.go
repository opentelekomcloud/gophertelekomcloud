package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/api"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/env"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
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

	createResp := CreatePolicy(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, policy.Delete(client, gatewayID, createResp.ID))
	})

	updateOpts := policy.UpdateOpts{
		GatewayID:      gatewayID,
		ThrottleID:     createResp.ID,
		Name:           createResp.Name + "_updated",
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

func TestThrottlingPolicyBinding(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createResp := CreatePolicy(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, policy.Delete(client, gatewayID, createResp.ID))
	})

	groupResp := CreateGroup(client, t, gatewayID)
	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, gatewayID, groupResp.ID))
	})

	groupID := groupResp.ID

	_, createAPIResp := CreateAPI(client, t, gatewayID, groupID)

	t.Cleanup(func() {
		th.AssertNoErr(t, api.Delete(client, gatewayID, createAPIResp.ID))
	})

	envResp := CreateEnv(client, t, gatewayID)

	t.Cleanup(func() {
		manageOpts := api.ManageOpts{
			GatewayID:   gatewayID,
			Action:      "offline",
			EnvID:       envResp.ID,
			ApiID:       createAPIResp.ID,
			Description: "test-api-publish",
		}

		_, err = api.ManageApi(client, manageOpts)
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, env.Delete(client, gatewayID, envResp.ID))
	})

	manageOpts := api.ManageOpts{
		GatewayID:   gatewayID,
		Action:      "online",
		EnvID:       envResp.ID,
		ApiID:       createAPIResp.ID,
		Description: "test-api-publish",
	}

	publishAPI, err := api.ManageApi(client, manageOpts)
	th.AssertNoErr(t, err)

	listUnbound, err := policy.ListAPIUnoundPolicy(client, policy.ListBoundOpts{
		GatewayID:  gatewayID,
		ThrottleID: createResp.ID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listUnbound)

	bindOpts := policy.BindOpts{
		GatewayID: gatewayID,
		PolicyID:  createResp.ID,
		PublishIds: []string{
			publishAPI.PublishID,
		},
	}

	bindResp, err := policy.BindPolicy(client, bindOpts)
	t.Cleanup(func() {
		th.AssertNoErr(t, policy.UnbindPolicy(client, gatewayID, bindResp[0].ID))
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, bindResp)

	listBound, err := policy.ListAPIBoundPolicy(client, policy.ListBoundOpts{
		GatewayID:  gatewayID,
		ThrottleID: createResp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listBound)

	listBoundPolicies, err := policy.ListBoundPolicies(client, policy.ListBindingOpts{
		GatewayID:  gatewayID,
		ThrottleID: createResp.ID,
		ApiID:      createAPIResp.ID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listBoundPolicies)
}

func CreatePolicy(client *golangsdk.ServiceClient, t *testing.T, gatewayID string) *policy.ThrottlingResp {
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
	return createResp
}
