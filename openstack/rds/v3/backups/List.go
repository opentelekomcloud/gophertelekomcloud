package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToBackupListQuery() (string, error)
}

type ListOpts struct {
	//
	InstanceID string `q:"instance_id"`
	//
	BackupID string `q:"backup_id"`
	//
	BackupType string `q:"backup_type"`
	//
	BeginTime string `q:"begin_time"`
	//
	EndTime string `q:"end_time"`
}

func (opts ListOpts) ToBackupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := c.ServiceURL("backups")
	if opts != nil {
		q, err := opts.ToBackupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return BackupPage{SinglePageBase: pagination.SinglePageBase(r)}
	})
}
