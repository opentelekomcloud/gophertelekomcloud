package testing

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "54602a9d-5e22-4239-9123-77e350df4a34"
	topicsPath = "/v2/" + projectID + "/instances/" + instanceID + "/topics"
)

// dmsClient mirrors the real dmsv2 catalog endpoint: https://dms.{region}.otc.t-systems.com/v2/{project_id}/
func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func handleCreate(t *testing.T, expectedBody string) {
	th.Mux.HandleFunc(topicsPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, expectedBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"name": "test01"}`)
	})
}

func TestCreateTopic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleCreate(t, `{
		"id": "test01",
		"partition": 3,
		"replication": 3,
		"retention_time": 72,
		"sync_message_flush": true,
		"sync_replication": true,
		"topic_other_configs": [
			{"name": "message.timestamp.type", "value": "LogAppendTime"},
			{"name": "max.message.bytes", "value": "10485760"}
		],
		"topic_desc": "test topic"
	}`)

	res, err := topics.Create(dmsClient(), instanceID, topics.CreateOpts{
		Name:             "test01",
		Partition:        3,
		Replication:      3,
		RetentionTime:    72,
		SyncMessageFlush: true,
		SyncReplication:  true,
		TopicOtherConfigs: []topics.TopicOtherConfigs{
			{Name: "message.timestamp.type", Value: "LogAppendTime"},
			{Name: "max.message.bytes", Value: "10485760"},
		},
		Description: "test topic",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test01", res.Name)
}

func TestCreateTopicMinimal(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// Optional fields left at zero values are omitted so that server defaults apply.
	handleCreate(t, `{"id": "test01"}`)

	res, err := topics.Create(dmsClient(), instanceID, topics.CreateOpts{Name: "test01"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test01", res.Name)
}

func TestCreateTopicOtherConfigsOmitEmpty(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// name and value are optional in topic_other_configs; empty strings are omitted.
	handleCreate(t, `{
		"id": "test01",
		"topic_other_configs": [{"name": "message.timestamp.type"}, {"value": "LogAppendTime"}]
	}`)

	_, err := topics.Create(dmsClient(), instanceID, topics.CreateOpts{
		Name: "test01",
		TopicOtherConfigs: []topics.TopicOtherConfigs{
			{Name: "message.timestamp.type"},
			{Value: "LogAppendTime"},
		},
	})
	th.AssertNoErr(t, err)
}

func TestCreateTopicValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	cases := map[string]struct {
		instanceID string
		opts       topics.CreateOpts
		errContain string
	}{
		"missing instance id": {"", topics.CreateOpts{Name: "test01"}, "instanceID"},
		"missing topic name":  {instanceID, topics.CreateOpts{Partition: 3}, "Name"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := topics.Create(dmsClient(), tc.instanceID, tc.opts)
			if err == nil {
				t.Fatalf("expected error, got response %+v", res)
			}
			if !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %q", tc.errContain, err)
			}
		})
	}
}

func TestCreateTopicAPIError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(topicsPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error_code": "DMS.TEST", "error_msg": "mocked error"}`)
	})

	res, err := topics.Create(dmsClient(), instanceID, topics.CreateOpts{Name: "test01"})
	if res != nil {
		t.Fatalf("expected nil response on error, got %+v", res)
	}
	var httpErr golangsdk.ErrDefault400
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault400, got %T: %v", err, err)
	}
	th.AssertEquals(t, true, strings.Contains(string(httpErr.Body), "DMS.TEST"))
}

const listResponse = `{
	"total": 1,
	"size": 0,
	"topics": [{
		"policiesOnly": false,
		"name": "Topic-test01",
		"replication": 3,
		"partition": 3,
		"retention_time": 72,
		"sync_replication": true,
		"sync_message_flush": false,
		"topic_other_configs": [{
			"name": "max.message.bytes",
			"valid_values": "[0...10485760]",
			"default_value": "10485760",
			"config_type": "dynamic",
			"value": "10485760",
			"value_type": "int"
		}],
		"external_configs": {"key": "value"},
		"topic_type": 0,
		"topic_desc": "This is a test topic",
		"created_at": 1688112779916
	}],
	"remain_partitions": 294,
	"max_partitions": 300,
	"topic_max_partitions": 200
}`

func handleList(t *testing.T, expectedQuery url.Values) {
	th.Mux.HandleFunc(topicsPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertDeepEquals(t, expectedQuery, r.URL.Query())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponse)
	})
}

func TestListTopics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// An explicit "0" offset must be sent, as in the API reference example.
	handleList(t, url.Values{"offset": {"0"}, "limit": {"10"}})

	res, err := topics.List(dmsClient(), instanceID, topics.ListOpts{Offset: "0", Limit: "10"})
	th.AssertNoErr(t, err)

	expected := &topics.ListResponse{
		Total:              1,
		Size:               0,
		RemainPartitions:   294,
		MaxPartitions:      300,
		TopicMaxPartitions: 200,
		Topics: []topics.Topic{{
			PoliciesOnly:     false,
			Name:             "Topic-test01",
			Replication:      3,
			Partition:        3,
			RetentionTime:    72,
			SyncReplication:  true,
			SyncMessageFlush: false,
			ExternalConfigs:  map[string]interface{}{"key": "value"},
			TopicType:        0,
			TopicOtherConfigs: []topics.TopicOtherConfigsResp{{
				Name:         "max.message.bytes",
				ValidValues:  "[0...10485760]",
				DefaultValue: "10485760",
				ConfigType:   "dynamic",
				Value:        "10485760",
				ValueType:    "int",
			}},
			Description: "This is a test topic",
			CreatedAt:   1688112779916,
		}},
	}
	th.AssertDeepEquals(t, expected, res)
}

func TestListTopicsWithoutQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleList(t, url.Values{})

	res, err := topics.List(dmsClient(), instanceID, topics.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(res.Topics))
}

func TestListTopicsMissingInstanceID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := topics.List(dmsClient(), "", topics.ListOpts{})
	if err == nil || !strings.Contains(err.Error(), "instanceID") {
		t.Fatalf("expected missing instanceID error, got %v", err)
	}
}
