package privateips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/privateips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips/private-ip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"privateip":` + privateIPJSON + `}`))
	})

	actual, err := privateips.Get(serviceClient(), "private-ip-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "private-ip-id", actual.ID)
	th.AssertEquals(t, "DOWN", actual.Status)
	th.AssertEquals(t, "network-id", actual.SubnetID)
	th.AssertEquals(t, "project-id", actual.TenantID)
	th.AssertEquals(t, "", actual.DeviceOwner)
	th.AssertEquals(t, "192.168.20.10", actual.IPAddress)
}

func TestGetZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips/private-ip-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"privateip":{"id":"private-ip-id"}}`))
	})

	actual, err := privateips.Get(serviceClient(), "private-ip-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "private-ip-id", actual.ID)
	th.AssertEquals(t, "", actual.Status)
	th.AssertEquals(t, "", actual.DeviceOwner)
	th.AssertEquals(t, "", actual.IPAddress)
}

func TestGetInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips/private-ip-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"privateip":`))
	})

	actual, err := privateips.Get(serviceClient(), "private-ip-id")
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestGetError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips/private-ip-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0206","error_msg":"Private IP does not exist."}`))
	})

	actual, err := privateips.Get(serviceClient(), "private-ip-id")
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestGetInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	actual, err := privateips.Get(serviceClient(), "../vpcs/vpc-id")
	if err == nil {
		t.Fatal("expected invalid ID error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
