package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

type CreateOptsBuilder interface {
	ToBackupCreateMap() (map[string]interface{}, error)
}

type BackupDatabase struct {
	Name string `json:"name"`
}

type CreateOpts struct {
	InstanceID  string           `json:"instance_id" required:"true"`
	Name        string           `json:"name" required:"true"`
	Description string           `json:"description,omitempty"`
	Databases   []BackupDatabase `json:"databases,omitempty"`
}

func (opts CreateOpts) ToBackupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToBackupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(c.ServiceURL("backups"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}
