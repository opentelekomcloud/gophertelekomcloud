package sync

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/drs/v3/public"
)

// GetSwitchVIPStatus Status Values:
// SWITCH_VIP_COMPLETE: The switchover is successful.
// SWITCH_VIP_FAILED: The switchover failed.
// SWITCH_VIP_START: The switchover is in progress.
func GetSwitchVIPStatus(client *golangsdk.ServiceClient, jobId string) (*public.IdJobResp, error) {
	// GET /v3/{project_id}/jobs/{job_id}/get-switch-vip-status
	raw, err := client.Get(client.ServiceURL("jobs", jobId, "get-switch-vip-status"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res public.IdJobResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
