package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/app"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAppLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID`needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createOpts, createResp := CreateApp(client, t, gatewayID)

	t.Cleanup(func() {
		th.AssertNoErr(t, app.Delete(client, gatewayID, createResp.ID))
	})

	createOpts.Name += "_updated"

	updateResp, err := app.Update(client, createResp.ID, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResp.Name+"_updated", updateResp.Name)

	tools.PrintResource(t, createResp)

	_, err = app.ResetSecret(client, app.ResetOpts{
		AppID:     createResp.ID,
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)

	getResp, err := app.Get(client, gatewayID, createResp.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getResp)

	verifyResp, err := app.VerifyApp(client, gatewayID, createResp.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, verifyResp)
}

func TestAppList(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := app.List(client, app.ListOpts{
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func CreateApp(client *golangsdk.ServiceClient, t *testing.T, gatewayID string) (app.CreateOpts, *app.AppResp) {
	name := tools.RandomString("test_api_", 5)

	createOpts := app.CreateOpts{
		GatewayID:   gatewayID,
		Description: "test app",
		Name:        name,
	}

	createResp, err := app.Create(client, createOpts)
	th.AssertNoErr(t, err)

	return createOpts, createResp
}
