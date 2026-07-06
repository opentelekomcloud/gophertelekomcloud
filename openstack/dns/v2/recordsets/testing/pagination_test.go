package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/recordsets"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListByZoneMarkerPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	var requests int
	th.Mux.HandleFunc("/zones/zone-id/recordsets", func(w http.ResponseWriter, r *http.Request) {
		requests++
		th.TestMethod(t, r, "GET")
		th.CheckEquals(t, "2", r.URL.Query().Get("limit"))

		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Query().Get("marker") {
		case "": // first page: the API includes a "next" link
			_, _ = fmt.Fprintf(w, `
				{
					"recordsets": [{"id": "r1"}, {"id": "r2"}],
					"links": {
						"self": "%[1]szones/zone-id/recordsets?limit=2",
						"next": "%[1]szones/zone-id/recordsets?limit=2&marker=r2"
					}
				}
			`, th.Endpoint())
		case "r2": // marker pages: only a "self" link, no "next" (issue #952)
			_, _ = fmt.Fprintf(w, `
				{
					"recordsets": [{"id": "r3"}, {"id": "r4"}],
					"links": {"self": "%szones/zone-id/recordsets?limit=2&marker=r2"}
				}
			`, th.Endpoint())
		case "r4":
			_, _ = fmt.Fprintf(w, `
				{
					"recordsets": [{"id": "r5"}],
					"links": {"self": "%szones/zone-id/recordsets?limit=2&marker=r4"}
				}
			`, th.Endpoint())
		case "r5": // past the end: empty page terminates the iteration
			_, _ = fmt.Fprint(w, `{"recordsets": [], "links": {}}`)
		default:
			t.Errorf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	var ids []string
	pages := 0
	err := recordsets.ListByZone(client.ServiceClient(), "zone-id", recordsets.ListOpts{Limit: 2}).
		EachPage(func(page pagination.Page) (bool, error) {
			pages++
			s, err := recordsets.ExtractRecordSets(page)
			th.AssertNoErr(t, err)
			for _, rs := range s {
				ids = append(ids, rs.ID)
			}
			return true, nil
		})
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, []string{"r1", "r2", "r3", "r4", "r5"}, ids)
	th.CheckEquals(t, 3, pages)
	th.CheckEquals(t, 4, requests)
}
