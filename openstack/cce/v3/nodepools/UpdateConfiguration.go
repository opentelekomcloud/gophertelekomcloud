package nodepools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
)

type UpdateConfigurationOpts struct {
	// API type. Must be "Configuration".
	Kind string `json:"kind" required:"true"`
	// API version. Must be "v3".
	APIVersion string `json:"apiVersion" required:"true"`
	// Configuration metadata.
	Metadata ConfigurationMetadata `json:"metadata" required:"true"`
	// Configuration specifications.
	Spec ClusterConfigurationsSpec `json:"spec" required:"true"`
}

type ConfigurationMetadata struct {
	// Configuration name (e.g., "configuration")
	Name string `json:"name" required:"true"`
	// Optional labels map.
	Labels map[string]string `json:"labels,omitempty"`
}

type ClusterConfigurationsSpec struct {
	// Component configuration item details.
	Packages []clusters.PackageConfiguration `json:"packages" required:"true"`
}

// UpdateConfiguration changes the values of configuration parameters of a node pool.
func UpdateConfiguration(client *golangsdk.ServiceClient, clusterId, nodepoolId string, opts UpdateConfigurationOpts) (*Configuration, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// PUT /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools/{nodepool_id}/configuration
	raw, err := client.Put(
		client.ServiceURL("clusters", clusterId, "nodepools", nodepoolId, "configuration"),
		b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}},
	)
	if err != nil {
		return nil, err
	}

	var res Configuration
	return &res, extract.Into(raw.Body, &res)
}

// Configuration represents the response body returned after updating configuration.
type Configuration struct {
	Kind       string                    `json:"kind"`
	APIVersion string                    `json:"apiVersion"`
	Metadata   ConfigurationMetadata     `json:"metadata"`
	Spec       ClusterConfigurationsSpec `json:"spec"`
	Status     map[string]interface{}    `json:"status,omitempty"`
}
