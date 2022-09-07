package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

type RestoreBackupOpts struct {
	Mappings []BackupRestoreServer `json:"mappings,omitempty"`
	PowerOn  bool                  `json:"power_on,omitempty"`
	ServerID string                `json:"server_id,omitempty"`
	VolumeID string                `json:"volume_id,omitempty"`
}

type BackupRestoreServer struct {
	BackupID string `json:"backup_id"`
	VolumeID string `json:"volume_id"`
}

func RestoreBackup(client *golangsdk.ServiceClient, backupID string, opts RestoreBackupOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "restore")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("backups", backupID, "restore"), b, nil, nil)
	return
}
