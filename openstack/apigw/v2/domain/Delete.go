package domain

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	GatewayID string
	GroupID   string
	DomainID  string
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	// DELETE /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}
	_, err = client.Delete(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "domains", opts.DomainID), nil)
	return
}
