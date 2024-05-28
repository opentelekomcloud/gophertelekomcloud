package channel

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, channelID string) (*ChannelResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "vpc-channels", channelID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res ChannelResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
