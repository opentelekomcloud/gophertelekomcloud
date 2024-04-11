package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tracker.
func Delete(client *golangsdk.ServiceClient, trackerName string) (err error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("tracker").WithQueryParams(&tracker{Tracker: trackerName}).Build()
	if err != nil {
		return err
	}

	// DELETE /v1.0/{project_id}/tracker?tracker_name={tracker_name}
	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return
}
