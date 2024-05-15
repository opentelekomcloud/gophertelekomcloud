package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	FuncUrn string `json:"-"`
	EventId string `json:"-"`
	Content string `json:"content" required:"true"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Event, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "events", opts.EventId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Event
	return &res, extract.Into(raw.Body, &res)
}
