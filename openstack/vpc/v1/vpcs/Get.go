package vpcs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Vpc, error) {
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "vpcs", id), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		VPC Vpc `json:"vpc"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.VPC, err
}
