package flow_logs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/flow_logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{
			"flow_log": {
				"name": "flow-log",
				"description": "description",
				"resource_type": "vpc",
				"resource_id": "vpc-id",
				"traffic_type": "all",
				"log_group_id": "group-id",
				"log_topic_id": "topic-id",
				"index_enabled": true
			}
		}`)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"flow_log":` + flowLogJSON + `}`))
	})

	indexEnabled := true
	actual, err := flow_logs.Create(serviceClient(), flow_logs.CreateOpts{
		Name:         "flow-log",
		Description:  "description",
		ResourceType: "vpc",
		ResourceID:   "vpc-id",
		TrafficType:  "all",
		LogGroupID:   "group-id",
		LogTopicID:   "topic-id",
		IndexEnabled: &indexEnabled,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "flow-log-id", actual.ID)
	th.AssertEquals(t, true, actual.IndexEnabled)
}

func TestCreateMissingRequiredField(t *testing.T) {
	actual, err := flow_logs.Create(serviceClient(), flow_logs.CreateOpts{})
	if err == nil {
		t.Fatal("expected required field validation error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
