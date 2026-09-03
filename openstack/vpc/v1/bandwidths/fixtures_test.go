package bandwidths_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const bandwidthJSON = `{
	"id": "bandwidth-id",
	"name": "bandwidth-6f78",
	"size": 5,
	"share_type": "PER",
	"public_border_group": "center",
	"publicip_info": [
		{
			"publicip_id": "publicip-id",
			"publicip_address": "10.10.10.62",
			"ip_version": 4,
			"publicip_type": "5_bgp"
		}
	],
	"tenant_id": "project-id",
	"bandwidth_type": "bgp",
	"charge_mode": "traffic",
	"enterprise_project_id": "0",
	"status": "NORMAL",
	"created_at": "2020-04-21T07:58:02Z",
	"updated_at": "2020-04-21T07:58:02Z"
}`
