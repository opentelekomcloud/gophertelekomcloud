package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

// PolicyCreate will create a new snapshot policy based on the values in PolicyCreateOpts.
func PolicyCreate(client *golangsdk.ServiceClient, opts CreateOptsBuilder, clusterId string) (r ErrorResult) {
	b, err := opts.ToSnapshotCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("clusters", clusterId, "index_snapshot/policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
