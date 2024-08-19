package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
)

type ModifyBackupPolicyOpts struct {
	InstanceId   string                    `json:"-"`
	BackupPolicy *instances.BackupStrategy `json:"backup_policy" required:"true"`
}

func SetBackupPolicy(client *golangsdk.ServiceClient, opts ModifyBackupPolicyOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/policy
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "backups", "policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
