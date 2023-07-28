package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListEventDetailOpts struct {
	// Specifies the event name.
	EventName string
	// Specifies the event type. Possible types are EVENT.SYS (system event) and EVENT.CUSTOM (custom event).
	EventType string `q:"event_type" required:"true"`
	// Specifies the event name. The name can be a system event name or a custom event name.
	EventSource string `q:"event_source,omitempty"`
	// Specifies the event severity. Possible severities are Critical, Major, Minor, and Info.
	EventLevel string `q:"event_level,omitempty"`
	// Specifies the name of the user who reports the event monitoring data. It can also be a project ID.
	EventUser string `q:"event_user,omitempty"`
	// Specifies the event status. Possible statuses are normal, warning, or incident.
	EventState string `q:"event_state,omitempty"`
	// Specifies the start time of the query. The time is a UNIX timestamp and the unit is ms. Example: 1605952700911
	From int64 `q:"from,omitempty"`
	// No
	// Specifies the end time of the query. The time is a UNIX timestamp and the unit is ms. The from value must be smaller than the to value.
	To int64 `q:"to,omitempty"`
	// Specifies the start value of pagination. The value is an integer. The default value is 0.
	Start int `q:"start,omitempty"`
	// Specifies the maximum number of records that can be queried at a time. Supported range: 1 to 100 (default)
	Limit int `q:"limit,omitempty"`
}

func ListEventDetail(client *golangsdk.ServiceClient, opts ListEventDetailOpts) (*ListEventDetailResponse, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /V1.0/{project_id}/event/{event_name}
	raw, err := client.Get(client.ServiceURL("event", opts.EventName)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListEventDetailResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListEventDetailResponse struct {
	// Specifies the event name. The name can be a system event name or a custom event name.
	EventName string `json:"event_name,omitempty"`
	// Specifies the event type. Possible types are EVENT.SYS (system event) and EVENT.CUSTOM (custom event).
	EventType string `json:"event_type,omitempty"`
	// Specifies the name of the user who reports the event. It can also be a project ID.
	EventUsers []string `json:"event_users,omitempty"`
	// Specifies the event source. If the event is a system event, the source is the namespace of each service.
	// To view the namespace of each service, see A.1 Services Interconnected with Cloud Eye. If the event is a custom event, the event source is defined by the user.
	EventSources []string `json:"event_sources,omitempty"`
	// Specifies details about one or more events.
	EventInfo []EventInfoDetail `json:"event_info,omitempty"`
	// Specifies the number of metadata records in the query result.
	MetaData TotalMetaData `json:"meta_data,omitempty"`
}

type EventInfoDetail struct {
	// Specifies the event name. Start with a letter. Enter 1 to 64 characters. Only letters, digits, and underscores (_) are allowed.
	EventName string `json:"event_name"`
	// Specifies the event source in the format of service.item. service and item each must start with
	// a letter and contain 3 to 32 characters, including only letters, digits, and underscores (_).
	EventSource string `json:"event_source"`
	// Specifies when the event occurred, which is a UNIX timestamp (ms).
	// Since there is a latency between the client and the server, the data timestamp to be inserted should be within
	// the period that starts from one hour before the current time plus 20s to 10 minutes after the current time minus 20s.
	// In this way, the timestamp will be inserted to the database without being affected by the latency.
	Time int64 `json:"time"`
	// Specifies the event details.
	Detail EventItemDetail `json:"detail"`
	// Specifies the event ID.
	EventId string `json:"event_id,omitempty"`
}

type TotalMetaData struct {
	// Specifies the total number of events.
	Total int `json:"total,omitempty"`
}
