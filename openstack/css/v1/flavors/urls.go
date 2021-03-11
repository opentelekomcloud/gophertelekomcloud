package flavors

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const rootURL = "flavors"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rootURL)
}
