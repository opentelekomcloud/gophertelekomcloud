package backup

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	InstanceId string `q:"instance_id"`
	BackupId   string `q:"backup_id"`
	BackupType string `q:"backup_type"`
	Offset     string `q:"offset"`
	Limit      string `q:"limit"`
	BeginTime  string `q:"begin_time"`
	EndTime    string `q:"end_time"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]BackupListInfo, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("backups").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BackupPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractBackups(pages)
}

func ExtractBackups(r pagination.NewPage) ([]BackupListInfo, error) {
	var s ListResponse
	err := extract.Into(bytes.NewReader((r.(BackupPage)).Body), &s)
	return s.Backups, err
}

type BackupPage struct {
	pagination.NewSinglePageBase
}

type ListResponse struct {
	Backups    []BackupListInfo `json:"backups"`
	TotalCount int64            `json:"total_count"`
}

type BackupListInfo struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	BeginTime   string         `json:"begin_time"`
	EndTime     string         `json:"end_time"`
	Status      string         `json:"status"`
	TakeUpTime  int            `json:"take_up_time"`
	Type        string         `json:"type"`
	Size        float64        `json:"size"`
	Datastore   MysqlDatastore `json:"datastore"`
	InstanceId  string         `json:"instance_id"`
	Description string         `json:"description"`
}

type MysqlDatastore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}
