package v2

import (
	"encoding/base64"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const appCode = `
# -*- coding:utf-8 -*-
import json
def handler (event, context):
    return {
        "statusCode": 200,
        "isBase64Encoded": False,
        "body": json.dumps(event),
        "headers": {
            "Content-Type": "application/json"
        }
    }
`

func TestFunctionGraphLifecycle(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, funcName := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	getCodeResp, err := function.GetCode(client, funcUrn)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResp.CodeURL, getCodeResp.CodeURL)

	updateFuncOpts := function.UpdateFuncCodeOpts{
		FuncUrn:  funcUrn,
		CodeType: "inline",
		FuncCode: function.FuncCode{
			File: base64.StdEncoding.EncodeToString([]byte(appCode)),
		},
	}

	updateResp, err := function.UpdateFuncCode(client, updateFuncOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateResp.CodeType, "inline")
	th.AssertEquals(t, updateResp.CodeFilename, "index.py")

	updateMetaOpts := function.UpdateFuncMetadataOpts{
		FuncUrn:    funcUrn,
		Name:       funcName,
		Runtime:    "Python3.6",
		Timeout:    10,
		Handler:    "index.py",
		MemorySize: 128,
	}

	updateMetaResp, err := function.UpdateFuncMetadata(client, updateMetaOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateMetaResp.FuncName, funcName)

	getMetaResp, err := function.GetMetadata(client, funcUrn)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getMetaResp.CPU, updateMetaResp.CPU)
	th.AssertEquals(t, getMetaResp.MemorySize, updateMetaResp.MemorySize)
	th.AssertEquals(t, getMetaResp.Timeout, updateMetaResp.Timeout)

	updateFuncInstance, err := function.UpdateMaxInstances(client,
		function.UpdateFuncInstancesOpts{
			FuncUrn:        funcUrn,
			MaxInstanceNum: 200,
		})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateFuncInstance.StrategyConfig.Concurrency, 200)

	// API not registered
	// err = function.UpdateStatus(client, funcUrn, "true")
	// th.AssertNoErr(t, err)
}

func TestFunctionGraphList(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	listOpts := function.ListOpts{}
	_, err = function.List(client, listOpts)
	th.AssertNoErr(t, err)
}

func createFunctionGraph(t *testing.T, client *golangsdk.ServiceClient) (*function.FuncGraph, string) {
	funcName := "funcgraph-" + tools.RandomString("acctest", 4)

	createOpts := function.CreateOpts{
		Name:       funcName,
		Package:    "default",
		Runtime:    "Python3.9",
		Timeout:    200,
		Handler:    "index.py",
		MemorySize: 512,
		CodeType:   "zip",
		CodeURL:    "https://regr-func-graph.obs.eu-de.otc.t-systems.com/index.py",
	}

	createResp, err := function.Create(client, createOpts)
	th.AssertNoErr(t, err)

	return createResp, funcName
}
