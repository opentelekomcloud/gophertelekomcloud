package trigger

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	FuncUrn         string                 `json:"-"`
	TriggerTypeCode string                 `json:"trigger_type_code" required:"true"`
	TriggerStatus   string                 `json:"trigger_status,omitempty"`
	EventTypeCode   string                 `json:"event_type_code,omitempty"`
	EventData       map[string]interface{} `json:"event_data" required:"true"`
}

type TriggerFuncInfo struct {
	FunctionUrn    string `json:"function_urn,omitempty"`
	InvocationType string `json:"invocation_type,omitempty"`
	Timeout        int    `json:"timeout" required:"true"`
	Version        string `json:"version,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*TriggerFuncResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "triggers", opts.FuncUrn), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res TriggerFuncResp
	return &res, extract.Into(raw.Body, &res)
}

type TriggerFuncResp struct {
	TriggerId       string                 `json:"trigger_id"`
	TriggerTypeCode string                 `json:"trigger_type_code"`
	TriggerStatus   string                 `json:"trigger_status"`
	EventData       map[string]interface{} `json:"event_data"`
	LastUpdatedTime string                 `json:"last_updated_time"`
	CreatedTime     string                 `json:"created_time"`
}
