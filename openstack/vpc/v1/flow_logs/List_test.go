package flow_logs_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/flow_logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		expected := map[string]string{
			"id":            "flow-log-id",
			"name":          "flow-log",
			"tenant_id":     "project-id",
			"description":   "description",
			"resource_type": "vpc",
			"resource_id":   "vpc-id",
			"traffic_type":  "all",
			"log_group_id":  "group-id",
			"log_topic_id":  "topic-id",
			"status":        "ACTIVE",
			"limit":         "2",
		}
		for key, value := range expected {
			if actual := r.URL.Query().Get(key); actual != value {
				t.Fatalf("unexpected %s: %q", key, actual)
			}
		}

		switch r.URL.Query().Get("marker") {
		case "":
			_, _ = fmt.Fprintf(
				w,
				`{"flow_logs":[%s,%s]}`,
				strings.ReplaceAll(flowLogJSON, "flow-log-id", "flow-log-1"),
				strings.ReplaceAll(flowLogJSON, "flow-log-id", "flow-log-2"),
			)
		case "flow-log-2":
			_, _ = fmt.Fprintf(
				w,
				`{"flow_logs":[%s]}`,
				strings.ReplaceAll(flowLogJSON, "flow-log-id", "flow-log-3"),
			)
		case "flow-log-3":
			_, _ = w.Write([]byte(`{"flow_logs":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	limit := 2
	actual, err := flow_logs.List(serviceClient(), flow_logs.ListOpts{
		ID:           "flow-log-id",
		Name:         "flow-log",
		TenantID:     "project-id",
		Description:  "description",
		ResourceType: "vpc",
		ResourceID:   "vpc-id",
		TrafficType:  "all",
		LogGroupID:   "group-id",
		LogTopicID:   "topic-id",
		Status:       "ACTIVE",
		Limit:        &limit,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "flow-log-1", actual[0].ID)
	th.AssertEquals(t, "flow-log-3", actual[2].ID)
	th.AssertEquals(t, true, actual[0].AdminState)
}

func TestListZeroLimit(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs", func(w http.ResponseWriter, r *http.Request) {
		if actual := r.URL.Query().Get("limit"); actual != "0" {
			t.Fatalf("unexpected limit %q", actual)
		}
		_, _ = w.Write([]byte(`{"flow_logs":[]}`))
	})

	limit := 0
	actual, err := flow_logs.List(serviceClient(), flow_logs.ListOpts{Limit: &limit})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"flow_logs":`))
	})

	actual, err := flow_logs.List(serviceClient(), flow_logs.ListOpts{})
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
