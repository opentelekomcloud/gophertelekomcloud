package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, instanceId string) ([]Users, error) {
	// GET /v2/{project_id}/instances/{instance_id}/users

	url := client.ServiceURL("instances", instanceId, "users")
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Users
	err = extract.IntoSlicePtr(raw.Body, &res, "users")
	return res, err
}

type Users struct {
	UserName    string  `json:"user_name"`
	Role        string  `json:"role"`
	DefaultApp  bool    `json:"default_app"`
	CreatedTime float64 `json:"created_time"`
}
