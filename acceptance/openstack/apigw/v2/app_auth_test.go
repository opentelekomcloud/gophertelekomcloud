package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/api"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/app"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/app_auth"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAppAuthLifecycle(t *testing.T) {
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

	groupResp := CreateGroup(client, t, gatewayID)
	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, gatewayID, groupResp.ID))
	})

	groupID := groupResp.ID

	_, createApiResp := CreateAPI(client, t, gatewayID, groupID)

	t.Cleanup(func() {
		th.AssertNoErr(t, api.Delete(client, gatewayID, createApiResp.ID))
	})

	createAuthResp, err := app_auth.Create(client, app_auth.CreateAuthOpts{
		GatewayID: gatewayID,
		EnvID:     "DEFAULT_ENVIRONMENT_RELEASE_ID",
		AppIDs: []string{
			createAppResp.ID,
		},
		ApiIDs: []string{
			createApiResp.ID,
		},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, app_auth.Delete(client, gatewayID, createAuthResp[0].ID))
	})

	tools.PrintResource(t, createAuthResp)

}
