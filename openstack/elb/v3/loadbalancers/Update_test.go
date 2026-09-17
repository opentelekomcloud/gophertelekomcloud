package loadbalancers_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/loadbalancers/loadbalancer-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"loadbalancer":{"name":"updated","admin_state_up":true,"description":"","ipv6_vip_virsubnet_id":"ipv6-subnet-id","vip_subnet_cidr_id":"subnet-id","vip_address":"192.0.2.20","l4_flavor_id":"l4-flavor-id","l7_flavor_id":"l7-flavor-id","ipv6_bandwidth":{"id":"bandwidth-id"},"ip_target_enable":true,"elb_virsubnet_ids":["network-id"],"deletion_protection_enable":true,"waf_failure_action":"forward","protection_status":"protection","protection_reason":"managed"}}`)
		_, _ = w.Write([]byte(`{"loadbalancer":` + loadBalancerJSON + `}`))
	})
	actual, err := loadbalancers.Update(serviceClient(), "loadbalancer-id", loadbalancers.UpdateOpts{
		Name: "updated", AdminStateUp: pointerto.Bool(true), Description: pointerto.String(""),
		IpV6VipSubnetID: pointerto.String("ipv6-subnet-id"), VipSubnetCidrID: pointerto.String("subnet-id"),
		VipAddress: "192.0.2.20", L4Flavor: "l4-flavor-id", L7Flavor: "l7-flavor-id",
		IpV6Bandwidth: &loadbalancers.BandwidthRef{ID: "bandwidth-id"}, IpTargetEnable: pointerto.Bool(true),
		ElbSubnetIDs: []string{"network-id"}, DeletionProtectionEnable: pointerto.Bool(true),
		WafFailureAction: "forward", ProtectionStatus: "protection", ProtectionReason: "managed",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "loadbalancer-id", actual.ID)
}

func TestUpdateInvalidID(t *testing.T) {
	actual, err := loadbalancers.Update(serviceClient(), "../listeners/listener-id", loadbalancers.UpdateOpts{Name: "updated"})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
