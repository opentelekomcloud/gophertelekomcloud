package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToBackupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a Backup.
type UpdateOpts struct {
	// Keep Days
	KeepDays *int `json:"keep_days" required:"true"`
	// Start Time
	StartTime string `json:"start_time,omitempty"`
	// Period
	Period string `json:"period,omitempty"`
}

// ToBackupUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToBackupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "backup_policy")
}

// Update accepts a UpdateOpts struct and uses the values to update a Backup.The response code from api is 200
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToBackupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := c.Put(c.ServiceURL("instances", id, "backups/policy"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}, MoreHeaders: openstack.StdRequestOpts().MoreHeaders})
	return
}
