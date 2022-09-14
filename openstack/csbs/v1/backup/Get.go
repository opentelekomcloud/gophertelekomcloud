package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get will get a single backup with specific ID. To extract the Backup object from the response,
// call the ExtractBackup method on the GetResult.
func Get(client *golangsdk.ServiceClient, backupID string) (*Backup, error) {
	raw, err := client.Get(client.ServiceURL("checkpoint_items", backupID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Backup
	err = extract.IntoStructPtr(raw.Body, &res, "checkpoint_item")
	return &res, err
}
