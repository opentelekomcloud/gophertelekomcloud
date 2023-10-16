package cluster

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v1.1/{project_id}/clusters/{cluster_id}
	type KeepBackup struct {
		KeepBackup int `json:"keep_last_manual_backup"`
	}
	_, err = client.DeleteWithBody(client.ServiceURL("clusters", id), KeepBackup{KeepBackup: 0}, nil)
	return
}
