package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetWhitelist(client *golangsdk.ServiceClient, id string) (*GetWhitelistResponse, error) {
	// POST /v1/{project_id}/vpc-endpoint-services/{vpc_endpoint_service_id}/permissions
	raw, err := client.Get(client.ServiceURL("vpc-endpoint-services", id, "permissions"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetWhitelistResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetWhitelistResponse struct {
	Permissions []Permission `json:"permissions"`
	TotalCount  int          `json:"total_count"`
}

type Permission struct {
	Id         string `json:"id"`
	Permission string `json:"permission"`
	CreatedAt  string `json:"created_at"`
}
