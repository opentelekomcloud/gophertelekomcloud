package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListAlarmsOpts struct {
	// The value ranges from 1 to 100, and is 100 by default.
	// This parameter is used to limit the number of query results.
	Limit int `q:"limit"`
	// Specifies the result sorting method, which is sorted by timestamp.
	// The default value is desc.
	// asc: The query results are displayed in the ascending order.
	// desc: The query results are displayed in the descending order.
	Order string `q:"order"`
	// Specifies the first queried alarm to be displayed on a page.
	Start string `q:"start"`
}

func ListAlarms(client *golangsdk.ServiceClient, opts ListAlarmsOpts) (*ListAlarmsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /V1.0/{project_id}/alarms
	raw, err := client.Get(client.ServiceURL("alarms")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var s ListAlarmsResponse
	err = extract.Into(raw.Body, &s)
	return &s, err
}
