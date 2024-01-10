package env_vars

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, envVarID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "env-variables", envVarID), nil)
	return
}
