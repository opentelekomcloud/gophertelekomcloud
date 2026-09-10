package securitygroups_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const groupJSON = `{
	"name":"sg-test",
	"description":"default",
	"id":"group-id",
	"vpc_id":"vpc-id",
	"enterprise_project_id":"0",
	"security_group_rules":[{
		"id":"rule-id",
		"description":"",
		"security_group_id":"group-id",
		"direction":"egress",
		"ethertype":"IPv4",
		"protocol":null,
		"port_range_min":null,
		"port_range_max":null,
		"remote_ip_prefix":null,
		"remote_group_id":null,
		"remote_address_group_id":null,
		"tenant_id":"project-id"
	}]
}`
