package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetRecyclePolicy(client *golangsdk.ServiceClient) (*RecyclePolicy, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/recycle-policy
	raw, err := client.Get(client.ServiceURL("instances", "recycle-policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res RecyclePolicy
	err = extract.Into(raw.Body, &res)
	return &res, err
}
