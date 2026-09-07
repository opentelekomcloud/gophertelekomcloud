package quotas_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const quotasJSON = `{
	"quotas": {
		"resources": [
			{
				"type": "publicIp",
				"used": 2,
				"quota": 10,
				"min": 0
			},
			{
				"type": "vpc",
				"used": 0,
				"quota": -1,
				"min": 0
			}
		]
	}
}`
