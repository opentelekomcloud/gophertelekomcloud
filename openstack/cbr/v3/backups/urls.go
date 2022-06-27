package backups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const rootURL = "backups"

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rootURL, id)
}
