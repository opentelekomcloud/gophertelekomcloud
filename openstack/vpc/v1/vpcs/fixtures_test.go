package vpcs_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const vpcJSON = `{
	"id": "vpc-id",
	"name": "vpc",
	"description": "description",
	"cidr": "192.168.0.0/16",
	"status": "OK",
	"routes": [
		{
			"destination": "0.0.0.0/0",
			"nexthop": "192.168.0.5"
		}
	],
	"enterprise_project_id": "0",
	"enable_shared_snat": false
}`
