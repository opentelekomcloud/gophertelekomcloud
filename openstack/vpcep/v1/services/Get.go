package services

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Service, error) {
	// GET /v1/{project_id}/vpc-endpoint-services/{vpc_endpoint_service_id}
	raw, err := client.Get(client.ServiceURL("vpc-endpoint-services", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Service
	return &res, extract.Into(raw.Body, &res)
}
