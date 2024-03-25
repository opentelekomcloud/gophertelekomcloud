package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func DeleteDefaultConfig(client *golangsdk.ServiceClient) (*TaskResponse, error) {
	// DELETE /v1/{project_id}/antiddos/default-config
	raw, err := client.Delete(client.ServiceURL("antiddos", "default-config"), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res TaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
