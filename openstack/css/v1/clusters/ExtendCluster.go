package clusters

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ClusterExtendOptsBuilder interface {
}

// ClusterExtendCommonOpts is used to extend cluster with only `common` nodes.
// Clusters with master, client, or cold data nodes cannot use this.
type ClusterExtendCommonOpts struct {
	// ModifySize - number of instances to be added.
	ModifySize int `json:"modifySize"`
}

// ClusterExtendSpecialOpts is used to extend cluster with special nodes.
// If a cluster has master, client, or cold data nodes, this should be used
type ClusterExtendSpecialOpts struct {
	// Type of the instance to be scaled out.
	// Select at least one from `ess`, `ess-cold`, `ess-master`, and `ess-client`.
	// You can only add instances, rather than increase storage capacity, on nodes of the `ess-master` and `ess-client` types.
	Type string `json:"type"`
	// NodeSize - number of instances to be scaled out.
	// The total number of existing instances and newly added instances in a cluster cannot exceed 32.
	NodeSize int `json:"nodesize"`
	// DiskSize - Storage capacity of the instance to be expanded.
	// The total storage capacity of existing instances and newly added instances in a cluster cannot exceed the maximum instance storage capacity allowed when a cluster is being created.
	// In addition, you can expand the instance storage capacity for a cluster for up to six times.
	// Unit: GB
	DiskSize int `json:"disksize"`
}

// ExtendCluster - extends cluster capacity.
// ClusterExtendCommonOpts should be used as options to extend cluster with only `common` nodes.
// ClusterExtendSpecialOpts should be used if extended cluster has master, client, or cold data nodes.
func ExtendCluster(client *golangsdk.ServiceClient, clusterID string, opts ClusterExtendOptsBuilder) (*ExtendedCluster, error) {
	url := ""
	switch opts.(type) {
	case ClusterExtendCommonOpts:
		url = client.ServiceURL("clusters", clusterID, "extend")
	case []ClusterExtendSpecialOpts:
		url = client.ServiceURL("clusters", clusterID, "role_extend")
	default:
		return nil, fmt.Errorf("invalid options type provided: %T", opts)
	}

	b, err := build.RequestBody(opts, "grow")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res ExtendedCluster
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ExtendedCluster struct {
	ID        string             `json:"id"`
	Instances []ExtendedInstance `json:"instances"`
}

type ExtendedInstance struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	ShardID string `json:"shard_id"`
}
