package endpoints

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Endpoint, error) {
	// GET /v1/{project_id}/vpc-endpoints/{vpc_endpoint_id}
	raw, err := client.Get(client.ServiceURL("vpc-endpoints", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Endpoint
	return &res, extract.Into(raw.Body, &res)
}
