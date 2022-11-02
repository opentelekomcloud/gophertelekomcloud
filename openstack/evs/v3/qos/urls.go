package qos

import "github.com/opentelekomcloud/gophertelekomcloud"

func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id)
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("qos-specs")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("qos-specs")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("qos-specs", id)
}

func updateURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id)
}

func deleteKeysURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "delete_keys")
}

func associateURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "associate")
}

func disassociateURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "disassociate")
}

func disassociateAllURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "disassociate_all")
}

func listAssociationsURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "associations")
}
