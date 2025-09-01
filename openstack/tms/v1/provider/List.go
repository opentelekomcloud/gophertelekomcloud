package provider

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts is the structure that used to query tags detail for specified resource.
type ListOpts struct {
	// Specifies the display language.
	Locale string `q:"locale"`
	// The maximum queries supported. The value 10 is used by default if this parameter is not set.
	// The value range is 1 to 200.
	Limit *int `q:"limit"`
	// Specifies the index position, which starts from the next data record specified by offset.
	// The value must be a number and cannot be negative. The default value is 0.
	Offset *int `q:"offset"`
	// Specifies the cloud service name.
	Provider *int `q:"provider"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Providers, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("tms", "providers").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v1.0/tms/providers
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []Providers
	err = extract.IntoSlicePtr(raw.Body, &res, "providers")
	return res, err
}

type Providers struct {
	// Specifies the cloud service name.
	Provider string `json:"provider"`
	// Specifies the display name of the resource.
	// You can set the language by setting the **locale** parameter.
	DisplayName string `json:"provider_i18n_display_name"`
	// Specifies the resource type.
	ResourceTypes []ResourceTypes `json:"resource_types"`
}

type ResourceTypes struct {
	// Specifies the resource type.
	ResourceType string `json:"resource_type"`
	// Specifies the display name of the resource type.
	// You can set the language by setting the **locale** parameter.
	DisplayName string `json:"provider_i18n_display_name"`
	// Specifies the supported regions.
	Regions []string `json:"regions"`
	// Specifies whether the resource is a global resource.
	Global bool `json:"global"`
}
