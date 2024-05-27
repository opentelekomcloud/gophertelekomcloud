package async_config

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	FuncUrn           string             `json:"-"`
	MaxEventAge       *int               `json:"max_async_event_age_in_seconds,omitempty"`
	MaxRetry          *int               `json:"max_async_retry_attempts,omitempty"`
	EnableStatusLog   *bool              `json:"enable_async_status_log,omitempty"`
	DestinationConfig *DestinationConfig `json:"destination_config,omitempty"`
}

type DestinationConfig struct {
	OnSuccess *Destination `json:"on_success,omitempty"`
	OnFailure *Destination `json:"on_failure,omitempty"`
}

type Destination struct {
	Destination string `json:"destination,omitempty"`
	Param       string `json:"param,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*AsyncInvokeResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "async-invoke-config"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AsyncInvokeResp
	return &res, extract.Into(raw.Body, &res)
}

type AsyncInvokeResp struct {
	FuncUrn           string             `json:"func_urn"`
	MaxEventAge       int                `json:"max_async_event_age_in_seconds"`
	MaxRetry          int                `json:"max_async_retry_attempts"`
	DestinationConfig *DestinationConfig `json:"destination_config,omitempty"`
	CreatedTime       string             `json:"created_time"`
	LastModified      string             `json:"last_modified"`
	EnableStatusLog   bool               `json:"enable_async_status_log"`
}
