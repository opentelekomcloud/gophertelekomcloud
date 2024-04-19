package trigger

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetTriggers(client *golangsdk.ServiceClient, funcURN string) ([]TriggerFuncResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "triggers", funcURN), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []TriggerFuncResp
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return res, err
}
