package backup

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will get a single backup with specific ID. To extract the Backup object from the response,
// call the ExtractBackup method on the GetResult.
func Get(client *golangsdk.ServiceClient, backupID string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("checkpoint_items", backupID), nil, nil)

	return
}
