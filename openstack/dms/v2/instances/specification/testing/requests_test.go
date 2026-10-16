package testing

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/specification"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "54602a9d-5e22-4239-9123-77e350df4a34"
	jobID      = "93b94287-728d-4bb1-a158-cb66cb0854e7"
	extendPath = "/v2/kafka/" + projectID + "/instances/" + instanceID + "/extend"
)

// dmsClient mirrors the real dmsv2 catalog endpoint: https://dms.{region}.otc.t-systems.com/v2/{project_id}/
func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func handleExtend(t *testing.T, path, expectedBody string) {
	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, expectedBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, `{"job_id": %q}`, jobID)
	})
}

func TestIncreaseSpecStorage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleExtend(t, extendPath, `{"oper_type": "storage", "new_storage_space": 600}`)

	res, err := specification.IncreaseSpec(dmsClient(), instanceID, specification.IncreaseSpecOpts{
		Engine:          "kafka",
		OperType:        "storage",
		NewStorageSpace: pointerto.Int(600),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, jobID, res.JobId)
}

func TestIncreaseSpecHorizontal(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleExtend(t, extendPath, `{
		"oper_type": "horizontal",
		"new_storage_space": 1600,
		"new_broker_num": 4,
		"publicip_id": "eip-1,eip-2",
		"tenant_ips": ["127.0.0.1", "127.0.0.2", "127.0.0.3"],
		"second_tenant_subnet_id": "subnet-2"
	}`)

	res, err := specification.IncreaseSpec(dmsClient(), instanceID, specification.IncreaseSpecOpts{
		Engine:               "kafka",
		OperType:             "horizontal",
		NewStorageSpace:      pointerto.Int(1600),
		NewBrokerNumber:      pointerto.Int(4),
		PublicIpId:           "eip-1,eip-2",
		TenantIps:            []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"},
		SecondTenantSubnetId: "subnet-2",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, jobID, res.JobId)
}

func TestIncreaseSpecVertical(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// Unset integer fields must be omitted entirely.
	handleExtend(t, extendPath, `{"oper_type": "vertical", "new_product_id": "c6.4u8g.cluster"}`)

	res, err := specification.IncreaseSpec(dmsClient(), instanceID, specification.IncreaseSpecOpts{
		Engine:       "kafka",
		OperType:     "vertical",
		NewProductId: "c6.4u8g.cluster",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, jobID, res.JobId)
}

func TestIncreaseSpecExplicitZeroIsSent(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// Explicit zeros must reach the API so that server-side validation applies.
	handleExtend(t, extendPath, `{"oper_type": "horizontal", "new_storage_space": 0, "new_broker_num": 0}`)

	_, err := specification.IncreaseSpec(dmsClient(), instanceID, specification.IncreaseSpecOpts{
		Engine:          "kafka",
		OperType:        "horizontal",
		NewStorageSpace: pointerto.Int(0),
		NewBrokerNumber: pointerto.Int(0),
	})
	th.AssertNoErr(t, err)
}

func TestIncreaseSpecValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	valid := specification.IncreaseSpecOpts{
		Engine:          "kafka",
		OperType:        "storage",
		NewStorageSpace: pointerto.Int(600),
	}
	noV2Client := client.ServiceClient()
	noV2Client.Endpoint = th.Endpoint() + projectID + "/"

	cases := map[string]struct {
		client     *golangsdk.ServiceClient
		instanceID string
		opts       specification.IncreaseSpecOpts
		errContain string
	}{
		"missing instance id": {dmsClient(), "", valid, "instanceId"},
		"missing engine": {dmsClient(), instanceID, specification.IncreaseSpecOpts{
			OperType: "storage", NewStorageSpace: pointerto.Int(600),
		}, "Engine"},
		"missing oper_type": {dmsClient(), instanceID, specification.IncreaseSpecOpts{
			Engine: "kafka", NewStorageSpace: pointerto.Int(600),
		}, "OperType"},
		"endpoint without v2": {noV2Client, instanceID, valid, "no v2 segment"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var (
				res *specification.IncreaseSpecResp
				err error
			)
			func() {
				defer func() {
					if p := recover(); p != nil {
						t.Fatalf("IncreaseSpec panicked: %v", p)
					}
				}()
				res, err = specification.IncreaseSpec(tc.client, tc.instanceID, tc.opts)
			}()
			if err == nil {
				t.Fatalf("expected error, got response %+v", res)
			}
			if !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %q", tc.errContain, err)
			}
		})
	}
}

func TestIncreaseSpecAPIError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(extendPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error_code": "DMS.00400026", "error_msg": "This operation is not allowed due to the instance status."}`)
	})

	res, err := specification.IncreaseSpec(dmsClient(), instanceID, specification.IncreaseSpecOpts{
		Engine:       "kafka",
		OperType:     "vertical",
		NewProductId: "c6.4u8g.cluster",
	})
	if res != nil {
		t.Fatalf("expected nil response on error, got %+v", res)
	}
	var httpErr golangsdk.ErrDefault400
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault400, got %T: %v", err, err)
	}
	th.AssertEquals(t, true, strings.Contains(string(httpErr.Body), "DMS.00400026"))
}
