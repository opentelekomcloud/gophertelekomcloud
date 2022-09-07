package nodepools

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
)

// UpdateOpts contains all the values needed to update a new node pool
type UpdateOpts struct {
	// API type, fixed value Node
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiversion" required:"true"`
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
	// Node type. Currently, only VM nodes are supported.
	Type string `json:"type,omitempty"`
	// Node template
	NodeTemplate UpdateNodeTemplate `json:"nodeTemplate,omitempty"`
	// Initial number of expected nodes
	InitialNodeCount int `json:"initialNodeCount"`
	// Auto scaling parameters
	Autoscaling AutoscalingSpec `json:"autoscaling,omitempty"`
}

type UpdateNodeTemplate struct {
	// Tag of a Kubernetes node, key value pair format
	K8sTags map[string]string `json:"k8sTags,omitempty"`
	// taints to created nodes to configure anti-affinity
	Taints []nodes.TaintSpec `json:"taints,omitempty"`
}

// Update allows node pools to be updated.
func Update(client *golangsdk.ServiceClient, clusterid, nodepoolid string, opts UpdateOpts) (*NodePool, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", clusterid, "nodepools", nodepoolid), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res NodePool
	err = extract.Into(raw.Body, &res)
	return &res, err
}
