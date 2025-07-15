package channel

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, channelID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "vpc-channels", channelID), nil)
	return
}
