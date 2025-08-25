package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPolicy(client *golangsdk.ServiceClient, instanceId string) (*BackupPolicy, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "backups", "policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res BackupPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "backup_policy")
}

type BackupPolicy struct {
	KeepDays  int    `json:"keep_days"`
	StartTime string `json:"start_time"`
	Period    string `json:"period"`
}
