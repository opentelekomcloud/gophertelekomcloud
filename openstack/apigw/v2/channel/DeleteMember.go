package channel

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DeleteMember(client *golangsdk.ServiceClient, gatewayID, channelID, memberID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "vpc-channels", channelID, "members", memberID), nil)
	return
}
