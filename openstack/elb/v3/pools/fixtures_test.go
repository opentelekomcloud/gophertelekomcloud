package pools_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

const poolJSON = `{
	"id":"pool-id","name":"pool-test","description":"pool","protocol":"HTTP",
	"lb_algorithm":"LEAST_CONNECTIONS","admin_state_up":true,"project_id":"project-id",
	"listeners":[{"id":"listener-id"}],"loadbalancers":[{"id":"loadbalancer-id"}],
	"members":[{"id":"member-id"}],"healthmonitor_id":"monitor-id",
	"session_persistence":{"type":"APP_COOKIE","cookie_name":"session","persistence_timeout":10},
	"ip_version":"v4","slow_start":{"enable":true,"duration":30},
	"member_deletion_protection_enable":true,"vpc_id":"vpc-id","type":"instance",
	"created_at":"2026-09-21T08:00:00Z","updated_at":"2026-09-21T08:01:00Z",
	"protection_status":"PROTECTED","protection_reason":"managed",
	"az_affinity":{"enable":true,"az_minimum_healthy_member_percentage":50,
	"az_minimum_healthy_member_count":2,"az_unhealthy_fallback_strategy":"LOCALLY_REPLICA"}
}`
