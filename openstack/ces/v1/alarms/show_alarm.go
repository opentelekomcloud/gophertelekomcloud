package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowAlarm(client *golangsdk.ServiceClient, id string) ([]MetricAlarms, error) {
	// GET /V1.0/{project_id}/alarms/{alarm_id}
	raw, err := client.Get(client.ServiceURL("alarms", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []MetricAlarms
	err = extract.IntoSlicePtr(raw.Body, &res, "metric_alarms")
	return res, err
}
