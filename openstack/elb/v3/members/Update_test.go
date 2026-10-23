package members_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id/members/member-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"member":{"name":"","weight":0,"admin_state_up":true,"availability_zone":"eu-de-01"}}`)
		_, _ = w.Write([]byte(`{"member":` + memberJSON + `}`))
	})
	empty := ""
	actual, err := members.Update(serviceClient(), "pool-id", "member-id", members.UpdateOpts{
		Name: &empty, Weight: pointerto.Int(0), AdminStateUp: pointerto.Bool(true), AvailabilityZone: "eu-de-01",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "member-id", actual.ID)
}

func TestUpdateInvalidIDs(t *testing.T) {
	actual, err := members.Update(serviceClient(), "pool-id", "../members/member-id", members.UpdateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid member ID error, got %#v, %v", actual, err)
	}
}
