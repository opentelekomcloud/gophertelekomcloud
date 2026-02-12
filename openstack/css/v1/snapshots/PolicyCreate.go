package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

type PolicyCreateOpts struct {
	// Whether to enable the automatic snapshot creation policy.
	Enable string `json:"enable" required:"true"`
	// Snapshot name prefix. Mandatory when Enable is set to true.
	// It must be 1 to 32 characters and can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	Prefix string `json:"prefix,omitempty"`
	// Time when a snapshot is created every day. Mandatory when Enable is set to true.
	// Format is HH:mm z, e.g. 00:00 GMT+08:00.
	Period string `json:"period,omitempty"`
	// Number of days for which automatically created snapshots are reserved. Mandatory when Enable is set to true.
	// Value range: 1 to 90.
	KeepDay int `json:"keepday,omitempty"`
	// Snapshot creation frequency.
	// Options: HOUR, DAY, SUN, MON, TUE, WED, THU, FRI, SAT. Default: DAY.
	Frequency string `json:"frequency,omitempty"`
	// Whether to delete all automatically created snapshots when the automatic snapshot creation policy is disabled.
	// Default: false.
	DeleteAuto string `json:"delete_auto,omitempty"`
	// Name of the index to be backed up. Supports wildcards, e.g. index*.
	Indices string `json:"indices,omitempty"`
}

// PolicyCreate will create a new snapshot policy based on the values in PolicyCreateOpts.
func PolicyCreate(client *golangsdk.ServiceClient, opts PolicyCreateOpts, clusterId string) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "index_snapshot/policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
