package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePolicyOpts struct {
	InstanceId               string `json:"-"`
	StartTime                string `json:"start_time" required:"true"`
	KeepDays                 int    `json:"keep_days" required:"true"`
	Period                   string `json:"period" required:"true"`
	RetentionNumBackupLevel1 *int   `json:"retention_num_backup_level-1,omitempty"`
}

func UpdatePolicy(client *golangsdk.ServiceClient, opts UpdatePolicyOpts) (*UpdatePolicyResponse, error) {
	b, err := build.RequestBody(opts, "backup_policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "backups", "policy", "update"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdatePolicyResponse
	return &res, extract.Into(raw.Body, &res)
}

type UpdatePolicyResponse struct {
	Status       string `json:"status"`
	InstanceId   string `json:"instance_id"`
	InstanceName string `json:"instance_name"`
}
