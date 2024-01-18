package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/api"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/env"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/key"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSignatureKeyLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test_sign_key_", 5)

	createOpts := key.CreateOpts{
		GatewayID:     gatewayID,
		Name:          name,
		SignType:      "aes",
		SignAlgorithm: "aes-256-cfb",
	}

	createResp, err := key.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, key.Delete(client, gatewayID, createResp.ID))
	})

	updateOpts := key.UpdateOpts{
		Name:          createOpts.Name,
		GatewayID:     createOpts.GatewayID,
		SignID:        createResp.ID,
		SignAlgorithm: "aes-128-cfb",
		SignType:      "aes",
	}

	updateResp, err := key.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updateResp)
}

func TestSignatureKeyList(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := key.List(client, key.ListOpts{
		GatewayID: gatewayID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestSignatureKeyBinding(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")

	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test_sign_key_", 5)

	createOpts := key.CreateOpts{
		GatewayID:     gatewayID,
		Name:          name,
		SignType:      "aes",
		SignAlgorithm: "aes-256-cfb",
	}

	createResp, err := key.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, key.Delete(client, gatewayID, createResp.ID))
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

	bindOpts := key.BindOpts{
		GatewayID: gatewayID,
		SignID:    createResp.ID,
		PublishIds: []string{
			publishAPI.PublishID,
		},
	}

	bindResp, err := key.BindKey(client, bindOpts)
	th.AssertNoErr(t, err)

	listBind, err := key.ListBoundKeys(client, key.ListBindingOpts{
		GatewayID: gatewayID,
		ApiID:     createAPIResp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listBind)

	listBindAPI, err := key.ListAPIBoundKeys(client, key.ListBoundOpts{
		GatewayID: gatewayID,
		SignID:    createResp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listBindAPI)

	th.AssertNoErr(t, key.UnbindKey(client, gatewayID, bindResp[0].ID))

	listUnbound, err := key.ListUnboundKeys(client, key.ListUnbindOpts{
		GatewayID: gatewayID,
		SignID:    createResp.ID,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listUnbound)

}
