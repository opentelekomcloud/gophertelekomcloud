package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/env_vars"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnvVarsLifecycle(t *testing.T) {
	gatewayID := clients.EnvOS.GetEnv("GATEWAY_ID")
	groupID := clients.EnvOS.GetEnv("GROUP_ID")
	envID := clients.EnvOS.GetEnv("ENV_ID")

	if gatewayID == "" || groupID == "" || envID == "" {
		t.Skip("All of `GATEWAY_ID`, `GROUP_ID` and `ENV_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test_var_env_", 5)

	createOpts := env_vars.CreateOpts{
		GatewayID:     gatewayID,
		GroupID:       groupID,
		EnvID:         envID,
		Name:          name,
		VariableName:  "test-name",
		VariableValue: "test-value",
	}

	createResp, err := env_vars.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, env_vars.Delete(client, gatewayID, createResp.ID))
	})

	getResp, err := env_vars.Get(client, gatewayID, createResp.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)
}

func TestEnvVarsList(t *testing.T) {
	gatewayID := clients.EnvOS.GetEnv("GATEWAY_ID")
	groupID := clients.EnvOS.GetEnv("GROUP_ID")

	if gatewayID == "" || groupID == "" {
		t.Skip("Both `GATEWAY_ID` and `GROUP_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := env_vars.List(client, env_vars.ListOpts{
		GatewayID: gatewayID,
		GroupID:   groupID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}
