package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/env"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnvLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createResp := CreateEnv(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, env.Delete(client, gatewayID, createResp.ID))
	})

	updateOpts := env.UpdateOpts{
		EnvID:       createResp.ID,
		GatewayID:   gatewayID,
		Description: "test env updated",
		Name:        createResp.Name + "_updated",
	}

	updateResp, err := env.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updateResp)
}

func TestEnvList(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := env.List(client, env.ListOpts{
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func CreateEnv(client *golangsdk.ServiceClient, t *testing.T, id string) *env.EnvResp {
	name := tools.RandomString("test_env_", 5)
	createOpts := env.CreateOpts{
		GatewayID:   id,
		Description: "test env",
		Name:        name,
	}

	createResp, err := env.Create(client, createOpts)
	th.AssertNoErr(t, err)
	return createResp
}
