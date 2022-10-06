package sync

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/drs/v3/public"
)

type SwitchVIPOpts struct {
	// Switchover type. Values:
	// source: Switch over the virtual IP address of the source database.
	// target: Switch over the virtual IP address of the destination database.
	Mode string `json:"mode"`
}

func SwitchVIP(client *golangsdk.ServiceClient, jobId string) (*public.IdJobResp, error) {
	// POST  /v3/{project_id}/jobs/{job_id}/switch-vip
	raw, err := client.Post(client.ServiceURL("jobs", jobId, "switch-vip"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var res public.IdJobResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
