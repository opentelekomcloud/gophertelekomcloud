package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List sends GET /v2/{project_id}/instances/{instance_id}/users
func List(client *golangsdk.ServiceClient, instanceId string) ([]Users, error) {
	if instanceId == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceId"}
	}

	raw, err := client.Get(client.ServiceURL("instances", instanceId, "users"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Users
	err = extract.IntoSlicePtr(raw.Body, &res, "users")
	return res, err
}

type Users struct {
	UserName    string `json:"user_name"`
	UserDesc    string `json:"user_desc"`
	Role        string `json:"role"`
	DefaultApp  bool   `json:"default_app"`
	CreatedTime int64  `json:"created_time"`
}
