package subnets_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id/subnets/subnet-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{
			"subnet": {
				"name": "subnet-updated",
				"description": "updated",
				"ipv6_enable": true,
				"dhcp_enable": false,
				"primary_dns": "100.125.4.25",
				"secondary_dns": "100.125.129.199",
				"dnsList": ["100.125.4.25", "100.125.129.199"],
				"extra_dhcp_opts": [
					{
						"opt_name": "ntp",
						"opt_value": "10.100.0.33,10.100.0.34"
					}
				]
			}
		}`)
		_, _ = w.Write([]byte(`{"subnet":` + subnetJSON + `}`))
	})

	description := "updated"
	enabled := true
	disabled := false
	actual, err := subnets.Update(serviceClient(), "vpc-id", "subnet-id", subnets.UpdateOpts{
		Name:         "subnet-updated",
		Description:  &description,
		EnableIpv6:   &enabled,
		EnableDHCP:   &disabled,
		PrimaryDNS:   "100.125.4.25",
		SecondaryDNS: "100.125.129.199",
		DNSList:      []string{"100.125.4.25", "100.125.129.199"},
		ExtraDHCPOpts: []subnets.ExtraDHCPOpt{
			{OptName: "ntp", OptValue: "10.100.0.33,10.100.0.34"},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "subnet-id", actual.ID)
	th.AssertEquals(t, "ACTIVE", actual.Status)
}

func TestUpdateEmptyDescription(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id/subnets/subnet-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"subnet":{"name":"subnet","description":""}}`)
		_, _ = w.Write([]byte(`{"subnet":` + subnetJSON + `}`))
	})

	description := ""
	actual, err := subnets.Update(serviceClient(), "vpc-id", "subnet-id", subnets.UpdateOpts{
		Name:        "subnet",
		Description: &description,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "subnet-id", actual.ID)
}

func TestUpdateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id/subnets/subnet-id", func(_ http.ResponseWriter, _ *http.Request) {
		t.Fatal("no request expected for invalid options")
	})

	actual, err := subnets.Update(serviceClient(), "vpc-id", "subnet-id", subnets.UpdateOpts{})
	if err == nil {
		t.Fatal("expected options validation error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestUpdateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id/subnets/subnet-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0202","error_msg":"Invalid subnet parameter values."}`))
	})

	actual, err := subnets.Update(serviceClient(), "vpc-id", "subnet-id", subnets.UpdateOpts{Name: "subnet"})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
