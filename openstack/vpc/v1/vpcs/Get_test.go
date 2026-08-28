package vpcs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"vpc":` + vpcJSON + `}`))
	})

	actual, err := vpcs.Get(serviceClient(), "vpc-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.ID)
	th.AssertEquals(t, "description", actual.Description)
	th.AssertEquals(t, "0", actual.EnterpriseProjectID)
	th.AssertEquals(t, false, actual.EnableSharedSnat)
}

func TestGetZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"vpc":{"id":"vpc-id"}}`))
	})

	actual, err := vpcs.Get(serviceClient(), "vpc-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.ID)
	th.AssertEquals(t, "", actual.Name)
	th.AssertEquals(t, "", actual.CIDR)
	th.AssertEquals(t, 0, len(actual.Routes))
}

func TestGetNotFound(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	actual, err := vpcs.Get(serviceClient(), "vpc-id")
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
