package vpc_endpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List VPC Endpoint Connections
func ListConnections(client *golangsdk.ServiceClient, id string) ([]ConnectionResp, error) {
	raw, err := client.Get(client.ServiceURL("clusters", id, "vpcepservice", "connections"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ConnectionResp
	err = extract.IntoSlicePtr(raw.Body, &res, "connections")
	return res, err
}

type ConnectionResp struct {
	ID                string `json:"id"`
	Status            string `json:"status"`
	MaxSession        string `json:"maxSession"`
	SpecificationName string `json:"specificationName"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"update_at"`
	DomainId          string `json:"domain_id"`
}
