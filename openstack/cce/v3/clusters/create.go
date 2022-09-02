package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

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

// Create accepts a CreateOpts struct and uses the values to create a new logical cluster.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Clusters, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	if err != nil {
		return nil, err
	}

	var res Clusters
	err = extract.Into(raw, &res)
	return &res, err
}
