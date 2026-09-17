package loadbalancers_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		requests++
		query := r.URL.Query()
		th.AssertEquals(t, "2", query.Get("limit"))
		th.AssertEquals(t, "true", query.Get("page_reverse"))
		th.AssertEquals(t, "false", query.Get("admin_state_up"))
		th.AssertEquals(t, "false", query.Get("guaranteed"))
		th.AssertEquals(t, "false", query.Get("deletion_protection_enable"))
		th.AssertEquals(t, "enterprise-project-id", query.Get("enterprise_project_id"))
		th.AssertEquals(t, "6", query.Get("ip_version"))
		th.AssertEquals(t, "protection", query.Get("protection_status"))
		switch query.Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"loadbalancers":[%s,%s],"page_info":{"next_marker":"lb-2","current_count":2}}`,
				strings.ReplaceAll(loadBalancerJSON, "loadbalancer-id", "lb-1"),
				strings.ReplaceAll(loadBalancerJSON, "loadbalancer-id", "lb-2"))
		case "lb-2":
			_, _ = fmt.Fprintf(w, `{"loadbalancers":[%s],"page_info":{"current_count":1}}`,
				strings.ReplaceAll(loadBalancerJSON, "loadbalancer-id", "lb-3"))
		default:
			t.Fatalf("unexpected marker %q", query.Get("marker"))
		}
	})
	actual, err := loadbalancers.List(serviceClient(), loadbalancers.ListOpts{
		Marker: "start", Limit: 2, PageReverse: pointerto.Bool(true),
		ID: []string{"loadbalancer-id"}, Name: []string{"loadbalancer-test"}, Description: []string{"load balancer"},
		AdminStateUp: pointerto.Bool(false), ProvisioningStatus: []string{"ACTIVE"}, OperatingStatus: []string{"ONLINE"},
		Guaranteed: pointerto.Bool(false), VpcID: []string{"vpc-id"}, VipPortID: []string{"port-id"},
		VipAddress: []string{"192.0.2.10"}, VipSubnetCidrID: []string{"subnet-id"},
		IpV6VipPortID: []string{"ipv6-port-id"}, IpV6VipAddress: []string{"2001:db8::10"},
		IpV6VipSubnetID: []string{"ipv6-subnet-id"}, Eips: []string{"eip-id"}, PublicIps: []string{"publicip-id"},
		AvailabilityZoneList: []string{"eu-de-01"}, L4FlavorID: []string{"l4-flavor-id"},
		L4ScaleFlavorID: []string{"l4-scale-flavor-id"}, L7FlavorID: []string{"l7-flavor-id"},
		L7ScaleFlavorID: []string{"l7-scale-flavor-id"}, MemberDeviceID: []string{"device-id"},
		MemberAddress: []string{"192.0.2.20"}, EnterpriseProjectID: []string{"enterprise-project-id"},
		IPVersion: []int{6}, DeletionProtectionEnable: pointerto.Bool(false),
		ElbSubnetType: []string{"dualstack"}, ProtectionStatus: []string{"protection"},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "global-eip-id", actual[0].GlobalEips[0].GlobalEipID)
	th.AssertEquals(t, 100, actual[0].CustomQosLimit.L4.Connection)
	th.AssertEquals(t, 3, requests)
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/loadbalancers", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"loadbalancers":`))
	})
	actual, err := loadbalancers.List(serviceClient(), loadbalancers.ListOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected extraction error, got %#v, %v", actual, err)
	}
}
