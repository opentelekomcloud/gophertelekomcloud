package clusters

import (
	"reflect"
)

var RequestOpts = map[string]string{"Content-Type": "application/json"}

// ListOpts allows the filtering of list data using given parameters.
type ListOpts struct {
	Name  string `json:"name"`
	ID    string `json:"uuid"`
	Type  string `json:"type"`
	VpcID string `json:"vpc"`
	Phase string `json:"phase"`
}

func filterClusters(clusters []Clusters, opts ListOpts) []Clusters {
	var refinedClusters []Clusters
	var matched bool
	m := map[string]FilterStruct{}

	if opts.Name != "" {
		m["Name"] = FilterStruct{Value: opts.Name, Driller: []string{"Metadata"}}
	}
	if opts.ID != "" {
		m["Id"] = FilterStruct{Value: opts.ID, Driller: []string{"Metadata"}}
	}
	if opts.Type != "" {
		m["Type"] = FilterStruct{Value: opts.Type, Driller: []string{"Spec"}}
	}
	if opts.VpcID != "" {
		m["VpcId"] = FilterStruct{Value: opts.VpcID, Driller: []string{"Spec", "HostNetwork"}}
	}
	if opts.Phase != "" {
		m["Phase"] = FilterStruct{Value: opts.Phase, Driller: []string{"Status"}}
	}

	if len(m) > 0 && len(clusters) > 0 {
		for _, cluster := range clusters {
			matched = true

			for key, value := range m {
				if sVal := GetStructNestedField(&cluster, key, value.Driller); !(sVal == value.Value) {
					matched = false
				}
			}
			if matched {
				refinedClusters = append(refinedClusters, cluster)
			}
		}

	} else {
		refinedClusters = clusters
	}

	return refinedClusters
}

type FilterStruct struct {
	Value   string
	Driller []string
}

func GetStructNestedField(v *Clusters, field string, structDriller []string) string {
	r := reflect.ValueOf(v)
	for _, drillField := range structDriller {
		f := reflect.Indirect(r).FieldByName(drillField).Interface()
		r = reflect.ValueOf(f)
	}
	f1 := reflect.Indirect(r).FieldByName(field)
	return f1.String()
}

// CreateOpts contains all the values needed to create a new cluster
type CreateOpts struct {
	// API type, fixed value Cluster
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiversion" required:"true"`
	// Metadata required to create a cluster
	Metadata CreateMetaData `json:"metadata" required:"true"`
	// specifications to create a cluster
	Spec Spec `json:"spec" required:"true"`
}

// Metadata required to create a cluster
type CreateMetaData struct {
	// Cluster unique name
	Name string `json:"name" required:"true"`
	// Cluster tag, key/value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Cluster annotation, key/value pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ExpirationOpts struct {
	Duration int `json:"duration" required:"true"`
}

// UpdateOpts contains all the values needed to update a new cluster
type UpdateOpts struct {
	Spec UpdateSpec `json:"spec" required:"true"`
}

type UpdateSpec struct {
	// Cluster description
	Description string `json:"description,omitempty"`
}

type UpdateIpOpts struct {
	Action    string `json:"action" required:"true"`
	Spec      IpSpec `json:"spec,omitempty"`
	ElasticIp string `json:"elasticIp"`
}

type IpSpec struct {
	ID string `json:"id" required:"true"`
}
