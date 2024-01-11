package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetUser(client *golangsdk.ServiceClient, id string) (*User, error) {
	// GET /v3.0/OS-USER/users/{user_id}
	raw, err := client.Get(client.ServiceURL("OS-USER", "users", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res User
	err = extract.IntoStructPtr(raw.Body, &res, "user")
	return &res, err
}
