package organizations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Organization, error) {
	// GET /v2/manage/namespaces/{namespace}
	raw, err := client.Get(client.ServiceURL("manage", "namespaces", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Organization
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Organization struct {
	// Organization ID
	ID int `json:"id"`
	// Organization name
	Name string `json:"name"`
	// IAM username
	CreatorName string `json:"creator_name"`
	// User permission
	// 7: Manage
	// 3: Write
	// 1: Read
	Auth int `json:"auth"`
}
