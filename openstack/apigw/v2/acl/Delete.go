package acl

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, aclID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "acls", aclID), nil)
	return
}
