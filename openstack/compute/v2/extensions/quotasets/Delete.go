package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete resets the quotas for the given tenant to their default values.
func Delete(client *golangsdk.ServiceClient, tenantID string) (r DeleteResult) {
	raw, err := client.Delete(deleteURL(client, tenantID), nil)
	return
}
