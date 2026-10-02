package features

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListResponse is returned by List.
type ListResponse struct {
	// Feature list.
	Features []Feature `json:"features"`
	// Total number of features.
	TotalRecord int `json:"totalRecord"`
}

type Feature struct {
	// Feature ID.
	FeatureID string `json:"featureId"`
	// Status: 1 if the feature is enabled, 0 otherwise.
	Status int `json:"status"`
	// Feature description.
	Description string `json:"description"`
}

// List the DMS feature switches.
// Send GET to /v2/config/features
func List(client *golangsdk.ServiceClient) (*ListResponse, error) {
	url := strings.Split(client.Endpoint, "/v2/")[0] + "/v2/config/features"

	raw, err := client.Get(url, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
