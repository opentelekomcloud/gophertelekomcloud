package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create will create a new snapshot based on the values in CreateOpts.
// To extract the result from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder, clusterId string) (r CreateResult) {
	b, err := opts.ToSnapshotCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("clusters", clusterId, "index_snapshot"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
