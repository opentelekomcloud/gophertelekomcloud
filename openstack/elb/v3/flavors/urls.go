package flavors

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "elb"
	resourcePath = "flavors"
)

func listURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL(rootPath, resourcePath)
}

func getURL(sc *golangsdk.ServiceClient, flavorID string) string {
	return sc.ServiceURL(rootPath, resourcePath, flavorID)
}
