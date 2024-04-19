package trigger

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	FuncUrn         string            `json:"-"`
	TriggerId       string            `json:"-"`
	TriggerTypeCode string            `json:"-"`
	TriggerStatus   string            `json:"trigger_status,omitempty"`
	EventData       []UpdateEventData `json:"event_data,omitempty"`
}

type UpdateEventData struct {
	IsSerial        *bool `json:"is_serial,omitempty"`
	MaxFetchBytes   *int  `json:"max_fetch_bytes,omitempty"`
	PollingInterval *int  `json:"polling_interval,omitempty"`
	PollingUnit     *int  `json:"polling_unit,omitempty"`
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
