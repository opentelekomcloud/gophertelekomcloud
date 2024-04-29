package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func UnbindPolicy(client *golangsdk.ServiceClient, gatewayID, aclBindId string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "acl-bindings", aclBindId), nil)
	return
}
