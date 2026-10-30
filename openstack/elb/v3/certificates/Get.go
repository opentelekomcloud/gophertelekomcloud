package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Certificate, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("certificates", id).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
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
