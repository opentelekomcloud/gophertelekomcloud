package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tracker.
func Delete(client *golangsdk.ServiceClient, trackerName string) (err error) {
	// DELETE /v1.0/{project_id}/tracker?tracker_name={tracker_name}
	_, err = client.Delete(client.ServiceURL("tracker")+"?tracker_name="+trackerName, nil)
	return
}
