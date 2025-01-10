package event

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAlarmWhitelistOpts struct {
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 10,
	// and a value greater than 200 will be automatically converted to 200.
	Limit int `q:"limit"`
	// Enterprise project ID.
	// The value 0 indicates the default enterprise project.
	// To query all enterprise projects, set this parameter to all_granted_eps.
	EnterpriseProjectId string `q:"enterprise_project_id"`
	// Hash value of the event whitelist description (SHA256 algorithm)
	Hash string `q:"hash"`
	// Event type. Its value can be:
	// 1001: malware
	// 1010 : Rootkit
	// 1011: ransomware
	// 1015 : Web shell
	// 1017: reverse shell
	// 2001: Common vulnerability exploit
	// 2047: redis vulnerability exploit
	// 2048: Hadoop vulnerability exploit
	// 2049: MySQL vulnerability exploit
	// 3002: file privilege escalation
	// 3003: process privilege escalation
	// 3004: critical file change
	// 3005: file/directory change
	// 3007: abnormal process behavior
	// 3015: high-risk command execution
	// 3018: abnormal shell
	// 3027: suspicious crontab task
	// 4002: brute-force attack
	// 4004: abnormal login
	// 4006: Invalid system account
	EventType string `q:"event_type"`
}

func ListAlarmWhitelist(client *golangsdk.ServiceClient, opts ListAlarmWhitelistOpts) ([]WhitelistsResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("event", "white-list", "alarm").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v5/{project_id}/event/white-list/alarm
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return AlarmWhitelistPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAlarmWhitelists(pages)
}

type AlarmWhitelistPage struct {
	pagination.NewSinglePageBase
}

func ExtractAlarmWhitelists(r pagination.NewPage) ([]WhitelistsResp, error) {
	var s struct {
		Whitelists []WhitelistsResp `json:"data_list"`
	}
	err := extract.Into(bytes.NewReader((r.(AlarmWhitelistPage)).Body), &s)
	return s.Whitelists, err
}

type WhitelistsResp struct {
	// Enterprise project name
	EnterpriseProjectName string `json:"enterprise_project_name"`
	// Hash value of the event whitelist description (SHA256 algorithm)
	Hash string `json:"hash"`
	// Description
	Description string `json:"description"`
	// Intrusion type
	EventType int `json:"event_type"`
	// Time when the event whitelist is updated, in milliseconds.
	UpdatedAt int64 `json:"update_time"`
}
