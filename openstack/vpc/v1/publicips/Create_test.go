package publicips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{
			"publicip": {
				"type": "5_bgp",
				"ip_address": "161.17.17.12",
				"ip_version": 4,
				"alias": "tom"
			},
			"bandwidth": {
				"name": "bandwidth-test",
				"size": 10,
				"share_type": "PER",
				"charge_mode": "traffic"
			},
			"enterprise_project_id": "0"
		}`)
		_, _ = w.Write([]byte(`{"publicip":` + publicIPJSON + `}`))
	})

	actual, err := publicips.Create(serviceClient(), publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type:      "5_bgp",
			IpAddress: "161.17.17.12",
			IPVersion: 4,
			Alias:     "tom",
		},
		Bandwidth: publicips.BandWidth{
			Name:       "bandwidth-test",
			Size:       10,
			ShareType:  "PER",
			ChargeMode: "traffic",
		},
		EnterpriseProjectId: "0",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
	th.AssertEquals(t, "DOWN", actual.Status)
	th.AssertEquals(t, "5_bgp", actual.Type)
	th.AssertEquals(t, "161.17.17.12", actual.PublicIpAddress)
	th.AssertEquals(t, 4, actual.IPVersion)
	th.AssertEquals(t, 10, actual.BandwidthSize)
	th.AssertEquals(t, "tom", actual.Alias)
	th.AssertEquals(t, "0", actual.EnterpriseProjectId)
	th.AssertEquals(t, "center", actual.PublicBorderGroup)
}

func TestCreateSharedBandwidth(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{
			"publicip": {
				"type": "5_bgp"
			},
			"bandwidth": {
				"name": "bandwidth-test",
				"size": 10,
				"id": "bandwidth-id",
				"share_type": "WHOLE"
			}
		}`)
		_, _ = w.Write([]byte(`{"publicip":` + publicIPJSON + `}`))
	})

	actual, err := publicips.Create(serviceClient(), publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{Type: "5_bgp"},
		Bandwidth: publicips.BandWidth{
			Name:      "bandwidth-test",
			Size:      10,
			ID:        "bandwidth-id",
			ShareType: "WHOLE",
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
}

func TestCreateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(_ http.ResponseWriter, _ *http.Request) {
		t.Fatal("no request expected for invalid options")
	})

	actual, err := publicips.Create(serviceClient(), publicips.CreateOpts{})
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

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0601","error_msg":"Invalid EIP parameter values."}`))
	})

	actual, err := publicips.Create(serviceClient(), publicips.CreateOpts{
		Publicip:  publicips.PublicIPRequest{Type: "5_bgp"},
		Bandwidth: publicips.BandWidth{Name: "bandwidth-test", Size: 10, ShareType: "PER"},
	})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
