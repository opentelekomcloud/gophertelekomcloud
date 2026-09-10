package securitygrouprules_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const ruleJSON = `{
	"id":"rule-id",
	"description":"allow web",
	"security_group_id":"group-id",
	"direction":"ingress",
	"ethertype":"IPv4",
	"protocol":"tcp",
	"port_range_min":80,
	"port_range_max":80,
	"remote_ip_prefix":"10.0.0.0/24",
	"remote_group_id":null,
	"remote_address_group_id":null,
	"tenant_id":"project-id"
}`
