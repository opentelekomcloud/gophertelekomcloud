package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ModifyUser(client *golangsdk.ServiceClient, id string, opts CreateOpts) (*User, error) {
	b, err := build.RequestBody(opts, "user")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-USER/users/{user_id}
	raw, err := client.Put(client.ServiceURL("OS-USER", "users", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res User
	return &res, extract.IntoStructPtr(raw.Body, &res, "user")
}
