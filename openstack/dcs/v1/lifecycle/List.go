package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
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
	// An indicator of whether to perform an exact or fuzzy match based on instance name.
	// Options:
	// true: exact match
	// false: fuzzy match
	// Default value: false.
	IsExactMatchName bool `q:"isExactMatchName"`
	// IP address for connecting to the DCS instance
	Ip string `q:"ip"`
	// Query based on the instance tag key and value. {key} indicates the tag key, and {value} indicates the tag value.
	// To query instances with multiple tag keys and values, separate key-value pairs with commas (,).
	Tags map[string]string `q:"tags"`
}

func List(client *golangsdk.ServiceClient, opts ListDcsInstanceOpts) (*ListDcsResponse, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v1.0/{project_id}/instances
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
