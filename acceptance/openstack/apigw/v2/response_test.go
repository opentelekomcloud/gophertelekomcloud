package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/response"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestResponseLifecycle(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}
	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create APIGW Group for instance: %s", gatewayId)
	groupResp := CreateGroup(client, t, gatewayId)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete APIGW Group: %s", groupResp.ID)
		th.AssertNoErr(t, group.Delete(client, gatewayId, groupResp.ID))
	})

	t.Logf("Attempting to Create APIGW Group Response for instance: %s", gatewayId)
	name := tools.RandomString("apigw-group-resp", 3)
	responses := make(map[string]response.ResponseInfo)
	opts := response.CreateOpts{
		GatewayID: gatewayId,
		GroupId:   groupResp.ID,
		Name:      name,
	}
	responses["AUTHORIZER_FAILURE"] = response.ResponseInfo{
		Body:   "{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}",
		Status: 401,
	}
	opts.Responses = responses
	responseRes, err := response.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, responseRes.Name, name)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete APIGW Group Response: %s", responseRes.ID)
		th.AssertNoErr(t, response.Delete(client, gatewayId, groupResp.ID, responseRes.ID))
	})

	t.Logf("Attempting to List APIGW Group Responses")
	_, err = response.List(client, response.ListOpts{
		GatewayID: gatewayId,
		GroupID:   groupResp.ID,
	})
	th.AssertNoErr(t, err)
	// List does not returns responses
	// th.AssertEquals(t, responseRes.Name, resps[1].Name)
	// th.AssertEquals(t, 401, resps[1].Responses["AUTHORIZER_FAILURE"].Status)

	t.Logf("Attempting to Update APIGW Group Response: %s", responseRes.ID)
	nameUpdate := name + "update"
	responsesUpdate := make(map[string]response.ResponseInfo)
	updateOpts := response.CreateOpts{
		GatewayID: gatewayId,
		GroupId:   groupResp.ID,
		Name:      nameUpdate,
	}
	responsesUpdate["AUTHORIZER_FAILURE"] = response.ResponseInfo{
		Body:   "{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}",
		Status: 403,
	}
	responsesUpdate["ACCESS_DENIED"] = response.ResponseInfo{
		Body:   "{\"error_code\":\"$context.error.code\",\"error_msg\":\"$context.error.message\"}",
		Status: 400,
	}
	updateOpts.Responses = responsesUpdate
	updateRes, err := response.Update(client, responseRes.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateRes.Name, nameUpdate)
	th.AssertEquals(t, 403, updateRes.Responses["AUTHORIZER_FAILURE"].Status)
	th.AssertEquals(t, 400, updateRes.Responses["ACCESS_DENIED"].Status)

	t.Logf("Attempting to Obtain APIGW Group Response: %s", responseRes.ID)
	getResp, err := response.Get(client, gatewayId, groupResp.ID, responseRes.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, nameUpdate, getResp.Name)
	th.AssertEquals(t, 403, getResp.Responses["AUTHORIZER_FAILURE"].Status)

	t.Logf("Attempting to Obtain APIGW Group Response Error Type: %s", "AUTHORIZER_FAILURE")
	errorType, err := response.GetErrorType(client, gatewayId, groupResp.ID, responseRes.ID, "AUTHORIZER_FAILURE")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 403, errorType.Status)

	t.Logf("Attempting to Update APIGW Group Response Error Type: %s", "AUTHORIZER_FAILURE")
	errorTypeUpdate, err := response.UpdateErrorType(client, "AUTHORIZER_FAILURE", response.UpdateErrorOpts{
		GatewayID: gatewayId,
		GroupId:   groupResp.ID,
		ID:        getResp.ID,
		Status:    401,
		Body:      "{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 401, errorTypeUpdate.Status)

	t.Logf("Attempting to Delete APIGW Group Response Error Type: %s", "AUTHORIZER_FAILURE")
	err = response.DeleteErrorType(client, gatewayId, groupResp.ID, responseRes.ID, "AUTHORIZER_FAILURE")
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Obtain DELETED APIGW Group Response Error Type: %s", "AUTHORIZER_FAILURE")
	errorTypeDeleted, err := response.GetErrorType(client, gatewayId, groupResp.ID, responseRes.ID, "AUTHORIZER_FAILURE")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 500, errorTypeDeleted.Status)
	th.AssertEquals(t, true, errorTypeDeleted.IsDefault)
}

func TestResponseList(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	groupResp := CreateGroup(client, t, gatewayId)

	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, gatewayId, groupResp.ID))
	})

	listResp, err := response.List(client, response.ListOpts{
		GatewayID: gatewayId,
		GroupID:   groupResp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}
