package trigger

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	FuncUrn         string                 `json:"-"`
	TriggerId       string                 `json:"-"`
	TriggerTypeCode string                 `json:"-"`
	TriggerStatus   string                 `json:"trigger_status,omitempty"`
	EventData       map[string]interface{} `json:"event_data,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*TriggerFuncResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "triggers", opts.FuncUrn, opts.TriggerTypeCode, opts.TriggerId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res TriggerFuncResp
	return &res, extract.Into(raw.Body, &res)
}
