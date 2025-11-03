package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/spec"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGeminiListFlavors(t *testing.T) {
	client, err := clients.NewGeminiDBSpecClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list Gemini db flavors")
	listResp, err := spec.ListFlavors(client, spec.ListFlavorsOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(listResp.Flavors) > 1, true)
}

func TestGeminiGetVersions(t *testing.T) {
	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list Gemini db flavors")
	listResp, err := spec.GetVersions(client, "cassandra")
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}
