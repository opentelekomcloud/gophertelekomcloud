package addons

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
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

// Specifications to create an addon
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

type UpdateMetadata struct {
	// Add-on annotations in the format of key-value pairs.
	// For add-on upgrade, the value is fixed at {"addon.upgrade/type":"upgrade"}.
	Annotations UpdateAnnotations `json:"annotations" required:"true"`
	// Add-on labels in the format of key-value pairs.
	Labels map[string]string `json:"metadata,omitempty"`
}

type UpdateAnnotations struct {
	AddonUpdateType string `json:"addon.upgrade/type" required:"true"`
}

type UpdateOpts struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata required to create an addon
	Metadata UpdateMetadata `json:"metadata" required:"true"`
	// specifications to create an addon
	Spec RequestSpec `json:"spec" required:"true"`
}

type ListOpts struct {
	Name string `q:"addon_template_name"`
}

// WaitForAddonRunning - wait until addon status is `running`
func WaitForAddonRunning(client *golangsdk.ServiceClient, id, clusterID string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		addon, err := Get(client, id, clusterID).Extract()
		if err != nil {
			return false, fmt.Errorf("error retriving addon status: %w", err)
		}

		if addon.Status.Status == "running" {
			return true, nil
		}

		return false, nil
	})
}

// WaitForAddonDeleted - wait until addon is deleted
func WaitForAddonDeleted(client *golangsdk.ServiceClient, id, clusterID string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		_, err := Get(client, id, clusterID).Extract()

		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, fmt.Errorf("error retriving addon status: %w", err)
		}

		return false, nil
	})
}
