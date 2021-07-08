package topicattributes

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	topics     = "topics"
	attributes = "attributes"
)

func listURL(client *golangsdk.ServiceClient, topicURN string) string {
	return client.ServiceURL(topics, topicURN, attributes)
}

func attributeURL(client *golangsdk.ServiceClient, topicURN, attribute string) string {
	return client.ServiceURL(topics, topicURN, attributes, attribute)
}
