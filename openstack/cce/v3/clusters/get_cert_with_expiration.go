package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ExpirationOpts struct {
	Duration int `json:"duration" required:"true"`
}

// GetCertWithExpiration retrieves a particular cluster certificate based on its unique ID.
func GetCertWithExpiration(client *golangsdk.ServiceClient, id string, opts ExpirationOpts) (*Certificate, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters", id, "clustercert"), b, nil, &golangsdk.RequestOpts{
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

type CertClusters struct {
	// Cluster name
	Name string `json:"name"`
	// Cluster information
	Cluster CertCluster `json:"cluster"`
}

type CertCluster struct {
	// Server IP address
	Server string `json:"server"`
	// Certificate data
	CertAuthorityData string `json:"certificate-authority-data"`
}

type CertUsers struct {
	// User name
	Name string `json:"name"`
	// Cluster information
	User CertUser `json:"user"`
}

type CertUser struct {
	// Client certificate
	ClientCertData string `json:"client-certificate-data"`
	// Client key data
	ClientKeyData string `json:"client-key-data"`
}

type CertContexts struct {
	// Context name
	Name string `json:"name"`
	// Context information
	Context CertContext `json:"context"`
}

type CertContext struct {
	// Cluster name
	Cluster string `json:"cluster"`
	// User name
	User string `json:"user"`
}
