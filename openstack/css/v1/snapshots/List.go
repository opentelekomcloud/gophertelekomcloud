package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List retrieves the Snapshots with the provided ID. To extract the Snapshot
// objects from the response, call the Extract method on the GetResult.
func List(client *golangsdk.ServiceClient, clusterId string) ([]Snapshot, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "index_snapshots"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Snapshot
	err = extract.IntoSlicePtr(raw.Body, &res, "backups")
	return res, err
}
