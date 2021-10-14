package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/services"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/vpc-endpoint-services", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, createResponse)
	})

	iFalse := false
	opts := &services.CreateOpts{
		PortID:          "4189d3c2-8882-4871-a3c2-d380272eed88",
		RouterID:        "4189d3c2-8882-4871-a3c2-d380272eed80",
		ApprovalEnabled: &iFalse,
		ServiceType:     services.ServiceTypeInterface,
		ServerType:      services.ServerTypeVM,
		Ports: []services.PortMapping{
			{
				ClientPort: 8080,
				ServerPort: 90,
				Protocol:   "TCP",
			},
			{
				ClientPort: 8081,
				ServerPort: 80,
				Protocol:   "TCP",
			},
		},
	}

	created, err := services.Create(client.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.RouterID, opts.RouterID)
}

func TestGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	id := "4189d3c2-8882-4871-a3c2-d380272eed80"

	th.Mux.HandleFunc(fmt.Sprintf("/vpc-endpoint-services/%s", id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, createResponse)
	})

	svc, err := services.Get(client.ServiceClient(), id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, id, svc.RouterID)
}

func TestUpdateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	id := "4189d3c2-8882-4871-a3c2-d380272eed80"

	th.Mux.HandleFunc(fmt.Sprintf("/vpc-endpoint-services/%s", id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, updateResponse)
	})

	iTrue := true
	opts := &services.UpdateOpts{
		ServiceName:     "test",
		ApprovalEnabled: &iTrue,
		Ports: []services.PortMapping{
			{
				ClientPort: 8081,
				ServerPort: 22,
				Protocol:   "TCP",
			},
			{
				ClientPort: 8082,
				ServerPort: 23,
				Protocol:   "UDP",
			},
		},
	}

	updated, err := services.Update(client.ServiceClient(), id, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updated.ApprovalEnabled, *opts.ApprovalEnabled)
	th.AssertDeepEquals(t, updated.Ports, opts.Ports)
}

func TestListRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	name := "test123"

	th.Mux.HandleFunc(fmt.Sprintf("/vpc-endpoint-services"), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestFormValues(t, r, map[string]string{"endpoint_service_name": name})

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponse)
	})

	opts := services.ListOpts{
		Name: name,
	}
	pages, err := services.List(client.ServiceClient(), opts).AllPages()
	th.AssertNoErr(t, err)
	list, err := services.ExtractServices(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(list))
	th.AssertEquals(t, name, list[0].ServiceName)
}
