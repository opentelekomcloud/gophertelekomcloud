package gateway

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// DisableIngressAccess is a method to unbind the eip associated with an existing APIG dedicated instance.
func DisableIngressAccess(client *golangsdk.ServiceClient, instanceId string) error {
	_, err := client.Delete(client.ServiceURL("apigw", "instances", instanceId, "eip"), nil)
	return err
}
