package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

type RestorePITROpts struct {
	//
	Source Source `json:"source"`
	//
	Target Target `json:"target"`
}

type Source struct {
	//
	BackupID string `json:"backup_id" required:"false"`
	//
	InstanceID string `json:"instance_id" required:"true"`
	//
	RestoreTime int64 `json:"restore_time" required:"false"`
	//
	Type string `json:"type" required:"true"`
}

type Target struct {
	//
	InstanceID string `json:"instance_id" required:"true"`
}

type RestorePITROptsBuilder interface {
	ToPITRRestoreMap() (map[string]interface{}, error)
}

func (opts RestorePITROpts) ToPITRRestoreMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func RestorePITR(c *golangsdk.ServiceClient, opts RestorePITROptsBuilder) (r RestoreResult) {
	b, err := opts.ToPITRRestoreMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := c.Post(c.ServiceURL("instances", "recovery"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
