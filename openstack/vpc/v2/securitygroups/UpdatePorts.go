package securitygroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePortsOpts struct {
	Ports  []Port `json:"ports" required:"true"`
	Action string `json:"action" required:"true"`
}

type Port struct {
	ID string `json:"id" required:"true"`
}

type FailedPort struct {
	ID        string `json:"id"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func UpdatePorts(client *golangsdk.ServiceClient, securityGroupID string, opts UpdatePortsOpts) ([]FailedPort, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "security-groups", securityGroupID, "instance", "action").
		Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		Fail []FailedPort `json:"fail"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return res.Fail, nil
}
