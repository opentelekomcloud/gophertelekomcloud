package tags

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const baseURL = "vault"
const tags = "tags"

// POST /v3/{project_id}/vault/{vault_id}/tags/action
func batchCreateAndDeleteVaultTagsURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(baseURL, id, tags, "action")
}

// GET/POST /v3/{project_id}/vault/{vault_id}/tags
func vaultTagsURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(baseURL, id, tags)
}

// DELETE /v3/{project_id}/vault/{vault_id}/tags/{key}
func deleteVaultTagURL(client *golangsdk.ServiceClient, id string, key string) string {
	return client.ServiceURL(baseURL, id, tags, key)
}

// GET /v3/{project_id}/vault/tags
func showVaultProjectTagURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(baseURL, tags)
}

// POST /v3/{project_id}/vault/resource_instances/action
func showVaultResourceInstancesURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(baseURL, "resource_instances", "action")
}
