package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetAutoExpansion(client *golangsdk.ServiceClient, instanceId string) (*GetAutoExpansionResponse, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "disk-auto-expansion"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res GetAutoExpansionResponse
	return &res, extract.Into(raw.Body, &res)
}

type GetAutoExpansionResponse struct {
	Policy *AutoEnlargePolicy `json:"policy"`
}

type AutoEnlargePolicy struct {
	Threshold int `json:"threshold"`
	Step      int `json:"step"`
	Size      int `json:"size"`
}
