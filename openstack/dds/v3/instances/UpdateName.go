package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateNameOpt struct {
	NewInstanceName string `json:"new_instance_name" required:"true"`
	InstanceId      string `json:"-"`
}

func UpdateName(client *golangsdk.ServiceClient, opts UpdateNameOpt) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "modify-name"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return err
	}
	return nil
}
