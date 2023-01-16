package events

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EventItem struct {
	// Specifies the event name.
	// Start with a letter. Enter 1 to 64 characters. Only letters, digits, and underscores (_) are allowed.
	EventName string `json:"event_name" required:"true"`
	// Specifies the event source.
	// The format is service.item. Set this parameter based on the site requirements.
	// service and item each must be a string that starts with a letter and contains 3 to 32 characters, including only letters, digits, and underscores (_).
	EventSource string `json:"event_source" required:"true"`
	// Specifies when the event occurred, which is a UNIX timestamp (ms).
	// NOTE
	// Since there is a latency between the client and the server, the data timestamp to be inserted should be within
	// the period that starts from one hour before the current time plus 20s to 10 minutes after the current time minus 20s.
	// In this way, the timestamp will be inserted to the database without being affected by the latency.
	// For example, if the current time is 2020.01.30 12:00:30, the timestamp inserted must be within the range
	// [2020.01.30 11:00:50, 2020.01.30 12:10:10]. The corresponding UNIX timestamp is [1580353250, 1580357410].
	Time   int64           `json:"time" required:"true"`
	Detail EventItemDetail `json:"detail" required:"true"`
}

type EventItemDetail struct {
	// Specifies the event content. Enter up to 4096 characters.
	Content string `json:"content,omitempty"`
	// Specifies the resource ID. Enter up to 128 characters, including letters, digits, underscores (_), hyphens (-), and colon (:).
	// Example: 6a69bf28-ee62-49f3-9785-845dacd799ec
	// To query the resource ID, perform the following steps:
	// 1.	Log in to the management console.
	// 2.	Under Computing, select Elastic Cloud Server.
	// On the Resource Overview page, obtain the resource ID.
	ResourceId string `json:"resource_id,omitempty"`
	// Specifies the resource name. Enter up to 128 characters, including letters, digits, underscores (_), and hyphens (-).
	ResourceName string `json:"resource_name,omitempty"`
	// Specifies the event status.
	// Valid value can be normal, warning, or incident.
	EventState string `json:"event_state,omitempty"`
	// Specifies the event severity.
	// Its value can be Critical, Major, Minor, or Info.
	EventLevel string `json:"event_level,omitempty"`
	// Specifies the event user.
	// Enter up to 64 characters, including letters, digits, underscores (_), hyphens (-), slashes (/), and spaces.
	EventUser string `json:"event_user,omitempty"`

	GroupId   string `json:"group_id,omitempty"`
	EventType string `json:"event_type,omitempty"`
}

func CreateEvents(client *golangsdk.ServiceClient, opts []EventItem) ([]CreateEventsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /V1.0/{project_id}/events
	raw, err := client.Post(client.ServiceURL("events"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []CreateEventsResponse
	err = extract.Into(raw.Body, &res)
	return res, err
}

type CreateEventsResponse struct {
	// Specifies the event ID.
	EventId string `json:"event_id"`
	// Specifies the event name.
	// Start with a letter. Enter 1 to 64 characters. Only letters, digits, and underscores (_) are allowed.
	EventName string `json:"event_name"`
}
