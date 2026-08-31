package subnets_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/subnet-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"subnet":` + subnetJSON + `}`))
	})

	actual, err := subnets.Get(serviceClient(), "subnet-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "subnet-id", actual.ID)
	th.AssertEquals(t, "subnet", actual.Name)
	th.AssertEquals(t, "vpc-id", actual.VpcID)
	th.AssertEquals(t, "network-id", actual.NetworkID)
	th.AssertEquals(t, "center", actual.Scope)
	th.AssertEquals(t, "project-id", actual.TenantID)
}

func TestGetZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/subnet-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"subnet":{"id":"subnet-id","dhcp_enable":false,"ipv6_enable":false}}`))
	})

	actual, err := subnets.Get(serviceClient(), "subnet-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "subnet-id", actual.ID)
	th.AssertEquals(t, false, actual.EnableDHCP)
	th.AssertEquals(t, false, actual.EnableIpv6)
	th.AssertEquals(t, "", actual.CidrV6)
	if actual.DNSList != nil {
		t.Fatalf("expected nil DNSList, got %#v", actual.DNSList)
	}
	if actual.ExtraDHCPOpts != nil {
		t.Fatalf("expected nil ExtraDHCPOpts, got %#v", actual.ExtraDHCPOpts)
	}
}

func TestGetError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/subnet-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0202","error_msg":"Subnet does not exist."}`))
	})

	actual, err := subnets.Get(serviceClient(), "subnet-id")
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
