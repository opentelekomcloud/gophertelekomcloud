package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List sends GET /v2/{project_id}/kafka/tags
func List(client *golangsdk.ServiceClient) (*ListResponse, error) {
	raw, err := client.Get(client.ServiceURL("kafka", "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Tags []ProjectTag `json:"tags"`
}

type ProjectTag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}
