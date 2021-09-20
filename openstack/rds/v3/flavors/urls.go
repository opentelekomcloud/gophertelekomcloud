package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

const rootPath = "flavors"

func listURL(client *golangsdk.ServiceClient, dbName string) string {
	return client.ServiceURL(rootPath, dbName)
}
