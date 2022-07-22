package tags

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const baseURL = "vault"

// POST /v3/{project_id}/vault/resource_instances/action
func showVaultResourceInstancesURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(baseURL, "resource_instances", "action")
}
