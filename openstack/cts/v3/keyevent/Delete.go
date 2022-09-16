package keyevent

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, notificationId []string) (err error) {
	// DELETE /v3/{project_id}/notifications
	url := client.ServiceURL("notifications") + "?notification_id=\"" + strings.Join(notificationId, ",") + "\""
	_, err = client.Delete(url, nil)
	return
}
