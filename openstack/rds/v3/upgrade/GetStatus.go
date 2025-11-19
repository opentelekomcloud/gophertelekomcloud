package upgrade

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetStatusOpts struct {
	InstanceId string `json:"-"`
	Action     string `json:"-"`
}

type VersionStatus struct {
	Status              string `json:"status"`
	TargetVersion       string `json:"target_version"`
	StartTime           string `json:"start_time"`
	CheckExpirationTime string `json:"check_expiration_time"`
	Detail              string `json:"detail"`
}

func GetStatus(client *golangsdk.ServiceClient, opts GetStatusOpts) (*VersionStatus, error) {
	url := client.ServiceURL("instances", opts.InstanceId, "major-version", "status")
	url += "?action=" + opts.Action

	raw, err := client.Get(url, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res VersionStatus
	err = extract.Into(raw.Body, &res)
	return &res, err
}
