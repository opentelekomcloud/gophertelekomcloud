package testing

import (
	"io"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const volumeID = "d32019d3-bc6e-4319-9c1d-6722fc136a22"

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		opts     volumes.DeleteOpts
		rawQuery string
	}{
		{
			name: "cascade",
			opts: volumes.DeleteOpts{
				Cascade: true,
			},
			rawQuery: "cascade=true",
		},
		{
			name: "zero options",
			opts: volumes.DeleteOpts{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()

			th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, r *http.Request) {
				th.TestMethod(t, r, http.MethodDelete)
				th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
				th.AssertEquals(t, tt.rawQuery, r.URL.RawQuery)

				body, err := io.ReadAll(r.Body)
				th.AssertNoErr(t, err)
				th.AssertEquals(t, "", string(body))

				w.WriteHeader(http.StatusAccepted)
			})

			err := volumes.Delete(client.ServiceClient(), volumeID, tt.opts)
			th.AssertNoErr(t, err)
		})
	}
}

func TestDeleteRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := volumes.Delete(client.ServiceClient(), volumeID, volumes.DeleteOpts{})
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
}

func TestDeleteErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "delete failed", http.StatusInternalServerError)
	})

	err := volumes.Delete(client.ServiceClient(), volumeID, volumes.DeleteOpts{})
	if err == nil {
		t.Fatal("expected an error response")
	}
}
