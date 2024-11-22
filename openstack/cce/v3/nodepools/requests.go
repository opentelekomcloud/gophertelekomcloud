package nodepools

import (
	"reflect"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json"},
}

// ListOpts allows the filtering of list data using given parameters.
type ListOpts struct {
	Name  string `json:"name"`
	Uid   string `json:"uid"`
	Phase string `json:"phase"`
}

// List returns collection of node pools.
func List(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) ([]NodePool, error) {
	var r ListResult
	_, r.Err = client.Get(rootURL(client, clusterID), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	allNodePools, err := r.ExtractNodePool()

	if err != nil {
		return nil, err
	}

	return FilterNodePools(allNodePools, opts), nil
}

func FilterNodePools(nodepools []NodePool, opts ListOpts) []NodePool {

	var refinedNodePools []NodePool
	var matched bool

	m := map[string]FilterStruct{}

	if opts.Name != "" {
		m["Name"] = FilterStruct{Value: opts.Name, Driller: []string{"Metadata"}}
	}
	if opts.Uid != "" {
		m["Id"] = FilterStruct{Value: opts.Uid, Driller: []string{"Metadata"}}
	}

	if opts.Phase != "" {
		m["Phase"] = FilterStruct{Value: opts.Phase, Driller: []string{"Status"}}
	}

	if len(m) > 0 && len(nodepools) > 0 {
		for _, nodepool := range nodepools {
			matched = true

			for key, value := range m {
				if sVal := GetStructNestedField(&nodepool, key, value.Driller); !(sVal == value.Value) {
					matched = false
				}
			}
			if matched {
				refinedNodePools = append(refinedNodePools, nodepool)
			}
		}
	} else {
		refinedNodePools = nodepools
	}
	return refinedNodePools
}

func GetStructNestedField(v *NodePool, field string, structDriller []string) string {
	r := reflect.ValueOf(v)
	for _, drillField := range structDriller {
		f := reflect.Indirect(r).FieldByName(drillField).Interface()
		r = reflect.ValueOf(f)
	}
	f1 := reflect.Indirect(r).FieldByName(field)
	return f1.String()
}

type FilterStruct struct {
	Value   string
	Driller []string
}

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

// CreateOptsBuilder Create accepts a CreateOpts struct and uses the values to create a new
// logical Node Pool. When it is created, the Node Pool does not have an internal
// interface
type CreateOptsBuilder interface {
	ToNodePoolCreateMap() (map[string]interface{}, error)
}

// ToNodePoolCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToNodePoolCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node pool.
func Create(c *golangsdk.ServiceClient, clusterid string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodePoolCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{201}}
	_, r.Err = c.Post(rootURL(c, clusterid), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular node pool based on its unique ID and cluster ID.
func Get(c *golangsdk.ServiceClient, clusterid, nodepoolid string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, clusterid, nodepoolid), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

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
func Update(c *golangsdk.ServiceClient, clusterid, nodepoolid string, opts UpdateOpts) (r UpdateResult) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, clusterid, nodepoolid), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular node pool based on its unique ID and cluster ID.
func Delete(c *golangsdk.ServiceClient, clusterid, nodepoolid string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, clusterid, nodepoolid), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
