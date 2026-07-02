package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/versions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"versions":[{"id":"v1.0","links":[{"href":"https://eps.eu-de.otc.t-systems.com/v1.0/","rel":"self"}],"status":"CURRENT","updated":"2018-09-30T00:00:00Z"}]}`)
	})

	sc := client.ServiceClient()
	// ServiceClient endpoint is like http://localhost:PORT/
	// versions.List strips /v... so it will hit "/"
	sc.Endpoint = th.Endpoint()

	actual, err := versions.List(sc)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "v1.0", actual[0].ID)
	th.AssertEquals(t, "CURRENT", actual[0].Status)
	th.AssertEquals(t, "self", actual[0].Links[0].Rel)
}
