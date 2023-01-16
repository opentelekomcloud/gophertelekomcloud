package keyevent

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	NotificationId []string `q:"notification_id"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return
	}

	// DELETE /v3/{project_id}/notifications
	url := client.ServiceURL("notifications") + q.String()
	_, err = client.Delete(url, nil)
	return
}
