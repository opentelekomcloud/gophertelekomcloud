package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListEventsOpts struct {
	// Specifies the event type. Possible types are EVENT.SYS (system event) and EVENT.CUSTOM (custom event).
	EventType string `q:"event_type,omitempty"`
	// Specifies the event name. The name can be a system event name or a custom event name.
	EventName string `q:"event_name,omitempty"`
	// Specifies the start time of the query. The time is a UNIX timestamp and the unit is ms. Example: 1605952700911
	From int64 `q:"from,omitempty"`
	// Specifies the end time of the query. The time is a UNIX timestamp and the unit is ms.
	// from must be smaller than to. For example, set to 1606557500911.
	To int64 `q:"to,omitempty"`
	// Specifies the start value of pagination. The value is an integer. The default value is 0.
	Start int `q:"start,omitempty"`
	// Specifies the maximum number of events that can be queried at a time. Supported range: 1 to 100 (default)
	Limit int `q:"limit,omitempty"`
}

func ListEvents(client *golangsdk.ServiceClient, opts ListEventsOpts) (*ListEventsResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /V1.0/{project_id}/events
	raw, err := client.Get(client.ServiceURL("events")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListEventsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListEventsResponse struct {
	// Specifies one or more pieces of event data.
	Events []EventInfo `json:"events,omitempty"`
	// Specifies the number of metadata records in the query result.
	MetaData TotalMetaData `json:"meta_data,omitempty"`
}

type EventInfo struct {
	// Specifies the event name.
	EventName string `json:"event_name,omitempty"`
	// Specifies the event type.
	EventType string `json:"event_type,omitempty"`
	// Specifies the number of occurrences of this event within the specified query time range.
	EventCount int `json:"event_count,omitempty"`
	// Specifies when the event last occurred.
	LatestOccurTime int64 `json:"latest_occur_time,omitempty"`
	// Specifies the event source. If the event is a system event, the value is the namespace of each service.
	// To view the namespace of each service, see A.1 Services Interconnected with Cloud Eye.
	// If the event is a custom event, the event source is defined by the user.
	LatestEventSource string `json:"latest_event_source,omitempty"`
}
