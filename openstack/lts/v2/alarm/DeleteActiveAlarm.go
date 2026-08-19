package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteActiveAlarmOpts struct {
	// Events contains the active alarms to delete.
	Events []DeleteActiveAlarmEvent `json:"events" required:"true"`
}

type DeleteActiveAlarmEvent struct {
	// Metadata contains the alarm information.
	Metadata *DeleteActiveAlarmMetadata `json:"metadata" required:"true"`
	// StartsAt is the alarm generation time as a timestamp.
	StartsAt *int64 `json:"starts_at" required:"true"`
}

type DeleteActiveAlarmMetadata struct {
	// EventType is the alarm type.
	EventType string `json:"event_type" required:"true"`
	// EventID is the alarm ID.
	EventID string `json:"event_id" required:"true"`
	// EventSeverity is the alarm severity.
	EventSeverity string `json:"event_severity" required:"true"`
	// EventName is the alarm name.
	EventName string `json:"event_name" required:"true"`
	// ResourceType is the resource type.
	ResourceType string `json:"resource_type" required:"true"`
	// ResourceID is the log group or stream name.
	ResourceID string `json:"resource_id" required:"true"`
	// ResourceProvider is the alarm source.
	ResourceProvider string `json:"resource_provider" required:"true"`
	// LTSAlarmType is the alarm rule type.
	LTSAlarmType string `json:"lts_alarm_type" required:"true"`
}

func DeleteActiveAlarm(client *golangsdk.ServiceClient, domainID string, opts DeleteActiveAlarmOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/{domain_id}/lts/alarms/sql-alarm/clear
	_, err = client.Post(
		client.ServiceURL(domainID, "lts", "alarms", "sql-alarm", "clear"),
		b,
		nil,
		&golangsdk.RequestOpts{
			MoreHeaders: map[string]string{
				"Content-Type": "application/json;charset=UTF-8",
			},
			OkCodes: []int{200},
		},
	)
	return err
}
