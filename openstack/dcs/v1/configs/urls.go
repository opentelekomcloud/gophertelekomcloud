package configs

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath   = "instances"
	configPath = "configs"
)

func rootURL(client *golangsdk.ServiceClient, instanceID string) string {
	return client.ServiceURL(rootPath, instanceID, configPath)
}
