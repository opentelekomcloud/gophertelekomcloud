package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListDcsInstanceOpts struct {
	// Instance ID.
	Id string `q:"id"`
	// DCS instance name.
	Name string `q:"name"`
	// Number of DCS instances displayed on each page.
	// Minimum value: 1
	// Maximum value: 2000
	// If this parameter is left unspecified, a maximum of 1000 DCS instances are displayed on each page.
	Limit int `q:"limit"`
	// Start number for querying DCS instances. It cannot be lower than 1.
	// By default, the start number is 1.
	Start int `q:"start"`
	// DCS instance status.
	Status string `q:"status"`
	// An indicator of whether the number of DCS instances that failed to be created will be returned to the API caller.
	// Options:
	// true: The number of DCS instances that failed to be created will be returned to the API caller.
	// false or others: The number of DCS instances that failed to be created will not be returned to the API caller.
	IncludeFailure bool `q:"includeFailure"`
	// An indicator of whether to perform an exact or fuzzy match based on instance name.
	// Options:
	// true: exact match
	// false: fuzzy match
	// Default value: false.
	IsExactMatchName bool `q:"isExactMatchName"`
	// IP address for connecting to the DCS instance
	Ip string `q:"ip"`
}

func List(client *golangsdk.ServiceClient, opts ListDcsInstanceOpts) (*ListDcsResponse, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("instances")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListDcsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListDcsResponse struct {
	// Array of DCS instance details.
	Instances []Instance `json:"instances"`
	// Number of DCS instances.
	TotalCount int `json:"instance_num"`
}
