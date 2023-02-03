package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete resets the quotas for the given tenant to their default values.
func Delete(client *golangsdk.ServiceClient, tenantID string) (err error) {
	_, err = client.Delete(client.ServiceURL("os-quota-sets", tenantID), nil)
	return
}
