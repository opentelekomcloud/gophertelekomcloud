package bootfromvolume

import "github.com/opentelekomcloud/gophertelekomcloud"

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-volumes_boot")
}
