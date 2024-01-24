package app_auth

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, appAuthID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "app-auths", appAuthID), nil)
	return
}
