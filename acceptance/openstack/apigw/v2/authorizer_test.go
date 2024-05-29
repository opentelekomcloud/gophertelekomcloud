package v2

import (
	"encoding/base64"
	"os"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/authorizer"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const appCode = `
# -*- coding:utf-8 -*-
import json
def handler(event, context):
    if event["headers"]["authorization"]=='Basic dXNlcjE6cGFzc3dvcmQ=':
        return {
            'statusCode': 200,
            'body': json.dumps({
                "status":"allow",
                "context":{
                    "user_name":"user1"
                }
            })
        }
    else:
        return {
            'statusCode': 200,
            'body': json.dumps({
                "status":"deny",
                "context":{
                    "code":"1001",
                    "message":"incorrect username or password"
                }
            })
        }
`

func TestCustomAuthorizerLifecycle(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}
	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	clientFg, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create FG function")
	funcResp, _ := createFunction(t, clientFg)

	funcUrn := strings.TrimSuffix(funcResp.FuncURN, ":latest")

	t.Cleanup(func() {
		t.Logf("Attempting to Delete FG function: %s", funcResp.FuncID)
		th.AssertNoErr(t, function.Delete(clientFg, funcUrn))
	})

	authName := "apigw_authorizer_" + tools.RandomString("acctest", 4)
	t.Logf("Attempting to Create APIGW Custom authorizer for instance: %s", gatewayId)
	opts := authorizer.CreateOpts{
		GatewayID:      gatewayId,
		Name:           authName,
		Type:           "FRONTEND",
		AuthorizerType: "FUNC",
		FunctionUrn:    funcUrn,
		Identities: []authorizer.Identity{
			{
				Name:     "user_name",
				Location: "QUERY",
			},
		},
		Ttl:      pointerto.Int(60),
		NeedBody: pointerto.Bool(true),
	}
	createResp, err := authorizer.Create(client, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete APIGW Custom authorizer: %s", createResp.ID)
		th.AssertNoErr(t, authorizer.Delete(client, gatewayId, createResp.ID))
	})

	t.Logf("Attempting to List APIGW Custom authorizers")
	authorizers, err := authorizer.List(client, authorizer.ListOpts{
		GatewayID: gatewayId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResp.Name, authorizers[0].Name)

	t.Logf("Attempting to Update APIGW Custom authorizer %s", createResp.ID)
	updateOpts := authorizer.CreateOpts{
		GatewayID:      gatewayId,
		Name:           authName + "updated",
		Type:           "FRONTEND",
		AuthorizerType: "FUNC",
		FunctionUrn:    funcUrn,
		Identities: []authorizer.Identity{
			{
				Name:     "user_name",
				Location: "QUERY",
			},
		},
		Ttl:      pointerto.Int(45),
		NeedBody: pointerto.Bool(true),
	}
	authorizerUpdate, err := authorizer.Update(client, createResp.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 45, *authorizerUpdate.Ttl)
	th.AssertEquals(t, authName+"updated", authorizerUpdate.Name)

	t.Logf("Attempting to Obtain updated APIGW Custom authorizer %s", authorizerUpdate.ID)
	ch, err := authorizer.Get(client, gatewayId, authorizerUpdate.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 45, *ch.Ttl)
	th.AssertEquals(t, "user_name", ch.Identities[0].Name)
	th.AssertEquals(t, "QUERY", ch.Identities[0].Location)
}

func TestCustomAuthorizerList(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")

	if gatewayId == "" {
		t.Skip("`GATEWAY_ID` need to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := authorizer.List(client, authorizer.ListOpts{
		GatewayID: gatewayId,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
}

func createFunction(t *testing.T, client *golangsdk.ServiceClient) (*function.FuncGraph, string) {
	funcName := "func-apigw" + tools.RandomString("acctest", 4)

	createOpts := function.CreateOpts{
		Name:        funcName,
		Description: "API custom authorization test",
		Package:     "default",
		Runtime:     "Python3.9",
		Timeout:     3,
		Handler:     "index.handler",
		MemorySize:  128,
		CodeType:    "inline",
		FuncCode: &function.FuncCode{
			File: base64.StdEncoding.EncodeToString([]byte(appCode))},
	}

	createResp, err := function.Create(client, createOpts)
	th.AssertNoErr(t, err)

	return createResp, funcName
}
