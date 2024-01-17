package api

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, envID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "apis", envID), nil)
	return
}
