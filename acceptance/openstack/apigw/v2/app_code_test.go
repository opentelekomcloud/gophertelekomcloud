package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/app"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/app_code"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAppCodeLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID`needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	_, createAppResp := CreateApp(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, app.Delete(client, gatewayID, createAppResp.ID))
	})

	createOpts := app_code.CreateOpts{
		GatewayID: gatewayID,
		AppID:     createAppResp.ID,
		AppCode:   tools.RandomString("test", 61),
	}

	createResp, err := app_code.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, app_code.Delete(client, gatewayID, createAppResp.ID, createResp.ID))
	})

	tools.PrintResource(t, createResp)

	generatedCode, err := app_code.GenerateAppCode(client, gatewayID, createAppResp.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, generatedCode)

	listResp, err := app_code.List(client, app_code.ListOpts{
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}
