package members_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

const memberJSON = `{
	"id":"member-id","availability_zone":"eu-de-01","name":"member-test","project_id":"project-id",
	"pool_id":"pool-id","admin_state_up":true,"subnet_cidr_id":"subnet-id","protocol_port":8080,
	"weight":10,"address":"192.0.2.10","ip_version":"v4","operating_status":"ONLINE",
	"status":[{"listener_id":"listener-id","operating_status":"ONLINE","reason":{"reason_code":"CONNECT_TIMEOUT","expected_response":"200","healthcheck_response":"500"},"created_at":"2026-10-21T08:00:00Z","updated_at":"2026-10-21T08:01:00Z"}],
	"created_at":"2026-10-21T08:00:00Z","updated_at":"2026-10-21T08:01:00Z","member_type":"instance",
	"instance_id":"instance-id","device_owner":"compute:eu-de-01","device_id":"device-id",
	"loadbalancer_id":"loadbalancer-id","loadbalancers":[{"id":"loadbalancer-id"}]
}`
