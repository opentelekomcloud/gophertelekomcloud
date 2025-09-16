package shares

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get will get a single SFS Turbo file system with given UUID
func Get(client *golangsdk.ServiceClient, shareId string) (*Turbo, error) {
	// GET /v1/{project_id}/sfs-turbo/shares/{share_id}
	raw, err := client.Get(client.ServiceURL("sfs-turbo", "shares", shareId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Turbo
	err = extract.Into(raw.Body, &res)
	return &res, err
}
