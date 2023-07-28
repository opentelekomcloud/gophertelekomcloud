package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/endpoints"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/vpc-endpoints", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, createResponse)
	})

	opts := endpoints.CreateOpts{
		NetworkID: "68bfbcc1-dff2-47e4-a9d4-332b9bc1b8de",
		ServiceID: expected.ServiceID,
		RouterID:  expected.RouterID,
		EnableDNS: true,
		Tags:      []tags.ResourceTag{{Key: "test1", Value: "test1"}},
	}

	ep, err := endpoints.Create(client.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expected, ep)
}

func TestGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	id := "4189d3c2-8882-4871-a3c2-d380272eed83"
	th.Mux.HandleFunc(fmt.Sprintf("/vpc-endpoints/%s", id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, createResponse)
	})

	ep, err := endpoints.Get(client.ServiceClient(), id).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expected, ep)
}

func TestListRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/vpc-endpoints", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponse)
	})

	opts := endpoints.ListOpts{
		RouterID: "84758cf5-9c62-43ae-a778-3dbd8370c0a4",
	}
	pages, err := endpoints.List(client.ServiceClient(), opts).AllPages()
	th.AssertNoErr(t, err)

	eps, err := endpoints.ExtractEndpoints(pages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, 2, len(eps))
}

func TestDeleteRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	id := "4189d3c2-8882-4871-a3c2-d380272eed83"
	th.Mux.HandleFunc(fmt.Sprintf("/vpc-endpoints/%s", id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, endpoints.Delete(client.ServiceClient(), id).Err)
}
