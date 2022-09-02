package addons

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains all the values needed to create a new addon
type CreateOpts struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata required to create an addon
	Metadata CreateMetadata `json:"metadata" required:"true"`
	// specifications to create an addon
	Spec RequestSpec `json:"spec" required:"true"`
}

type CreateMetadata struct {
	Annotations CreateAnnotations `json:"annotations" required:"true"`
}

type CreateAnnotations struct {
	AddonInstallType string `json:"addon.install/type" required:"true"`
}

// RequestSpec to create an addon
type RequestSpec struct {
	// For the addon version.
	Version string `json:"version" required:"true"`
	// Cluster ID.
	ClusterID string `json:"clusterID" required:"true"`
	// Addon Template Name.
	AddonTemplateName string `json:"addonTemplateName" required:"true"`
	// Addon Parameters
	Values Values `json:"values" required:"true"`
}

type Values struct {
	Basic    map[string]interface{} `json:"basic" required:"true"`
	Advanced map[string]interface{} `json:"custom,omitempty"`
}

// Create accepts a CreateOpts struct and uses the values to create a new addon.
func Create(client *golangsdk.ServiceClient, opts CreateOpts, clusterId string) (*Addon, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(fmt.Sprintf("https://%s.%s", clusterId, client.ResourceBaseURL()[8:])+
		strings.Join([]string{"addons"}, "/"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	if err != nil {
		return nil, err
	}

	var res Addon
	err = extract.Into(raw, &res)
	return &res, err
}
