package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Backup, error) {
	raw, err := client.Get(client.ServiceURL("backups", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Backup
	err = extract.IntoStructPtr(raw.Body, &res, "backup")
	return &res, err
}
