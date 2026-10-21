package members_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"member":{"address":"192.0.2.10","protocol_port":8080,"name":"member-test","project_id":"project-id","weight":10,"subnet_cidr_id":"subnet-id","admin_state_up":true,"availability_zone":"eu-de-01"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"member":` + memberJSON + `}`))
	})
	actual, err := members.Create(serviceClient(), "pool-id", members.CreateOpts{
		Address: "192.0.2.10", ProtocolPort: 8080, Name: "member-test", ProjectID: "project-id", Weight: pointerto.Int(10),
		SubnetID: "subnet-id", AdminStateUp: pointerto.Bool(true), AvailabilityZone: "eu-de-01",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "member-id", actual.ID)
}

func TestCreateInvalidPoolID(t *testing.T) {
	actual, err := members.Create(serviceClient(), "../pools/pool-id", members.CreateOpts{Address: "192.0.2.10", ProtocolPort: 8080})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid pool ID error, got %#v, %v", actual, err)
	}
}
