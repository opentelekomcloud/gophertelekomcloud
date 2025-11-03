package spec

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetVersions(client *golangsdk.ServiceClient, datastoreName string) (*GetVersionsResponse, error) {
	raw, err := client.Get(client.ServiceURL("datastores", datastoreName, "versions"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res GetVersionsResponse
	return &res, extract.Into(raw.Body, &res)
}

type GetVersionsResponse struct {
	Versions []string `json:"versions"`
}
