package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/async_config"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAsyncInvokeLifecycle(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

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
