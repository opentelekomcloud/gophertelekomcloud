package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	FuncUrn string `json:"-"`
	Name    string `json:"name" required:"true"`
	Content string `json:"content" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*EventFuncResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions", opts.FuncUrn, "events"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res EventFuncResp
	return &res, extract.Into(raw.Body, &res)
}

type EventFuncResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
