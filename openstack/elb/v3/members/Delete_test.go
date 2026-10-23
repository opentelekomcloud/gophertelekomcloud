package members_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id/members/member-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, members.Delete(serviceClient(), "pool-id", "member-id"))
}

func TestDeleteInvalidIDs(t *testing.T) {
	if err := members.Delete(serviceClient(), "../pools/pool-id", "member-id"); err == nil {
		t.Fatal("expected invalid pool ID error")
	}
}
