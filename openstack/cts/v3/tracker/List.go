package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func List(client *golangsdk.ServiceClient, trackerName string) ([]Tracker, error) {
	// GET /v3/{project_id}/trackers
	raw, err := client.Get(client.ServiceURL("trackers")+"?tracker_name="+trackerName, nil, nil)
	return extraStruct(err, raw)
}
