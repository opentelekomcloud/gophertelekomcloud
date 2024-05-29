package authorizer

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, authorizerID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "authorizers", authorizerID), nil)
	return
}
