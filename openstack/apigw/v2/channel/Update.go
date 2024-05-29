package channel

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Update(client *golangsdk.ServiceClient, channelID string, opts CreateOpts) (*ChannelResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "vpc-channels", channelID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ChannelResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
