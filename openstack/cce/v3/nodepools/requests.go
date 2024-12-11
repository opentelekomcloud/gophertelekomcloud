package nodepools

import (
	"reflect"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
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
