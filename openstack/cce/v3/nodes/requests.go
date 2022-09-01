package nodes

import (
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
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

// List returns collection of nodes.
func List(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) ([]Nodes, error) {
	var res ListResult
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodes"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	allNodes, err := res.ExtractNode()

	if err != nil {
		return nil, err
	}

	return filterNodes(allNodes, opts), nil
}

func filterNodes(nodes []Nodes, opts ListOpts) []Nodes {
	var refinedNodes []Nodes
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

	if len(m) > 0 && len(nodes) > 0 {
		for _, nodes := range nodes {
			matched = true

			for key, value := range m {
				if sVal := GetStructNestedField(&nodes, key, value.Driller); !(sVal == value.Value) {
					matched = false
				}
			}
			if matched {
				refinedNodes = append(refinedNodes, nodes)
			}
		}
	} else {
		refinedNodes = nodes
	}
	return refinedNodes
}

func GetStructNestedField(v *Nodes, field string, structDriller []string) string {
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

// CreateOpts is a struct contains the parameters of creating Node
type CreateOpts struct {
	// API type, fixed value Node
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata required to create a Node
	Metadata CreateMetaData `json:"metadata"`
	// specifications to create a Node
	Spec Spec `json:"spec" required:"true"`
}

// CreateMetaData required to create a Node
type CreateMetaData struct {
	// Node name
	Name string `json:"name,omitempty"`
	// Node tag, key value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Node annotation, key value pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToNodeCreateMap() (map[string]interface{}, error)
}

// ToNodeCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToNodeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node.
func Create(client *golangsdk.ServiceClient, clusterID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodeCreateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters", clusterID, "nodes"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return
}

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterID, nodeID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodes", nodeID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNodeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a new node
type UpdateOpts struct {
	Metadata UpdateMetadata `json:"metadata,omitempty"`
}

type UpdateMetadata struct {
	Name string `json:"name,omitempty"`
}

// ToNodeUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToNodeUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update allows nodes to be updated.
func Update(client *golangsdk.ServiceClient, clusterID, nodeID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToNodeUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", clusterID, "nodes", nodeID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular node based on its unique ID and cluster ID.
func Delete(client *golangsdk.ServiceClient, clusterID, nodeID string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("clusters", clusterID, "nodes", nodeID), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

// GetJobDetails retrieves a particular job based on its unique ID
func GetJobDetails(client *golangsdk.ServiceClient, jobID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("jobs", jobID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
