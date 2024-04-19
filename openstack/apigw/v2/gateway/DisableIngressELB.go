package gateway

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// DisableElbIngressAccess is a method to unbind the ingress eip associated with an existing APIG dedicated instance.
// Supported only when load balancer_provider is set to elb.
func DisableElbIngressAccess(client *golangsdk.ServiceClient, instanceId string) error {
	_, err := client.Delete(client.ServiceURL("apigw", "instances", instanceId, "ingress-eip"), nil)
	return err
}
