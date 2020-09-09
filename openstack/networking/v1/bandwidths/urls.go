package bandwidths

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "bandwidths"

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, resourcePath, id)
}
