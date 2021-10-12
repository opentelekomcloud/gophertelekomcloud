package alarmnotifications

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const (
	configPath   = "config"
	resourcePath = "alert"
)

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(configPath, resourcePath)
}

func resourceURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(configPath, resourcePath, id)
}
