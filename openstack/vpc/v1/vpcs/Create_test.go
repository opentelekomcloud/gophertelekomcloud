package vpcs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{
			"vpc": {
				"name": "vpc",
				"description": "description",
				"cidr": "192.168.0.0/16",
				"enterprise_project_id": "0"
			}
		}`)
		_, _ = w.Write([]byte(`{"vpc":` + vpcJSON + `}`))
	})

	actual, err := vpcs.Create(serviceClient(), vpcs.CreateOpts{
		Name:                "vpc",
		Description:         "description",
		CIDR:                "192.168.0.0/16",
		EnterpriseProjectID: "0",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.ID)
	th.AssertEquals(t, "vpc", actual.Name)
	th.AssertEquals(t, "192.168.0.0/16", actual.CIDR)
	th.AssertEquals(t, "OK", actual.Status)
	th.AssertDeepEquals(t, []vpcs.Route{{DestinationCIDR: "0.0.0.0/0", NextHop: "192.168.0.5"}}, actual.Routes)
}

func TestCreateEmptyOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"vpc":{}}`)
		_, _ = w.Write([]byte(`{"vpc":` + vpcJSON + `}`))
	})

	actual, err := vpcs.Create(serviceClient(), vpcs.CreateOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.ID)
}

func TestCreateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0101","error_msg":"Invalid VPC parameter values."}`))
	})

	actual, err := vpcs.Create(serviceClient(), vpcs.CreateOpts{})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
