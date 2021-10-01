package subscriptions

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath   = "subscriptions"
	topicsPath = "topics"
)

func createURL(c *golangsdk.ServiceClient, topicUrn string) string {
	return c.ServiceURL(topicsPath, topicUrn, rootPath)
}

func deleteURL(c *golangsdk.ServiceClient, subscriptionUrn string) string {
	return c.ServiceURL(rootPath, subscriptionUrn)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func listFromTopicURL(c *golangsdk.ServiceClient, topicUrn string) string {
	return c.ServiceURL(topicsPath, topicUrn, rootPath)
}
