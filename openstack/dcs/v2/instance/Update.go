package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifyInstanceOpt struct {
	InstanceId      string                    `json:"-"`
	Name            string                    `json:"name,omitempty"`
	Description     *string                   `json:"description,omitempty"`
	Port            *int                      `json:"port,omitempty"`
	MaintainBegin   string                    `json:"maintain_begin,omitempty"`
	MaintainEnd     string                    `json:"maintain_end,omitempty"`
	SecurityGroupId *string                   `json:"security_group_id,omitempty"`
	BackupPolicy    *InstanceBackupPolicyOpts `json:"instance_backup_policy,omitempty"`
	RenameCommands  *RenameCommand            `json:"rename_commands,omitempty"`
}

type RenameCommand struct {
	Command  string `json:"command,omitempty"`
	Keys     string `json:"keys,omitempty"`
	Flushdb  string `json:"flushdb,omitempty"`
	Flushall string `json:"flushall,omitempty"`
	Hgetall  string `json:"hgetall,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts ModifyInstanceOpt) (err error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
