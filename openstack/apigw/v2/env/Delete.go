package env

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, envID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "envs", envID), nil)
	return
}
