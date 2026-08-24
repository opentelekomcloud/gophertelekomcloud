package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestUpdateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataResponse(t)

	actual, err := snapshots.UpdateMetadata(client.ServiceClient(), "123", snapshots.UpdateMetadataOpts{
		Metadata: map[string]string{
			"empty": "",
			"key":   "v1",
		},
	})
	th.AssertNoErr(t, err)

	expected := &map[string]string{
		"empty": "",
		"key":   "v1",
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestUpdateMetadataMissingMetadata(t *testing.T) {
	actual, err := snapshots.UpdateMetadata(client.ServiceClient(), "123", snapshots.UpdateMetadataOpts{})
	if err == nil {
		t.Fatal("expected an error for missing metadata")
	}
	if actual != nil {
		t.Fatalf("expected nil metadata, got %#v", actual)
	}
}

func TestUpdateMetadataErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots/123/metadata", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "metadata update failed", http.StatusInternalServerError)
	})

	actual, err := snapshots.UpdateMetadata(client.ServiceClient(), "123", snapshots.UpdateMetadataOpts{
		Metadata: map[string]string{"key": "v1"},
	})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil metadata, got %#v", actual)
	}
}

func TestUpdateMetadataInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots/123/metadata", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"metadata":`))
	})

	actual, err := snapshots.UpdateMetadata(client.ServiceClient(), "123", snapshots.UpdateMetadataOpts{
		Metadata: map[string]string{"key": "v1"},
	})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual == nil {
		t.Fatal("expected the allocated metadata map with an extraction error")
	}
}
