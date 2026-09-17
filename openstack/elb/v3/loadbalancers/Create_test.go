package loadbalancers_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"loadbalancer":{"project_id":"project-id","name":"loadbalancer-test","description":"load balancer","vip_address":"192.0.2.10","vip_subnet_cidr_id":"subnet-id","ipv6_vip_virsubnet_id":"ipv6-subnet-id","provider":"vlb","l4_flavor_id":"l4-flavor-id","l7_flavor_id":"l7-flavor-id","guaranteed":true,"vpc_id":"vpc-id","availability_zone_list":["eu-de-01"],"enterprise_project_id":"enterprise-project-id","tags":[{"key":"key","value":"value"}],"admin_state_up":true,"ipv6_bandwidth":{"id":"bandwidth-id"},"publicip_ids":["publicip-id"],"publicip":{"ip_version":4,"network_type":"5_bgp","billing_info":"billing","description":"public IP","bandwidth":{"name":"bandwidth","size":10,"charge_mode":"traffic","share_type":"PER","billing_info":"bandwidth-billing","id":"bandwidth-id"}},"elb_virsubnet_ids":["network-id"],"ip_target_enable":true,"deletion_protection_enable":true,"waf_failure_action":"forward","charge_mode":"flavor","protection_status":"protection","protection_reason":"managed"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"loadbalancer":` + loadBalancerJSON + `}`))
	})
	actual, err := loadbalancers.Create(serviceClient(), loadbalancers.CreateOpts{
		ProjectID: "project-id", Name: "loadbalancer-test", Description: "load balancer",
		VipAddress: "192.0.2.10", VipSubnetCidrID: "subnet-id", IpV6VipSubnetID: "ipv6-subnet-id",
		Provider: "vlb", L4Flavor: "l4-flavor-id", L7Flavor: "l7-flavor-id", Guaranteed: pointerto.Bool(true),
		VpcID: "vpc-id", AvailabilityZoneList: []string{"eu-de-01"}, EnterpriseProjectID: "enterprise-project-id",
		Tags: []tags.ResourceTag{{Key: "key", Value: "value"}}, AdminStateUp: pointerto.Bool(true),
		IPV6Bandwidth: &loadbalancers.BandwidthRef{ID: "bandwidth-id"}, PublicIpIDs: []string{"publicip-id"},
		PublicIp: &loadbalancers.PublicIp{
			IpVersion: 4, NetworkType: "5_bgp", BillingInfo: "billing", Description: "public IP",
			Bandwidth: loadbalancers.Bandwidth{
				Name: "bandwidth", Size: 10, ChargeMode: "traffic", ShareType: "PER",
				BillingInfo: "bandwidth-billing", ID: "bandwidth-id",
			},
		},
		ElbSubnetIDs: []string{"network-id"}, IpTargetEnable: pointerto.Bool(true),
		DeletionProtectionEnable: pointerto.Bool(true), WafFailureAction: "forward",
		ChargeMode: "flavor", ProtectionStatus: "protection", ProtectionReason: "managed",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "loadbalancer-id", actual.ID)
	th.AssertEquals(t, "enterprise-project-id", actual.EnterpriseProjectID)
}

func TestCreateMissingAvailabilityZone(t *testing.T) {
	actual, err := loadbalancers.Create(serviceClient(), loadbalancers.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected validation error, got %#v, %v", actual, err)
	}
}
