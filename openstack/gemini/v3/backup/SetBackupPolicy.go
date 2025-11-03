package backup

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type SetBackupPolicyOpts struct {
	InstanceId     string                   `json:"-"`
	BackupPolicy   BackupPolicy             `json:"backup_policy"`
	DatabaseTables []PutDatabaseTablePolicy `json:"database_tables,omitempty"`
}

func SetBackupPolicy(client *golangsdk.ServiceClient, opts SetBackupPolicyOpts) error {
	_, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "backups", "policy"), opts, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}

type BackupPolicy struct {
	KeepDays  int    `json:"keep_days"`
	StartTime string `json:"start_time,omitempty"`
	Period    string `json:"period,omitempty"`
}

type PutDatabaseTablePolicy struct {
	DatabaseName string   `json:"database_name"`
	TableNames   []string `json:"table_names,omitempty"`
}
