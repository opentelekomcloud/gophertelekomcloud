package upgrade

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpgradeMajorVersionOpts struct {
	InstanceId               string `json:"-"`
	TargetVersion            string `json:"target_version"`
	IsChangePrivateIp        bool   `json:"is_change_private_ip"`
	StatisticsCollectionMode string `json:"statistics_collection_mode,omitempty"`
}

type JobResponse struct {
	JobId string `json:"job_id"`
}

func UpgradeMajorVersion(client *golangsdk.ServiceClient, opts UpgradeMajorVersionOpts) (*JobResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "major-version", "upgrade"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res JobResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
