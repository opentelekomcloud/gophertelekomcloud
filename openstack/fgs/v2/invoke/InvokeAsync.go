package invoke

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// LaunchAsync is used to execute a function asynchronously.
func LaunchAsync(client *golangsdk.ServiceClient, funcUrn string, JSONBody interface{}) (*LaunchAsyncResp, error) {

	raw, err := client.Post(client.ServiceURL("fgs", "functions", funcUrn, "invocations-async"), JSONBody, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
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
