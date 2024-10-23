package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/dependency_version"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDependencyVersionLifecycle(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	createOpts := dependency_version.CreateOpts{
		Name:       tools.RandomString("dep-version-", 4),
		Runtime:    "Python2.7",
		DependType: "zip",
		DependFile: "",
	}

	createDepResp, err := dependency_version.Create(client, createOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, createDepResp)
}
