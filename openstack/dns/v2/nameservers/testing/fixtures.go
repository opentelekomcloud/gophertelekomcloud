package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

// ListOutput is a sample response to a List call.
const ListOutput string = `
{
    "nameservers": [
        {
            "hostname": "ns1.example.com.",
            "priority": 1
        },
        {
            "hostname": "ns2.example.com.",
            "priority": 2
        }
    ]
}
`

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/2c9eb155587194ec01587224c9f90149/nameservers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// nolint
		_, _ = fmt.Fprintf(w, ListOutput)
	})
}
