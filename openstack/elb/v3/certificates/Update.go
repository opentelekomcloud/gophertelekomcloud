package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Domain      string  `json:"domain,omitempty"`
	PrivateKey  string  `json:"private_key,omitempty"`
	Certificate string  `json:"certificate,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Certificate, error) {
	b, err := build.RequestBody(opts, "certificate")
	if err != nil {
		return nil, err
	}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("certificates", id).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Certificate Certificate `json:"certificate"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Certificate, nil
}
