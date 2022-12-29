package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListFlavorOpts struct {
	Region        string `q:"region"`
	EngineName    string `q:"engine_name"`
	EngineVersion string `q:"engine_version"`
}

func List(client *golangsdk.ServiceClient, opts ListFlavorOpts) (*ListResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("flavors")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Flavors    []FlavorResponse `json:"flavors"`
	TotalCount int              `json:"total_count"`
}

type FlavorResponse struct {
	EngineName     string            `json:"engine_name"`
	Type           string            `json:"type"`
	Vcpus          string            `json:"vcpus"`
	Ram            string            `json:"ram"`
	SpecCode       string            `json:"spec_code"`
	AZStatus       map[string]string `json:"az_status"`
	EngineVersions []string          `json:"engine_versions"`
}
