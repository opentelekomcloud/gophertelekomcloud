package alarm

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]KeywordRule, error) {
	// GET /v2/{project_id}/lts/alarms/keywords-alarm-rule
	raw, err := client.Get(client.ServiceURL("lts", "alarms", "keywords-alarm-rule"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res []KeywordRule
	err = extract.IntoSlicePtr(raw.Body, &res, "keywords_alarm_rules")
	return res, err
}
