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
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestApiLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID`needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	groupResp := CreateGroup(client, t, gatewayID)
	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, gatewayID, groupResp.ID))
	})

	groupID := groupResp.ID

	createOpts, createResp := CreateAPI(client, t, gatewayID, groupID)

	t.Cleanup(func() {
		th.AssertNoErr(t, api.Delete(client, gatewayID, createResp.ID))
	})

	createOpts.Name += "_updated"
	createOpts.Type = 1
	createOpts.ReqMethod = "POST"
	createOpts.ReqProtocol = "HTTPS"

	updateResp, err := api.Update(client, createResp.ID, createOpts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, createOpts.Name, updateResp.Name)
	th.AssertEquals(t, createOpts.Type, updateResp.Type)
	th.AssertEquals(t, createOpts.ReqMethod, updateResp.ReqMethod)
	th.AssertEquals(t, createOpts.ReqProtocol, updateResp.ReqProtocol)

	envResp := CreateEnv(client, t, gatewayID)

	t.Cleanup(func() {
		manageOpts := api.ManageOpts{
			GatewayID:   gatewayID,
			Action:      "offline",
			EnvID:       envResp.ID,
			ApiID:       createResp.ID,
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
		ApiID:       createResp.ID,
		Description: "test-api-publish",
	}

	manageResp, err := api.ManageApi(client, manageOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, manageResp.EnvID, envResp.ID)
	th.AssertEquals(t, manageResp.Description, manageOpts.Description)

	getResp, err := api.Get(client, gatewayID, createResp.ID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, getResp.Name, createOpts.Name)
	th.AssertEquals(t, getResp.Type, createOpts.Type)
	th.AssertEquals(t, getResp.ReqMethod, createOpts.ReqMethod)
	th.AssertEquals(t, getResp.ReqProtocol, createOpts.ReqProtocol)
}

func TestApiList(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := api.List(client, api.ListOpts{
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func CreateAPI(client *golangsdk.ServiceClient, t *testing.T, gatewayID, groupID string) (api.CreateOpts, *api.ApiResp) {
	name := tools.RandomString("test_api_", 5)

	createOpts := api.CreateOpts{
		GatewayID:   gatewayID,
		Description: "test env",
		Name:        name,
		GroupID:     groupID,
		Type:        2,
		ReqProtocol: "HTTP",
		ReqMethod:   "GET",
		ReqUri:      "/test/http",
		AuthType:    "IAM",
		BackendType: "HTTP",
		BackendApi: &api.BackendApi{
			UrlDomain:   "192.168.189.156:12346",
			ReqProtocol: "HTTP",
			ReqMethod:   "GET",
			ReqUri:      "/test/benchmark",
			Timeout:     5000,
			RetryCount:  "-1",
		},
	}

	createResp, err := api.Create(client, createOpts)
	th.AssertNoErr(t, err)

	return createOpts, createResp
}
