package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	secgrouppath = "os-security-groups"
	rulepath     = "os-security-group-rules"
)

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(secgrouppath, id)
}

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(secgrouppath)
}

func listByServerURL(client *golangsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, secgrouppath)
}

func rootRuleURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rulepath)
}

func resourceRuleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(rulepath, id)
}

func serverActionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
