package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ClusterFlavorOpts struct {
	// Indicates whether to verify replicas.
	NeedCheckReplica *bool `json:"needCheckReplica,omitempty"`
	// ID of the new flavor.
	NewFlavorID string `json:"newFlavorId" required:"true"`
	// Type of the cluster node to modify.
	NodeType string `json:"-"`
}

// UpdateClusterFlavor is used to modify the specifications of a cluster or specifications of a specified node type.
func UpdateClusterFlavor(client *golangsdk.ServiceClient, clusterID string, opts ClusterFlavorOpts) error {
	// Construct the URL dynamically based on the optional NodeType.

	var url string

	if opts.NodeType != "" {
		url = client.ServiceURL("clusters", clusterID, opts.NodeType, "flavor")
	} else {
		url = client.ServiceURL("clusters", clusterID, "flavor")
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
