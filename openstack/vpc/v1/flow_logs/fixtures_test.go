package flow_logs_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

const flowLogJSON = `{
	"id": "flow-log-id",
	"name": "flow-log",
	"tenant_id": "project-id",
	"description": "description",
	"resource_type": "vpc",
	"resource_id": "vpc-id",
	"traffic_type": "all",
	"log_group_id": "group-id",
	"log_topic_id": "topic-id",
	"index_enabled": true,
	"admin_state": true,
	"status": "ACTIVE",
	"created_at": "2019-01-14T11:03:02",
	"updated_at": "2019-01-14T12:03:02"
}`
