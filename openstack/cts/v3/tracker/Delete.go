package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tracker.
func Delete(client *golangsdk.ServiceClient, trackerName string) (err error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("trackers").WithQueryParams(&tracker{Tracker: trackerName}).Build()
	if err != nil {
		return err
	}

	// DELETE /v3/{project_id}/trackers?tracker_name=system
	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return
}
