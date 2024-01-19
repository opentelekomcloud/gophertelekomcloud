package app_code

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, appID, appCodeID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "apps", appID,
		"app-codes", appCodeID), nil)
	return
}
