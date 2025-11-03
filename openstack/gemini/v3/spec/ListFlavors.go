package spec

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListFlavors(client *golangsdk.ServiceClient, opts ListFlavorsOpts) (*ListFlavorsResponse, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("flavors")+query.String(), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListFlavorsResponse
	return &res, extract.Into(raw.Body, &res)
}

type ListFlavorsOpts struct {
	EngineName string `q:"engine_name"`
	Offset     *int   `q:"offset"`
	Limit      *int   `q:"limit"`
}

type ListFlavorsResponse struct {
	TotalCount int       `json:"total_count"`
	Flavors    []Flavors `json:"flavors"`
}

type Flavors struct {
	EngineName       string            `json:"engine_name"`
	EngineVersion    string            `json:"engine_version"`
	Vcpus            string            `json:"vcpus"`
	Ram              string            `json:"ram"`
	SpecCode         string            `json:"spec_code"`
	AvailabilityZone []string          `json:"availability_zone"`
	AzStatus         map[string]string `json:"az_status"`
}
