package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestoreToPITOpts struct {
	// Specifies the database information.
	RestoreCollections []RestoreCollections `json:"restore_collections" required:"true"`
}

type RestoreCollections struct {
	// Specifies the database name.
	Database string `json:"database" required:"true"`
	// Specifies the collection information.
	Collections []Collections `json:"collections,omitempty"`
	// Specifies the database restoration time point.
	RestoreDatabaseTime string `json:"restore_database_time,omitempty"`
}

type Collections struct {
	// Specifies the original table name before the restoration.
	OldName string `json:"old_name" required:"true"`
	// Specifies the table name after the restoration.
	NewName string `json:"new_name,omitempty"`
	// Specifies the collection restoration time point.
	// The value is a UNIX timestamp, in milliseconds. The time zone is UTC.
	RestoreCollectionTime string `json:"restore_collection_time,omitempty"`
}

func RestoreToPIT(client *golangsdk.ServiceClient, opts RestoreToPITOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	return ExtractJob(err, raw)
}
