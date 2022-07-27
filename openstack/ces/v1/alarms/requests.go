package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type ListAlarmsRequest struct {
	// The value ranges from 1 to 100, and is 100 by default.
	// This parameter is used to limit the number of query results.
	Limit int32 `json:"limit,omitempty"`
	// Specifies the result sorting method, which is sorted by timestamp.
	// The default value is desc.
	// asc: The query results are displayed in the ascending order.
	// desc: The query results are displayed in the descending order.
	Order string `json:"order,omitempty"`
	// Specifies the first queried alarm to be displayed on a page.
	Start string `json:"start,omitempty"`
}

func ListAlarms(client *golangsdk.ServiceClient, req ListAlarmsRequest) (r ListAlarmsResult, err error) {
	url := alarmsURL(client)
	query, err := golangsdk.BuildQueryString(req)
	if err != nil {
		return
	}

	url += query.String()
	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
