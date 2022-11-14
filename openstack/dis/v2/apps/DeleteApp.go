package apps

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteApp(client *golangsdk.ServiceClient, appName string) (err error) {
	// DELETE /v2/{project_id}/apps/{app_name}
	_, err = client.Delete(client.ServiceURL("apps", appName), nil)
	return
}
