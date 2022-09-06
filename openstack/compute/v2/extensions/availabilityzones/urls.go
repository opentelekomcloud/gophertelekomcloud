package availabilityzones

import "github.com/opentelekomcloud/gophertelekomcloud"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-availability-zone")
}

func listDetailURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-availability-zone", "detail")
}
