package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tracker.
func Delete(client *golangsdk.ServiceClient, trackerName string) (err error) {
	// DELETE /v3/{project_id}/trackers?tracker_name=system
	_, err = client.Delete(client.ServiceURL("trackers")+"?tracker_name="+trackerName, nil)
	return
}
