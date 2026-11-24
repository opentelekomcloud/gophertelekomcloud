package resolverrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// DisassociateVPC is used to disassociate a VPC from an endpoint rule.
// Send POST /v2.1/resolverrules/{resolverrule_id}/disassociaterouter
func DisassociateVPC(client *golangsdk.ServiceClient, ruleID string, opts RouterOpts) (*Router, error) {
	return routerAction(client, ruleID, "disassociaterouter", opts)
}
