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

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		q := r.URL.Query()
		th.AssertEquals(t, "true", q.Get("page_reverse"))
		th.AssertEquals(t, "v4", q.Get("ip_version"))
		th.AssertEquals(t, "instance", q.Get("member_type"))
		th.AssertEquals(t, "instance-id", q.Get("instance_id"))
		th.AssertEquals(t, "eu-de-01", q.Get("availability_zone"))
		th.AssertEquals(t, "subnet-id", q.Get("subnet_cidr_id"))
		th.AssertEquals(t, "legacy-subnet-id", q.Get("subnet_sidr_id"))
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
	actual, err := members.List(serviceClient(), "pool-id", members.ListOpts{
		Name: []string{"member-test"}, Weight: []int{10}, AdminStateUp: pointerto.Bool(true),
		Address: []string{"192.0.2.10"}, ProtocolPort: []int{8080}, ID: []string{"member-id"},
		OperatingStatus: []string{"ONLINE"}, EnterpriseProjectID: []string{"enterprise-id"},
		Limit: 1, Marker: "start", PageReverse: true, IPVersion: []string{"v4"},
		MemberType: []string{"instance"}, InstanceID: []string{"instance-id"},
		AvailabilityZone: []string{"eu-de-01"}, SubnetCIDRID: []string{"subnet-id"},
		SubnetID: []string{"legacy-subnet-id"}, SortKey: "weight", SortDir: "desc",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "member-1", actual[0].ID)
	th.AssertEquals(t, "instance-id", actual[0].InstanceID)
}
