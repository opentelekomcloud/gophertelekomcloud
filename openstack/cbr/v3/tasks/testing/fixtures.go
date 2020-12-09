package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tasks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	expectedGetResponse = `
{
  "operation_log" : {
    "status" : "success",
    "provider_id" : "0daac4c5-6707-4851-97ba-169e36266b66",
    "checkpoint_id" : "b432511f-d889-428f-8b0e-5f47c524c6b6",
    "updated_at" : "2019-05-23T14:35:23.584418",
    "error_info" : {
      "message" : "",
      "code" : ""
    },
    "started_at" : "2019-05-23T14:31:36.007230",
    "id" : "4827f2da-b008-4507-ab7d-42d0df5ed912",
    "extra_info" : {
      "resource" : {
        "type" : "OS::Nova::Server",
        "id" : "1dab32fa-ebf2-415a-ab0b-eabe6353bc86",
        "name" : "ECS-0001"
      },
      "backup" : {
        "backup_name" : "manualbk_1234",
        "backup_id" : "0e5d0ef6-7f0a-4890-b98c-cb12490e31c1"
      },
      "common" : {
        "progress" : 100,
        "request_id" : "req-cdb98cc4-e87b-4f40-9b4a-57ec036620bc"
      }
    },
    "ended_at" : "2019-05-23T14:35:23.511155",
    "created_at" : "2019-05-23T14:31:36.039365",
    "operation_type" : "backup",
    "project_id" : "04f1829c788037ac2fb8c01eb2b04b95"
  }
}
`
)

var (
	expectedGetResponseData = &tasks.OperationLog{
		CheckpointID: "b432511f-d889-428f-8b0e-5f47c524c6b6",
		CreatedAt:    "2019-05-23T14:31:36.039365",
		EndedAt:      "2019-05-23T14:35:23.511155",
		ErrorInfo: tasks.OperationErrorInfo{
			Code:    "",
			Message: "",
		},
		ExtraInfo: tasks.OperationExtraInfo{
			Backup: tasks.OperationExtendInfoBackup{
				BackupID:   "0e5d0ef6-7f0a-4890-b98c-cb12490e31c1",
				BackupName: "manualbk_1234",
			},
			Common: tasks.OperationExtendInfoCommon{
				Progress:  100,
				RequestID: "req-cdb98cc4-e87b-4f40-9b4a-57ec036620bc",
			},
		},
		ID:            "4827f2da-b008-4507-ab7d-42d0df5ed912",
		OperationType: "backup",
		ProjectID:     "04f1829c788037ac2fb8c01eb2b04b95",
		ProviderID:    "0daac4c5-6707-4851-97ba-169e36266b66",
		StartedAt:     "2019-05-23T14:31:36.007230",
		Status:        "success",
		UpdatedAt:     "2019-05-23T14:35:23.584418",
	}
)

func handleTaskGet(t *testing.T) {
	th.Mux.HandleFunc("operation-logs/4827f2da-b008-4507-ab7d-42d0df5ed912", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, expectedGetResponse)
	})
}
