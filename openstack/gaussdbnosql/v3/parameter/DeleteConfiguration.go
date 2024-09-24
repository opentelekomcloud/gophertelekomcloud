package parameter

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteConfiguration(client *golangsdk.ServiceClient, configId string) (err error) {
	// DELETE https://{Endpoint}/v3/{project_id}/configurations/{config_id}
	_, err = client.Delete(client.ServiceURL("configurations", configId), nil)
	return
}
