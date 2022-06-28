package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, id), &r.Body, nil)
	return
}

type ListOptsBuilder interface {
	ToBackupListQuery() (string, error)
}

type ImageType string
type MemberStatus string
type OwnType string
type Status string

type ListOpts struct {
	CheckpointId   string       `q:"checkpoint_id"`
	DedicatedCloud bool         `q:"dec"`
	EndTime        string       `q:"end_time"`
	ImageType      ImageType    `q:"image_type"`
	MemberStatus   MemberStatus `q:"member_status"`
	Name           string       `q:"name"`
	OwnType        OwnType      `q:"own_type"`
	ParentId       string       `q:"parent_id"`
	ResourceAZ     string       `q:"resource_az"`
	ResourceID     string       `q:"resource_id"`
	ResourceName   string       `q:"resource_name"`
	ResourceType   string       `q:"resource_type"`
	StartTime      string       `q:"start_time"`
	Status         Status       `q:"status"`
	UserPercent    string       `q:"user_percent"`
	VaultId        string       `q:"vault_id"`
}

func (opts ListOpts) ToBackupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToBackupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return BackupPage{pagination.SinglePageBase(r)}
	})
}
