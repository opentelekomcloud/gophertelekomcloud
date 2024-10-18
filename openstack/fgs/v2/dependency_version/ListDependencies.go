package dependency_version

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListDependenciesOpts struct {
	Marker         string `q:"marker"`
	MaxItems       string `q:"max_items"`
	IsPublic       string `json:"is_public"`
	DependencyType string `json:"dependency_type"`
	Runtime        string `json:"runtime"`
	Name           string `json:"name"`
	Limit          string `json:"limit"`
}

func ListDependencies(client *golangsdk.ServiceClient, opts ListDependenciesOpts) (*ListDepVersionResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "dependencies").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListDepVersionResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
