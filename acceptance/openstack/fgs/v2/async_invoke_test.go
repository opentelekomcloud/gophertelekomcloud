package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/async_config"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAsyncInvokeLifecycle(t *testing.T) {
	// agency := os.Getenv("AGENCY")
	// if agency == "" {
	// 	t.Skip("`AGENCY`needs to be defined to run this test")
	// }
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	funcName := "funcgraph-" + tools.RandomString("acctest", 4)

	createOpts := function.CreateOpts{
		Name:       funcName,
		Package:    "default",
		Runtime:    "Python2.7",
		Timeout:    200,
		Handler:    "index.py",
		MemorySize: 512,
		CodeType:   "inline",
		// Xrole:      agency,
		FuncCode: &function.FuncCode{
			File: "e42a37a22f4988ba7a681e3042e5c7d13c04e6c1"},
	}

	createResp, err := function.Create(client, createOpts)
	th.AssertNoErr(t, err)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	asyncOpts := async_config.UpdateOpts{
		FuncUrn:     funcUrn,
		MaxEventAge: pointerto.Int(1000),
	}

	updateResp, err := async_config.Update(client, asyncOpts)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, id string) {
		err = async_config.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	getResp, err := async_config.Get(client, funcUrn)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, updateResp.MaxEventAge, getResp.MaxEventAge)
	th.AssertEquals(t, updateResp.MaxRetry, getResp.MaxRetry)
	th.AssertEquals(t, updateResp.FuncUrn, getResp.FuncUrn)
	th.AssertEquals(t, updateResp.CreatedTime, getResp.CreatedTime)
}
