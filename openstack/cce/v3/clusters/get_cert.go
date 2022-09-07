package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetCert retrieves a particular cluster certificate based on its unique ID.
func GetCert(client *golangsdk.ServiceClient, id string) (*Certificate, error) {
	raw, err := client.Get(client.ServiceURL("clusters", id, "clustercert"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts,
	})
	if err != nil {
		return nil, err
	}

	var res Certificate
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Certificate struct {
	// API type, fixed value Config
	Kind string `json:"kind"`
	// API version, fixed value v1
	ApiVersion string `json:"apiVersion"`
	// Cluster list
	Clusters []CertClusters `json:"clusters"`
	// User list
	Users []CertUsers `json:"users"`
	// Context list
	Contexts []CertContexts `json:"contexts"`
	// The current context
	CurrentContext string `json:"current-context"`
}
