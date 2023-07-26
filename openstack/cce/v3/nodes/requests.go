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
	var r ListResult
	_, r.Err = client.Get(rootURL(client, clusterID), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	allNodes, err := r.ExtractNode()

	if err != nil {
		return nil, err
	}

	return FilterNodes(allNodes, opts), nil
}

func FilterNodes(nodes []Nodes, opts ListOpts) []Nodes {

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
	ToNodeCreateMap() (map[string]any, error)
}

// ToNodeCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToNodeCreateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node.
func Create(c *golangsdk.ServiceClient, clusterID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{201}}
	_, r.Err = c.Post(rootURL(c, clusterID), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(c *golangsdk.ServiceClient, clusterID, nodeID string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, clusterID, nodeID), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNodeUpdateMap() (map[string]any, error)
}

// UpdateOpts contains all the values needed to update a new node
type UpdateOpts struct {
	Metadata UpdateMetadata `json:"metadata,omitempty"`
}

type UpdateMetadata struct {
	Name string `json:"name,omitempty"`
}

// ToNodeUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToNodeUpdateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update allows nodes to be updated.
func Update(c *golangsdk.ServiceClient, clusterID, nodeID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToNodeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, clusterID, nodeID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular node based on its unique ID and cluster ID.
func Delete(c *golangsdk.ServiceClient, clusterID, nodeID string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, clusterID, nodeID), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

// GetJobDetails retrieves a particular job based on its unique ID
func GetJobDetails(c *golangsdk.ServiceClient, jobID string) (r GetResult) {
	_, r.Err = c.Get(getJobURL(c, jobID), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
