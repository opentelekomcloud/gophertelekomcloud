package nodepools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
)

// UpdateOpts contains all the values needed to update a new node pool
type UpdateOpts struct {
	// Metadata required to update a Node Pool
	Metadata UpdateMetaData `json:"metadata" required:"true"`
	// specifications to update a Node Pool
	Spec UpdateSpec `json:"spec,omitempty" required:"true"`
}

// UpdateMetaData required to update a Node Pool
type UpdateMetaData struct {
	// Name of the node pool.
	Name string `json:"name" required:"true"`
}

// UpdateSpec describes Node pools update specification
type UpdateSpec struct {
	// Node template
	NodeTemplate UpdateNodeTemplate `json:"nodeTemplate" required:"true"`
	// Initial number of expected nodes
	InitialNodeCount int `json:"initialNodeCount" required:"true"`
	// Auto scaling parameters
	Autoscaling UpdateAutoscalingSpec `json:"autoscaling,omitempty"`
}

type UpdateNodeTemplate struct {
	// Tag of a Kubernetes node, key value pair format
	K8sTags map[string]string `json:"k8sTags,omitempty"`
	// taints to created nodes to configure anti-affinity
	Taints []nodes.TaintSpec `json:"taints,omitempty"`
}

type UpdateAutoscalingSpec struct {
	// Whether to enable auto scaling
	Enable bool `json:"enable,omitempty"`
	// Minimum number of nodes allowed if auto scaling is enabled
	MinNodeCount int `json:"minNodeCount,omitempty"`
	// This value must be greater than or equal to the value of minNodeCount
	MaxNodeCount int `json:"maxNodeCount,omitempty"`
	// Interval between two scaling operations, in minutes
	ScaleDownCooldownTime int `json:"scaleDownCooldownTime,omitempty"`
	// Weight of a node pool
	Priority int `json:"priority,omitempty"`
}

// Update allows node pools to be updated.
func Update(client *golangsdk.ServiceClient, clusterId, nodepoolId string, opts UpdateOpts) (*NodePool, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools/{nodepool_id}
	raw, err := client.Put(client.ServiceURL("clusters", clusterId, "nodepools", nodepoolId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res NodePool
	return &res, extract.Into(raw.Body, &res)
}
