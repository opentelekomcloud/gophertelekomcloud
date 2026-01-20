package v2

import (
	"os"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/invoke"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphSync(t *testing.T) {
	if funcGraph := os.Getenv("FUNCGRAPH_TEST"); funcGraph == "" {
		t.Skip("`FUNCGRAPH_TEST`needs to be defined to run this test")
	}
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	body := map[string]interface{}{
		"k": "v",
		"t": "start",
	}

	syncResp, syncRespHeaders, err := invoke.LaunchSync(client, funcUrn, body, invoke.NewLaunchSyncHeaders())

	th.AssertNoErr(t, err)

	// test for function error
	th.AssertEquals(t, false, syncRespHeaders.IsFuncErr)

	// test for http status
	th.AssertEquals(t, 200, syncResp.Status)

	tools.PrintResource(t, syncResp)
	tools.PrintResource(t, syncRespHeaders)
}

func TestFunctionGraphAsync(t *testing.T) {
	if funcGraph := os.Getenv("FUNCGRAPH_TEST"); funcGraph == "" {
		t.Skip("`FUNCGRAPH_TEST`needs to be defined to run this test")
	}
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	asyncOpts := map[string]string{
		"k":    "v",
		"test": "start",
	}

	syncResp, err := invoke.LaunchAsync(client, funcUrn, asyncOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, syncResp)
}
