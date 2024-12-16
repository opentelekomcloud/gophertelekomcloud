package addons

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

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

func Update(client *golangsdk.ServiceClient, addonId string, clusterId string, opts UpdateOpts) (*Addon, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("addons", addonId).WithQueryParams(&ClusterIdQueryParam{ClusterId: clusterId}).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(CCEServiceURL(client, clusterId, url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Addon
	return &res, extract.Into(raw.Body, &res)
}
