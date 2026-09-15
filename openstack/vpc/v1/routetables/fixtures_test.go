package routetables_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const routeTableJSON = `{
	"id":"route-table-id",
	"name":"route-table-test",
	"default":false,
	"routes":[{"type":"ecs","destination":"10.10.10.0/24","nexthop":"server-id","description":"route"}],
	"subnets":[{"id":"network-id"}],
	"tenant_id":"project-id",
	"vpc_id":"vpc-id",
	"description":"route table",
	"created_at":"2026-09-15T08:00:00",
	"updated_at":"2026-09-15T08:01:00"
}`
