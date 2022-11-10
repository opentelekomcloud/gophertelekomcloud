package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// PolicyGet retrieves the snapshot policy with the provided cluster ID.
// To extract the snapshot policy object from the response, call the Extract method on the GetResult.
func PolicyGet(client *golangsdk.ServiceClient, clusterId string) (*Policy, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "index_snapshot/policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Policy
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// Policy contains all the information associated with a snapshot policy.
type Policy struct {
	KeepDay       int    `json:"keepday"`
	Period        string `json:"period"`
	Prefix        string `json:"prefix"`
	Bucket        string `json:"bucket"`
	BasePath      string `json:"basePath"`
	Agency        string `json:"agency"`
	Enable        string `json:"enable"`
	SnapshotCmkID string `json:"snapshotCmkId"`
}
