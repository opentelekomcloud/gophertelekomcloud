package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	backupID = "6df2b54c-dd62-4059-a07c-1b8f24f2725d"
)

var (
	getBackupResponse = fmt.Sprintf(`
{
  "backup" : {
    "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
    "checkpoint_id" : "8b0851a8-adf3-4f4c-a914-dead08bf9664",
    "enterprise_project_id" : "0",
    "vault_id" : "3b5816b5-f29c-4172-9d9a-76c719a659ce",
    "id" : "%s",
    "resource_az" : "az1.dc1",
    "image_type" : "backup",
    "resource_id" : "94eba8b2-acc9-4d82-badc-127144cc5526",
    "resource_size" : 40,
    "children" : [ {
      "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
      "checkpoint_id" : "8b0851a8-adf3-4f4c-a914-dead08bf9664",
      "vault_id" : "3b5816b5-f29c-4172-9d9a-76c719a659ce",
      "id" : "5d822633-2bbf-4af8-a16e-5ab1c7705235",
      "image_type" : "backup",
      "resource_id" : "eccbcfdd-f843-4bbb-b2c0-a5ce861f9376",
      "resource_size" : 40,
      "children" : [ ],
      "parent_id" : "6df2b54c-dd62-4059-a07c-1b8f24f2725d",
      "extend_info" : {
        "auto_trigger" : true,
        "snapshot_id" : "5230a977-1a94-4092-8edd-519303a44cda",
        "bootable" : true,
        "encrypted" : true
      },
      "project_id" : "4229d7a45436489f8c3dc2b1d35d4987",
      "status" : "available",
      "resource_name" : "ecs-1f0f-0002",
      "replication_records" : [ ],
      "name" : "autobk_a843_ecs-1f0f-0002",
      "created_at" : "2019-05-10T07:59:59.450700",
      "resource_type" : "OS::Cinder::Volume"
    } ],
    "extend_info" : {
      "auto_trigger" : true,
      "supported_restore_mode" : "backup",
      "contain_system_disk" : true,
      "support_lld" : true,
      "app_consistency" : {
        "app_consistency_error_code" : "0",
        "app_consistency_status" : "0",
        "app_consistency_error_message" : "",
        "app_consistency" : "0"
      }
    },
    "project_id" : "4229d7a45436489f8c3dc2b1d35d4987",
    "status" : "available",
    "resource_name" : "ecs-1f0f-0002",
    "description" : "backup_description",
    "replication_records" : [ ],
    "name" : "backup_name",
    "created_at" : "2019-05-10T07:59:12.084695",
    "resource_type" : "OS::Nova::Server"
  }
}
`, backupID)
)

func handleBackup(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/backups/%s", backupID), func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, getBackupResponse)
		}
	})
}
