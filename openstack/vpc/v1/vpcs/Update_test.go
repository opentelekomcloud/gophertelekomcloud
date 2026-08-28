package vpcs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{
			"vpc": {
				"name": "vpc-updated",
				"description": "updated",
				"cidr": "192.168.0.0/16",
				"routes": [
					{
						"destination": "0.0.0.0/0",
						"nexthop": "192.168.0.5"
					}
				],
				"enable_shared_snat": true
			}
		}`)
		_, _ = w.Write([]byte(`{"vpc":` + vpcJSON + `}`))
	})

	description := "updated"
	sharedSnat := true
	actual, err := vpcs.Update(serviceClient(), "vpc-id", vpcs.UpdateOpts{
		Name:             "vpc-updated",
		Description:      &description,
		CIDR:             "192.168.0.0/16",
		Routes:           []vpcs.Route{{DestinationCIDR: "0.0.0.0/0", NextHop: "192.168.0.5"}},
		EnableSharedSnat: &sharedSnat,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.ID)
}

func TestUpdateEmptyDescription(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"vpc":{"description":""}}`)
		_, _ = w.Write([]byte(`{"vpc":` + vpcJSON + `}`))
	})

	description := ""
	actual, err := vpcs.Update(serviceClient(), "vpc-id", vpcs.UpdateOpts{Description: &description})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.ID)
}

func TestUpdateDisableSharedSnat(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"vpc":{"enable_shared_snat":false}}`)
		_, _ = w.Write([]byte(`{"vpc":` + vpcJSON + `}`))
	})

	sharedSnat := false
	actual, err := vpcs.Update(serviceClient(), "vpc-id", vpcs.UpdateOpts{EnableSharedSnat: &sharedSnat})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, actual.EnableSharedSnat)
}
