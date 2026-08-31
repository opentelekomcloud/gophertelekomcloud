package subnets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Subnet, error) {
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "subnets", id), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Subnet Subnet `json:"subnet"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Subnet, err
}
