package clusters

import (
	"reflect"
)

var RequestOpts = map[string]string{"Content-Type": "application/json"}

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

// Metadata required to create a cluster
type CreateMetaData struct {
	// Cluster unique name
	Name string `json:"name" required:"true"`
	// Cluster tag, key/value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Cluster annotation, key/value pair format
	Annotations map[string]string `json:"annotations,omitempty"`
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
