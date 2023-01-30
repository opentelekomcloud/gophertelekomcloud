package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Get(client *golangsdk.ServiceClient, trackerName string) (*Tracker, error) {
	// GET /v1.0/{project_id}/tracker?tracker_name={tracker_name}
	raw, err := client.Get(client.ServiceURL("tracker")+"?tracker_name="+trackerName, nil, nil)
	return extra(err, raw)
}
