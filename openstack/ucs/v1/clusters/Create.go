package clusters

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is the request body of POST /v1/clusters (registering_a_cluster.rst, Table 2).
type CreateOpts struct {
	Kind       string         `json:"kind" required:"true"`
	APIVersion string         `json:"apiVersion" required:"true"`
	Metadata   CreateMetadata `json:"metadata" required:"true"`
	Spec       CreateSpec     `json:"spec" required:"true"`
}

// CreateMetadata is the request metadata object (Table 3).
type CreateMetadata struct {
	UID         string            `json:"uid,omitempty"`
	Name        string            `json:"name" required:"true"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// CreateSpec is the request spec object (Table 4).
// Note: in the request, provider is a map (e.g. {"CCE":"cce"}), unlike the response.
type CreateSpec struct {
	ClusterGroupID string            `json:"clusterGroupID,omitempty"`
	Category       string            `json:"category" required:"true"`
	Type           string            `json:"type" required:"true"`
	Provider       map[string]string `json:"provider" required:"true"`
	Country        string            `json:"country" required:"true"`
	City           string            `json:"city,omitempty"`
	Region         string            `json:"region,omitempty"`
	ProjectID      string            `json:"projectID,omitempty"`
	ManageType     string            `json:"manageType" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{http.StatusCreated}, // 201
	})
	if err != nil {
		return "", err
	}

	var res struct {
		UID string `json:"uid"`
	}
	return res.UID, extract.Into(raw.Body, &res)
}
