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

	syncResp, err := invoke.LaunchSync(client, funcUrn)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, syncResp)
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
