package key

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	GatewayID     string `json:"-"`
	SignID        string `json:"-"`
	Name          string `json:"name" required:"true"`
	SignType      string `json:"sign_type,omitempty"`
	SignKey       string `json:"sign_key,omitempty"`
	SignSecret    string `json:"sign_secret,omitempty"`
	SignAlgorithm string `json:"sign_algorithm,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*SignKeyResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "signs", opts.SignID),
		b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res SignKeyResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
