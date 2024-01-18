package key

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID     string `json:"-"`
	Name          string `json:"name" required:"true"`
	SignType      string `json:"sign_type,omitempty"`
	SignKey       string `json:"sign_key,omitempty"`
	SignSecret    string `json:"sign_secret,omitempty"`
	SignAlgorithm string `json:"sign_algorithm,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SignKeyResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "signs"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res SignKeyResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SignKeyResp struct {
	Name          string `json:"name"`
	SignType      string `json:"sign_type"`
	SignKey       string `json:"sign_key"`
	SignSecret    string `json:"sign_secret"`
	SignAlgorithm string `json:"sign_algorithm"`
	UpdateTime    string `json:"update_time"`
	CreateTime    string `json:"create_time"`
	ID            string `json:"id"`
}
