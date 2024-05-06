package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, aclID string) (*AclResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "acls", aclID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res AclResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
