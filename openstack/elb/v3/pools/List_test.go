package pools_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		q := r.URL.Query()
		th.AssertEquals(t, "true", q.Get("page_reverse"))
		th.AssertEquals(t, "false", q.Get("member_deletion_protection_enable"))
		th.AssertEquals(t, "4", q.Get("ip_version"))
		th.AssertEquals(t, "member-address", q.Get("member_address"))
		th.AssertEquals(t, "member-device-id", q.Get("member_device_id"))
		th.AssertEquals(t, "listener-id", q.Get("listener_id"))
		th.AssertEquals(t, "member-instance-id", q.Get("member_instance_id"))
		th.AssertEquals(t, "vpc-id", q.Get("vpc_id"))
		th.AssertEquals(t, "instance", q.Get("type"))
		th.AssertEquals(t, "PROTECTED", q.Get("protection_status"))
		th.AssertEquals(t, "sort-key", q.Get("sort_key"))
		th.AssertEquals(t, "desc", q.Get("sort_dir"))
		switch q.Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"pools":[%s],"page_info":{"previous_marker":"pool-1","next_marker":"pool-2"}}`,
				strings.ReplaceAll(poolJSON, "pool-id", "pool-1"))
		case "pool-1":
			_, _ = fmt.Fprintf(w, `{"pools":[%s],"page_info":{}}`,
				strings.ReplaceAll(poolJSON, "pool-id", "pool-2"))
		default:
			t.Fatalf("unexpected marker %q", q.Get("marker"))
		}
	})
	actual, err := pools.List(serviceClient(), pools.ListOpts{
		Marker: "start", Limit: 1, PageReverse: true,
		Description: []string{"pool"}, HealthMonitorID: []string{"monitor-id"},
		LBMethod: []string{"LEAST_CONNECTIONS"}, Protocol: []string{"HTTP"},
		AdminStateUp: pointerto.Bool(true), Name: []string{"pool-test"}, ID: []string{"pool-id"},
		LoadbalancerID: []string{"loadbalancer-id"}, EnterpriseProjectID: []string{"enterprise-project-id"},
		IPVersion: []int{4}, MemberAddress: []string{"member-address"},
		MemberDeviceID: []string{"member-device-id"}, MemberDeletionProtectionEnable: pointerto.Bool(false),
		ListenerID: []string{"listener-id"}, MemberInstanceID: []string{"member-instance-id"},
		VpcID: []string{"vpc-id"}, Type: []string{"instance"}, ProtectionStatus: []string{"PROTECTED"},
		SortKey: "sort-key", SortDir: "desc",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "pool-1", actual[0].ID)
	th.AssertEquals(t, "2026-09-21T08:00:00Z", actual[0].CreatedAt)
	th.AssertEquals(t, "LOCALLY_REPLICA", actual[0].AZAffinity.AZUnhealthyFallbackStrategy)
}
