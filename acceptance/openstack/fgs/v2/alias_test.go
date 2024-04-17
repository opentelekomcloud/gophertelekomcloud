package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/alias"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphListAliases(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)
	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	publishOpts := alias.PublishOpts{
		FuncUrn:     funcUrn,
		Version:     "new-version",
		Description: "terraform",
	}

	publishOptsResp, err := alias.PublishVersion(client, publishOpts)
	th.AssertNoErr(t, err)

	createAliasOpts := alias.CreateAliasOpts{
		Name:        "test-alias",
		Version:     "new-version",
		Description: "terraform alias",
		FuncUrn:     publishOptsResp.FuncURN,
	}

	listVersion, err := alias.ListVersion(client, alias.ListVersionOpts{
		FuncUrn: publishOptsResp.FuncURN,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listVersion)

	createAliasResp, err := alias.CreateAlias(client, createAliasOpts)
	th.AssertNoErr(t, err)

	listAliasResp, err := alias.ListAlias(client, publishOptsResp.FuncURN)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, listAliasResp[0].Name, createAliasResp.Name)
	th.AssertEquals(t, listAliasResp[0].Version, createAliasResp.Version)

	updateAliasResp, err := alias.UpdateAlias(client, alias.UpdateAliasOpts{
		FuncUrn:     funcUrn,
		AliasName:   "test-alias",
		Version:     "new-version",
		Description: "new description",
	})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, listAliasResp[0].Name, updateAliasResp.Name)
	th.AssertEquals(t, updateAliasResp.Description, "new description")

	getAliasResp, err := alias.GetAlias(client, publishOptsResp.FuncURN, "test-alias")
	th.AssertNoErr(t, err)

	th.AssertEquals(t, listAliasResp[0].Name, getAliasResp.Name)
	th.AssertEquals(t, getAliasResp.Description, updateAliasResp.Description)
	th.AssertEquals(t, listAliasResp[0].AliasUrn, getAliasResp.AliasUrn)

	th.AssertNoErr(t, alias.Delete(client, publishOptsResp.FuncURN, "test-alias"))
}
