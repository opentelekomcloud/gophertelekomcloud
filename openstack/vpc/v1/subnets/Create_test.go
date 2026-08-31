package subnets_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{
			"subnet": {
				"name": "subnet",
				"description": "description",
				"cidr": "192.168.20.0/24",
				"gateway_ip": "192.168.20.1",
				"ipv6_enable": true,
				"dhcp_enable": true,
				"primary_dns": "100.125.4.25",
				"secondary_dns": "100.125.129.199",
				"dnsList": ["100.125.4.25", "100.125.129.199"],
				"availability_zone": "eu-de-01",
				"vpc_id": "vpc-id",
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

	enabled := true
	actual, err := subnets.Create(serviceClient(), subnets.CreateOpts{
		Name:             "subnet",
		Description:      "description",
		CIDR:             "192.168.20.0/24",
		GatewayIP:        "192.168.20.1",
		EnableIpv6:       &enabled,
		EnableDHCP:       &enabled,
		PrimaryDNS:       "100.125.4.25",
		SecondaryDNS:     "100.125.129.199",
		DNSList:          []string{"100.125.4.25", "100.125.129.199"},
		AvailabilityZone: "eu-de-01",
		VpcID:            "vpc-id",
		ExtraDHCPOpts: []subnets.ExtraDHCPOpt{
			{OptName: "ntp", OptValue: "10.100.0.33,10.100.0.34"},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "subnet-id", actual.ID)
	th.AssertEquals(t, "subnet", actual.Name)
	th.AssertEquals(t, "192.168.20.0/24", actual.CIDR)
	th.AssertEquals(t, "192.168.20.1", actual.GatewayIP)
	th.AssertEquals(t, "vpc-id", actual.VpcID)
	th.AssertEquals(t, "ACTIVE", actual.Status)
	th.AssertEquals(t, true, actual.EnableIpv6)
	th.AssertEquals(t, "2407:c080:802:be7::/64", actual.CidrV6)
	th.AssertEquals(t, "eu-de-01", actual.AvailabilityZone)
	th.AssertEquals(t, "neutron-subnet-id", actual.SubnetID)
	th.AssertDeepEquals(t, []string{"100.125.4.25", "100.125.129.199"}, actual.DNSList)
	th.AssertDeepEquals(t, []subnets.ExtraDHCPOpt{
		{OptName: "ntp", OptValue: "10.100.0.33,10.100.0.34"},
	}, actual.ExtraDHCPOpts)
}

func TestCreateMinimalOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{
			"subnet": {
				"name": "subnet",
				"cidr": "192.168.20.0/24",
				"gateway_ip": "192.168.20.1",
				"vpc_id": "vpc-id"
			}
		}`)
		_, _ = w.Write([]byte(`{"subnet":` + subnetJSON + `}`))
	})

	actual, err := subnets.Create(serviceClient(), subnets.CreateOpts{
		Name:      "subnet",
		CIDR:      "192.168.20.0/24",
		GatewayIP: "192.168.20.1",
		VpcID:     "vpc-id",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "subnet-id", actual.ID)
}

func TestCreateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(_ http.ResponseWriter, _ *http.Request) {
		t.Fatal("no request expected for invalid options")
	})

	actual, err := subnets.Create(serviceClient(), subnets.CreateOpts{})
	if err == nil {
		t.Fatal("expected options validation error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestCreateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0202","error_msg":"Invalid subnet parameter values."}`))
	})

	actual, err := subnets.Create(serviceClient(), subnets.CreateOpts{
		Name:      "subnet",
		CIDR:      "192.168.20.0/24",
		GatewayIP: "192.168.20.1",
		VpcID:     "vpc-id",
	})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
