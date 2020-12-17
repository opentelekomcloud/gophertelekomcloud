package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func mockEvacuateResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `
		{
		    "evacuate": {
		    "adminPass": "MySecretPass",
		    "host": "derp",
		    "onSharedStorage": false
		  }

		}
		      `)
		w.WriteHeader(http.StatusOK)
	})
}

func mockEvacuateResponseWithHost(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `
		{
		    "evacuate": {
		    "host": "derp",
		    "onSharedStorage": false
		  }

		}
		      `)
		w.WriteHeader(http.StatusOK)
	})
}

func mockEvacuateResponseWithNoOpts(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `
		{
		    "evacuate": {
		    "onSharedStorage": false
		  }

		}
		      `)
		w.WriteHeader(http.StatusOK)
	})
}

const EvacuateResponse = `
{
  "adminPass": "MySecretPass"
}
`

func mockEvacuateAdminpassResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `
    {
        "evacuate": {
          "onSharedStorage": false
        }
    }
          `)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, EvacuateResponse)
	})
}
