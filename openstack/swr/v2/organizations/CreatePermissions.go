package organizations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type Auth struct {
	// User ID, which needs to be obtained from the IAM service.
	UserID string `json:"user_id"`
	// Username, which needs to be obtained from the IAM service.
	Username string `json:"user_name"`
	// User permission
	// 7: Manage
	// 3: Write
	// 1: Read
	Auth int `json:"auth"`
}

func CreatePermissions(client *golangsdk.ServiceClient, organization string, opts []Auth) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v2/manage/namespaces/{namespace}/access
	_, err = client.Post(client.ServiceURL("manage", "namespaces", organization, "access"), b, nil, nil)
	return
}
