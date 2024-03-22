package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangePasswordOpts struct {
	UserId           string `json:"-"`
	OriginalPassword string `json:"original_password"`
	NewPassword      string `json:"password"`
}

func ChangePassword(client *golangsdk.ServiceClient, opts ChangePasswordOpts) error {
	b, err := build.RequestBody(opts, "user")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("users", opts.UserId, "password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return err
	}

	return nil
}
