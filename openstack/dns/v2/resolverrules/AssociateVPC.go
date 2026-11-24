package resolverrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// RouterOpts is the RouterForRule object used to associate or disassociate a VPC.
type RouterOpts struct {
	// ID of the Router (VPC) associated with the zone.
	RouterID string `json:"router_id" required:"true"`
}

// AssociateVPC is used to associate a VPC with an endpoint rule.
// Send POST /v2.1/resolverrules/{resolverrule_id}/associaterouter
func AssociateVPC(client *golangsdk.ServiceClient, ruleID string, opts RouterOpts) (*Router, error) {
	return routerAction(client, ruleID, "associaterouter", opts)
}

func routerAction(client *golangsdk.ServiceClient, ruleID, action string, opts RouterOpts) (*Router, error) {
	if ruleID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "ruleID"}
	}

	b, err := build.RequestBody(opts, "router")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(v21URL(client, "resolverrules", ruleID, action), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Router
	err = extract.Into(raw.Body, &res)
	return &res, err
}
