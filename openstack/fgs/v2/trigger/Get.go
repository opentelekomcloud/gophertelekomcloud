package trigger

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, funcURN, triggerType, triggerId string) (*TriggerFuncResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "triggers", funcURN, triggerType, triggerId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res *TriggerFuncResp
	err = extract.Into(raw.Body, &res)
	return res, err
}
