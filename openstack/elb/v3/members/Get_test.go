package members_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id/members/member-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"member":` + memberJSON + `}`))
	})
	actual, err := members.Get(serviceClient(), "pool-id", "member-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "member-id", actual.ID)
	th.AssertEquals(t, "compute:eu-de-01", actual.DeviceOwner)
}

func TestGetInvalidIDs(t *testing.T) {
	actual, err := members.Get(serviceClient(), "../pools/pool-id", "member-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid pool ID error, got %#v, %v", actual, err)
	}
}
