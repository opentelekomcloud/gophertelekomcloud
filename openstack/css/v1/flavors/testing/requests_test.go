package testing

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

var emptyFlavor = flavors.Flavor{}

func TestCSSClusterFlavorsListResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, listResponseBody)
	})

	client := fake.ServiceClient()

	pages, err := flavors.List(client).AllPages()
	th.AssertNoErr(t, err)

	empty, err := pages.IsEmpty()
	th.AssertNoErr(t, err)
	if empty {
		t.Fatal("version list is empty")
	}

	versions, _ := flavors.ExtractVersions(pages)
	for _, version := range versions {
		if version.Version == "" {
			t.Error("version object has no object")
		}
		if version.Type == "" {
			t.Error("version object has no type")
		}
		for _, flavor := range version.Flavors {
			if reflect.DeepEqual(emptyFlavor, flavor) {
				t.Error("flavor is empty")
			}
		}
	}
}
