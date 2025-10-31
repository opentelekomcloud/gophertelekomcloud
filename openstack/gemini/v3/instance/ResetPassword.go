package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResetPasswordOpts struct {
	InstanceId string `json:"-"`
	Password   string `json:"password" required:"true"`
}

func ResetPassword(client *golangsdk.ServiceClient, opts ResetPasswordOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
