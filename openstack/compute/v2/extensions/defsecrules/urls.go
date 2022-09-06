package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

const rulepath = "os-security-group-default-rules"

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rulepath, id)
}

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rulepath)
}
