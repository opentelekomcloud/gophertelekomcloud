package publicips_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const publicIPJSON = `{
	"id": "publicip-id",
	"status": "DOWN",
	"profile": {
		"order_id": "order-id",
		"product_id": "product-id",
		"region_id": "eu-de",
		"user_id": "user-id"
	},
	"type": "5_bgp",
	"public_ip_address": "161.17.17.12",
	"ip_version": 4,
	"private_ip_address": "192.168.10.5",
	"port_id": "port-id",
	"tenant_id": "project-id",
	"create_time": "2015-07-16 04:32:50",
	"bandwidth_id": "bandwidth-id",
	"bandwidth_size": 10,
	"bandwidth_share_type": "PER",
	"bandwidth_name": "bandwidth-test",
	"alias": "tom",
	"enterprise_project_id": "0",
	"public_border_group": "center",
	"allow_share_bandwidth_types": ["share"],
	"tags": ["key=value"]
}`
