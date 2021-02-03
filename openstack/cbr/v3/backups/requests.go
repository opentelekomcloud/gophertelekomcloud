package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, id), &r.Body, nil)
	return
}

type ListOptsBuilder interface {
	ToBackupListQuery() (string, error)
}

type ListOpts struct {
	CheckpointID        string       `q:"checkpoint_id"`
	Dedicated           bool         `q:"dec"`
	EndTime             string       `q:"end_time"`
	EnterpriseProjectID string       `q:"enterprise_project_id"`
	ImageType           ImageType    `q:"image_type"`
	MemberStatus        string       `q:"member_status"`
	Name                string       `q:"name"`
	OwnType             string       `q:"own_type"`
	ParentID            string       `q:"parent_id"`
	ResourceAZ          string       `q:"resource_az"`
	ResourceID          string       `q:"resource_id"`
	ResourceName        string       `q:"resource_name"`
	ResorceType         string       `q:"resorce_type"`
	Sort                string       `q:"sort"`
	StartTime           string       `q:"start_time"`
	Status              BackupStatus `q:"status"`
	VaultID             string       `q:"vault_id"`
}

func (opts *ListOpts) ToBackupListQuery() (string, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return query.String(), nil
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
		return BackupPage{pagination.MarkerPageBase{PageResult: r}}
	})
}
