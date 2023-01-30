package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangePasswordOpt struct {
	UserPwd    string `json:"user_pwd" required:"true"`
	InstanceId string `json:"-"`
}

func ChangePassword(client *golangsdk.ServiceClient, opts ChangePasswordOpt) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "reset-password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return err
	}
	return nil
}
