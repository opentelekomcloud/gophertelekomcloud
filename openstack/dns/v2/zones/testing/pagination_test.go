package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListMarkerPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	var requests int
	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		requests++
		th.TestMethod(t, r, "GET")
		th.CheckEquals(t, "2", r.URL.Query().Get("limit"))

		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Query().Get("marker") {
		case "": // first page: the API includes a "next" link
			_, _ = fmt.Fprintf(w, `
				{
					"zones": [{"id": "z1"}, {"id": "z2"}],
					"links": {
						"self": "%[1]szones?limit=2",
						"next": "%[1]szones?limit=2&marker=z2"
					}
				}
			`, th.Endpoint())
		case "z2": // marker pages: only a "self" link, no "next"
			_, _ = fmt.Fprintf(w, `
				{
					"zones": [{"id": "z3"}],
					"links": {"self": "%szones?limit=2&marker=z2"}
				}
			`, th.Endpoint())
		case "z3": // past the end: empty page terminates the iteration
			_, _ = fmt.Fprint(w, `{"zones": [], "links": {}}`)
		default:
			t.Errorf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	var ids []string
	pages := 0
	err := zones.List(client.ServiceClient(), zones.ListOpts{Limit: 2}).
		EachPage(func(page pagination.Page) (bool, error) {
			pages++
			s, err := zones.ExtractZones(page)
			th.AssertNoErr(t, err)
			for _, z := range s {
				ids = append(ids, z.ID)
			}
			return true, nil
		})
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, []string{"z1", "z2", "z3"}, ids)
	th.CheckEquals(t, 2, pages)
	th.CheckEquals(t, 3, requests)
}
