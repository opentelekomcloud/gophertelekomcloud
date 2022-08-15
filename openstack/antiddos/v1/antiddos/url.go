package antiddos

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func CreateURL(c *golangsdk.ServiceClient, floatingIpId string) string {
	return c.ServiceURL("antiddos", floatingIpId)
}

func DeleteURL(c *golangsdk.ServiceClient, floatingIpId string) string {
	return c.ServiceURL("antiddos", floatingIpId)
}

func GetTaskURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("query_task_status")
}

func UpdateURL(c *golangsdk.ServiceClient, floatingIpId string) string {
	return c.ServiceURL("antiddos", floatingIpId)
}
