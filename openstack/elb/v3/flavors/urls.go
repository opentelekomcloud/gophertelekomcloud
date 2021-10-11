package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	v3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3"
)

const (
	resourcePath = "flavors"
)

func listURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL(v3.RootPath, resourcePath)
}

func getURL(sc *golangsdk.ServiceClient, flavorID string) string {
	return sc.ServiceURL(v3.RootPath, resourcePath, flavorID)
}
