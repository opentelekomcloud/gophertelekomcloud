package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains options for creating a snapshot.
// This object is passed to the snapshots.Create function.
type CreateOpts struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
	Indices     string `json:"indices,omitempty"`
}

// Create will create a new snapshot based on the values in CreateOpts.
// To extract the result from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts, clusterId string) (*Snapshot, error) {
	b, err := build.RequestBodyMap(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters", clusterId, "index_snapshot"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res Snapshot
	err = extract.IntoStructPtr(raw.Body, &res, "backup")
	return &res, err
}
