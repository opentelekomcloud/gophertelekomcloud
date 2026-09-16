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

func TestGeminiListInfluxFlavors(t *testing.T) {
	client, err := clients.NewGeminiDBSpecClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list GeminiDB Influx flavors")
	listResp, err := spec.ListFlavors(client, spec.ListFlavorsOpts{
		EngineName: "influxdb",
	})
	th.AssertNoErr(t, err)

	for _, flavor := range listResp.Flavors {
		th.AssertEquals(t, "influxdb", flavor.EngineName)
	}
	tools.PrintResource(t, listResp)
}

func TestGeminiListCloudNativeFlavors(t *testing.T) {
	client, err := clients.NewGeminiDBSpecClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list GeminiDB Influx flavors with cloud native storage")
	listResp, err := spec.ListFlavors(client, spec.ListFlavorsOpts{
		EngineName: "influxdb",
		Mode:       "CloudNativeCluster",
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestGeminiGetVersions(t *testing.T) {
	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list Gemini db flavors")
	listResp, err := spec.GetVersions(client, "cassandra")
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}
