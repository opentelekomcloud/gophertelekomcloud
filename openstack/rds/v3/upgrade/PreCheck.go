package upgrade

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type PreCheckOpts struct {
	InstanceId    string `json:"-"`
	TargetVersion string `json:"target_version"`
}

func PreCheck(client *golangsdk.ServiceClient, opts PreCheckOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "major-version", "inspection"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ReportId string `json:"report_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ReportId, err
}
