package testing

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumetypes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const volumeTypeID = "d32019d3-bc6e-4319-9c1d-6722fc136a22"

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		body, err := io.ReadAll(r.Body)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, "", string(body))

		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, `{"invalid":`)
	})

	err := volumetypes.Delete(client.ServiceClient(), volumeTypeID)
	th.AssertNoErr(t, err)
}

func TestDeleteRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := volumetypes.Delete(client.ServiceClient(), volumeTypeID)
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
}

func TestDeleteErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "delete failed", http.StatusInternalServerError)
	})

	err := volumetypes.Delete(client.ServiceClient(), volumeTypeID)
	if err == nil {
		t.Fatal("expected an error response")
	}
}
