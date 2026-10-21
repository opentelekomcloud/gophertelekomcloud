package members_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/members"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		q := r.URL.Query()
		th.AssertEquals(t, "true", q.Get("page_reverse"))
		th.AssertEquals(t, "pool-id", q.Get("pool_id"))
		th.AssertEquals(t, "loadbalancer-id", q.Get("loadbalancer_id"))
		switch q.Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"members":[%s],"page_info":{"previous_marker":"member-1","next_marker":"member-2"}}`,
				strings.ReplaceAll(memberJSON, "member-id", "member-1"))
		case "member-1":
			_, _ = fmt.Fprintf(w, `{"members":[%s],"page_info":{}}`,
				strings.ReplaceAll(memberJSON, "member-id", "member-2"))
		default:
			t.Fatalf("unexpected marker %q", q.Get("marker"))
		}
	})
	actual, err := members.ListAll(serviceClient(), members.ListAllOpts{
		Marker: "start", Limit: 1, PageReverse: true, Name: []string{"member-test"}, Weight: []int{10},
		AdminStateUp: pointerto.Bool(true), SubnetCIDRID: []string{"subnet-id"}, Address: []string{"192.0.2.10"},
		ProtocolPort: []int{8080}, ID: []string{"member-id"}, OperatingStatus: []string{"ONLINE"},
		EnterpriseProjectID: []string{"enterprise-id"}, IPVersion: []string{"v4"}, PoolID: []string{"pool-id"},
		LoadbalancerID: []string{"loadbalancer-id"},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "loadbalancer-id", actual[0].LoadbalancerID)
}
