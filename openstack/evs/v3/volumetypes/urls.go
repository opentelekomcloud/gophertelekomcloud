package volumetypes

import "github.com/opentelekomcloud/gophertelekomcloud"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func extraSpecsListURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "extra_specs")
}

func extraSpecsGetURL(client *golangsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("types", id, "extra_specs", key)
}

func extraSpecsCreateURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "extra_specs")
}

func extraSpecUpdateURL(client *golangsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("types", id, "extra_specs", key)
}

func extraSpecDeleteURL(client *golangsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("types", id, "extra_specs", key)
}

func accessURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "os-volume-type-access")
}

func accessActionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "action")
}
