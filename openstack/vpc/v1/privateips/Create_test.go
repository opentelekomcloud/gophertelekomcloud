package privateips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/privateips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{
			"privateips": [
				{
					"subnet_id": "network-id"
				},
				{
					"subnet_id": "network-id",
					"ip_address": "192.168.20.11"
				}
			]
		}`)
		_, _ = w.Write([]byte(`{"privateips":[` + privateIPJSON + `]}`))
	})

	actual, err := privateips.Create(serviceClient(), privateips.CreateOpts{
		PrivateIPs: []privateips.PrivateIPRequest{
			{SubnetID: "network-id"},
			{SubnetID: "network-id", IPAddress: "192.168.20.11"},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "private-ip-id", actual[0].ID)
	th.AssertEquals(t, "DOWN", actual[0].Status)
	th.AssertEquals(t, "network-id", actual[0].SubnetID)
	th.AssertEquals(t, "project-id", actual[0].TenantID)
	th.AssertEquals(t, "", actual[0].DeviceOwner)
	th.AssertEquals(t, "192.168.20.10", actual[0].IPAddress)
}

func TestCreateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips", func(_ http.ResponseWriter, _ *http.Request) {
		t.Fatal("no request expected for invalid options")
	})

	actual, err := privateips.Create(serviceClient(), privateips.CreateOpts{})
	if err == nil {
		t.Fatal("expected options validation error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestCreateInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"privateips":`))
	})

	actual, err := privateips.Create(serviceClient(), privateips.CreateOpts{
		PrivateIPs: []privateips.PrivateIPRequest{{SubnetID: "network-id"}},
	})
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestCreateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0207","error_msg":"Invalid private IP parameter."}`))
	})

	actual, err := privateips.Create(serviceClient(), privateips.CreateOpts{
		PrivateIPs: []privateips.PrivateIPRequest{{SubnetID: "network-id"}},
	})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
