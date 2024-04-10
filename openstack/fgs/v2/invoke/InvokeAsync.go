package invoke

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func LaunchAsync(client *golangsdk.ServiceClient, funcUrn string, opts map[string]string) (*LaunchAsyncResp, error) {
	b, err := build.RequestBody(opts, "body")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions", funcUrn, "invocations-async"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res LaunchAsyncResp
	return &res, extract.Into(raw.Body, &res)
}

type LaunchAsyncResp struct {
	RequestID string `json:"request_id"`
}
