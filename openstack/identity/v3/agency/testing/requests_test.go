package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/agency"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	mock "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestAgencyList(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(th.TeardownHTTP)

	th.Mux.HandleFunc("/OS-AGENCY/agencies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", mock.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, agencyListResponse)
	})

	opts := agency.ListOpts{}
	pages, err := agency.List(mock.ServiceClient(), opts).AllPages()
	th.AssertNoErr(t, err)

	agencies, err := agency.ExtractAgencies(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(agencies))
	th.AssertEquals(t, "afca8ddf2e92469a8fd26a635da5206f", agencies[0].ID)
}
