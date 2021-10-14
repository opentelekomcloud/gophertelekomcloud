package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	resourcePath = "flavors"
)

func listURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL(resourcePath)
}

func getURL(sc *golangsdk.ServiceClient, flavorID string) string {
	return sc.ServiceURL(resourcePath, flavorID)
}
