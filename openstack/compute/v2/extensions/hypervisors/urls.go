package hypervisors

import "github.com/opentelekomcloud/gophertelekomcloud"

func hypervisorsListDetailURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-hypervisors", "detail")
}

func hypervisorsStatisticsURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("os-hypervisors", "statistics")
}

func hypervisorsGetURL(client *golangsdk.ServiceClient, hypervisorID string) string {
	return client.ServiceURL("os-hypervisors", hypervisorID)
}

func hypervisorsUptimeURL(client *golangsdk.ServiceClient, hypervisorID string) string {
	return client.ServiceURL("os-hypervisors", hypervisorID, "uptime")
}
