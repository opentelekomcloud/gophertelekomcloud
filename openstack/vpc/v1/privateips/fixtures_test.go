package privateips_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const privateIPJSON = `{
	"status": "DOWN",
	"id": "private-ip-id",
	"subnet_id": "network-id",
	"tenant_id": "project-id",
	"device_owner": "",
	"ip_address": "192.168.20.10"
}`
