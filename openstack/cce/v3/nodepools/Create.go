package nodepools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
)

// CreateOpts allows extensions to add additional parameters to the
// Create request.
type CreateOpts struct {
	// API type, fixed value Node
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiversion" required:"true"`
	// Metadata required to create a Node Pool
	Metadata CreateMetaData `json:"metadata" required:"true"`
	// specifications to create a Node Pool
	Spec CreateSpec `json:"spec" required:"true"`
}

// CreateMetaData required to create a Node Pool
type CreateMetaData struct {
	// Name of the node pool.
	Name string `json:"name" required:"true"`
}

// CreateSpec describes Node pools specification
type CreateSpec struct {
	// Node pool type. Currently, only `vm`(ECSs) are supported.
	Type string `json:"type" required:"true"`
	// Node template
	NodeTemplate nodes.Spec `json:"nodeTemplate" required:"true"`
	// Initial number of expected nodes
	InitialNodeCount int `json:"initialNodeCount"`
	// Auto scaling parameters
	Autoscaling AutoscalingSpec `json:"autoscaling,omitempty"`
	// Node management parameters
	NodeManagement NodeManagementSpec `json:"nodeManagement,omitempty"`
	// Custom security group settings for a node pool
	CustomSecurityGroupIds []string `json:"customSecurityGroups,omitempty"`
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node pool.
func Create(client *golangsdk.ServiceClient, clusterId string, opts CreateOpts) (*NodePool, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools
	raw, err := client.Post(client.ServiceURL("clusters", clusterId, "nodepools"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res NodePool
	return &res, extract.Into(raw.Body, &res)
}
