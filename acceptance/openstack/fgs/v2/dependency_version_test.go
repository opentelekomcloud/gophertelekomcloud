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
		Name:        tools.RandomString("dep-version-", 4),
		Runtime:     "Python3.9",
		DependType:  "obs",
		DependLink:  "https://regr-func-graph.obs.eu-de.otc.t-systems.com/requirements.zip",
		Description: "test dependency",
	}

	createDepResp, err := dependency_version.Create(client, createOpts)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, id string) {
		err = dependency_version.Delete(client, createDepResp.DepId, "1")
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	getDepResp, err := dependency_version.Get(client, createDepResp.DepId, "1")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createDepResp.Name, getDepResp.Name)
	th.AssertEquals(t, createDepResp.Runtime, getDepResp.Runtime)
	th.AssertEquals(t, createDepResp.Etag, getDepResp.Etag)

	listDepResp, err := dependency_version.ListDependencies(client, dependency_version.ListDependenciesOpts{
		DependencyType: "obs",
		Runtime:        "Python3.9",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(listDepResp.Dependencies) > 1, true)

	listDepVerResp, err := dependency_version.ListDependencyVersions(client, dependency_version.ListDependencyVersionOpts{
		DependId: createDepResp.DepId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listDepVerResp.Dependencies[0].Name, getDepResp.Name)
	th.AssertEquals(t, listDepVerResp.Dependencies[0].Runtime, getDepResp.Runtime)
	th.AssertEquals(t, listDepVerResp.Dependencies[0].Etag, getDepResp.Etag)

}
