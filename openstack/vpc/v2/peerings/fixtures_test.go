package peerings_test

const peeringJSON = `{
	"id":"peering-id",
	"name":"peering-test",
	"status":"ACTIVE",
	"request_vpc_info":{
		"vpc_id":"request-vpc-id",
		"tenant_id":"request-tenant-id"
	},
	"accept_vpc_info":{
		"vpc_id":"accept-vpc-id",
		"tenant_id":"accept-tenant-id"
	},
	"description":"peering description",
	"created_at":"2026-09-11T10:00:00",
	"updated_at":"2026-09-11T10:01:00"
}`
