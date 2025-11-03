package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetBackupPolicyOpts struct {
	InstanceId string `json:"-"`
	Type       string `q:"type"`
}

func GetBackupPolicy(client *golangsdk.ServiceClient, opts GetBackupPolicyOpts) (*GetBackupPolicyResponse, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("instances", opts.InstanceId, "backups", "policy")+query.String(), nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res GetBackupPolicyResponse
	return &res, extract.Into(raw.Body, &res)
}

type GetBackupPolicyResponse struct {
	BackupPolicy   *ShowBackupPolicyResult  `json:"backup_policy"`
	DatabaseTables []QueryDatabaseTableInfo `json:"database_tables,omitempty"`
}

type ShowBackupPolicyResult struct {
	KeepDays           int    `json:"keep_days"`
	DifferentialPeriod string `json:"differential_period"`
	IncrementalPeriod  string `json:"incremental_period"`
	StartTime          string `json:"start_time"`
	Period             string `json:"period"`
}

type QueryDatabaseTableInfo struct {
	DatabaseName string   `json:"database_name"`
	TableNames   []string `json:"table_names"`
}
