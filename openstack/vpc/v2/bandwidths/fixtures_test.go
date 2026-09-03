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

const sharedBandwidthJSON = `{
	"id": "bandwidth-id",
	"name": "bandwidth123",
	"size": 10,
	"share_type": "WHOLE",
	"publicip_info": [
		{
			"publicip_id": "publicip-id",
			"publicip_address": "99.10.10.82",
			"publicip_type": "5_bgp",
			"ip_version": 4
		}
	],
	"tenant_id": "project-id",
	"charge_mode": "traffic",
	"billing_info": "",
	"bandwidth_type": "share",
	"status": "NORMAL"
}`
