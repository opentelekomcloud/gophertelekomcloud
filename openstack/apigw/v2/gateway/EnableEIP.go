package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EipOpts struct {
	ID                    string `json:"-"`
	BandwidthSize         string `json:"bandwidth_size,omitempty"`
	BandwidthChargingMode string `json:"bandwidth_charging_mode,omitempty"`
}

func EnableEIP(client *golangsdk.ServiceClient, opts EipOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("apigw", "instances", opts.ID, "nat-eip"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return err
	}

	return nil
}
