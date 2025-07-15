package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePasswordOpts struct {
	InstanceId  string `json:"-"`
	OldPassword string `json:"old_password" required:"true"`
	NewPassword string `json:"new_password" required:"true"`
}

func UpdatePassword(client *golangsdk.ServiceClient, opts UpdatePasswordOpts) (*UpdatePasswordResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdatePasswordResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdatePasswordResponse struct {
	LockTime       string `json:"lock_time"`
	Result         string `json:"result"`
	LockTimeLeft   string `json:"lock_time_left"`
	RetryTimesLeft string `json:"retry_times_left"`
	Message        string `json:"message"`
}
