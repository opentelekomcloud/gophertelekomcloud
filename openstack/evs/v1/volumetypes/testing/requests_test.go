package testing

import (
	"fmt"
	"net/http"
	"testing"

	volumetypes2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v1/volumetypes"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := volumetypes2.List(client.ServiceClient())
	if err != nil {
		t.Errorf("Failed to extract volume types: %v", err)
	}

	expected := []volumetypes2.VolumeType{
		{
			ID:   "289da7f8-6440-407c-9fb4-7db01ec49164",
			Name: "vol-type-001",
			ExtraSpecs: map[string]interface{}{
				"capabilities": "gpu",
			},
		},
		{
			ID:         "96c3bda7-c82a-4f50-be73-ca7621794835",
			Name:       "vol-type-002",
			ExtraSpecs: map[string]interface{}{},
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	vt, err := volumetypes2.Get(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, vt.ExtraSpecs, map[string]interface{}{"serverNumber": "2"})
	th.AssertEquals(t, vt.Name, "vol-type-001")
	th.AssertEquals(t, vt.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "volume_type": {
        "name": "vol-type-001"
    }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		_, _ = fmt.Fprint(w, `
{
    "volume_type": {
        "name": "vol-type-001",
        "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22"
    }
}
		`)
	})

	options := volumetypes2.CreateOpts{Name: "vol-type-001"}
	n, err := volumetypes2.Create(client.ServiceClient(), options)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Name, "vol-type-001")
	th.AssertEquals(t, n.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})

	err := volumetypes2.Delete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, err)
}
