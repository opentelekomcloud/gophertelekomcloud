package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeSecGroupOpts struct {
	InstanceId      string `json:"-"`
	SecurityGroupId string `json:"security_group_id" required:"true"`
}

func ChangeSecGroup(client *golangsdk.ServiceClient, opts ChangeSecGroupOpts) (*ChangeSecGroupResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "security-group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res ChangeSecGroupResp
	return &res, extract.Into(raw.Body, &res)
}

type ChangeSecGroupResp struct {
	JobId string `json:"job_id"`
}
