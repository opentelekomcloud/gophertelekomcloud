package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifySGOpt struct {
	SecurityGroupId string `json:"security_group_id" required:"true"`
	InstanceId      string `json:"-"`
}

func ModifySG(client *golangsdk.ServiceClient, opts ModifySGOpt) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "modify-security-group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
