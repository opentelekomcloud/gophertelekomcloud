package keyevent

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	NotificationId []string `q:"notification_id"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return
	}

	// DELETE /v3/{project_id}/notifications
	url := client.ServiceURL("notifications") + q.String()
	_, err = client.Delete(url, nil)
	return
}
