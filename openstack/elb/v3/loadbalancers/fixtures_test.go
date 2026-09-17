package loadbalancers_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

const loadBalancerJSON = `{
	"id":"loadbalancer-id",
	"name":"loadbalancer-test",
	"description":"load balancer",
	"provisioning_status":"ACTIVE",
	"admin_state_up":true,
	"provider":"vlb",
	"pools":[{"id":"pool-id"}],
	"listeners":[{"id":"listener-id"}],
	"operating_status":"ONLINE",
	"project_id":"project-id",
	"vip_subnet_cidr_id":"subnet-id",
	"vip_address":"192.0.2.10",
	"vip_port_id":"port-id",
	"tags":[{"key":"key","value":"value"}],
	"created_at":"2026-09-17T09:00:00Z",
	"updated_at":"2026-09-17T09:01:00Z",
	"guaranteed":true,
	"vpc_id":"vpc-id",
	"eips":[{"eip_id":"eip-id","eip_address":"198.51.100.10","ip_version":4}],
	"ipv6_vip_address":"2001:db8::10",
	"ipv6_vip_virsubnet_id":"ipv6-subnet-id",
	"ipv6_vip_port_id":"ipv6-port-id",
	"availability_zone_list":["eu-de-01"],
	"enterprise_project_id":"enterprise-project-id",
	"billing_info":"billing",
	"l4_flavor_id":"l4-flavor-id",
	"l4_scale_flavor_id":"l4-scale-flavor-id",
	"l7_flavor_id":"l7-flavor-id",
	"l7_scale_flavor_id":"l7-scale-flavor-id",
	"publicips":[{"publicip_id":"publicip-id","publicip_address":"203.0.113.10","ip_version":4}],
	"global_eips":[{"global_eip_id":"global-eip-id","global_eip_address":"203.0.113.20","ip_version":4}],
	"elb_virsubnet_ids":["network-id"],
	"elb_virsubnet_type":"dualstack",
	"ip_target_enable":true,
	"frozen_scene":"",
	"ipv6_bandwidth":{"id":"bandwidth-id"},
	"deletion_protection_enable":true,
	"autoscaling":{"enable":true,"min_l7_flavor_id":"min-l7-flavor-id"},
	"public_border_group":"center",
	"waf_failure_action":"forward",
	"charge_mode":"flavor",
	"protection_status":"protection",
	"protection_reason":"managed",
	"loadbalancer_type":"gateway",
	"gw_flavor_id":"gw-flavor-id",
	"instance_type":"service",
	"instance_id":"instance-id",
	"log_group_id":"log-group-id",
	"log_topic_id":"log-topic-id",
	"custom_qos_limit":{"l4":{"connection":100,"cps":10},"l7":{"connection":200,"cps":20}},
	"service_lb_mode":"standard",
	"proxy_protocol_extensions":[{"vip_address":"192.0.2.10","ipv6_vip_address":"2001:db8::10","extension":{"ep_id":"endpoint-id","ep_service_id":"endpoint-service-id"}}]
}`
