package aggregates

import "github.com/opentelekomcloud/gophertelekomcloud"

func aggregatesListURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-aggregates")
}

func aggregatesCreateURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-aggregates")
}

func aggregatesDeleteURL(client *golangsdk.ServiceClient, aggregateID string) string {
	return client.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesGetURL(client *golangsdk.ServiceClient, aggregateID string) string {
	return client.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesUpdateURL(client *golangsdk.ServiceClient, aggregateID string) string {
	return client.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesAddHostURL(client *golangsdk.ServiceClient, aggregateID string) string {
	return client.ServiceURL("os-aggregates", aggregateID, "action")
}

func aggregatesRemoveHostURL(client *golangsdk.ServiceClient, aggregateID string) string {
	return client.ServiceURL("os-aggregates", aggregateID, "action")
}

func aggregatesSetMetadataURL(client *golangsdk.ServiceClient, aggregateID string) string {
	return client.ServiceURL("os-aggregates", aggregateID, "action")
}
