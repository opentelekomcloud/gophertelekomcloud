package listeners

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Listener, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("listeners", id).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "listener")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Listener Listener `json:"listener"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Listener, nil
}
