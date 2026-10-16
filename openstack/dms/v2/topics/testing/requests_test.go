package testing

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
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

func handleUpdate(t *testing.T, expectedBody string) {
	th.Mux.HandleFunc(topicsPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, expectedBody)
		w.WriteHeader(http.StatusNoContent)
	})
}

func TestUpdateTopic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleUpdate(t, `{
		"topics": [{
			"id": "test01",
			"retention_time": 72,
			"sync_replication": false,
			"sync_message_flush": false,
			"new_partition_numbers": 6,
			"new_partition_brokers": [1, 2],
			"topic_other_configs": [
				{"name": "message.timestamp.type", "value": "LogAppendTime"},
				{"name": "max.message.bytes", "value": "10485760"}
			],
			"topic_desc": "This is a test topic"
		}]
	}`)

	err := topics.Update(dmsClient(), instanceID, topics.UpdateOpts{
		Topics: []topics.UpdateItem{{
			Name:                "test01",
			RetentionTime:       pointerto.Int(72),
			SyncReplication:     pointerto.Bool(false),
			SyncMessageFlush:    pointerto.Bool(false),
			NewPartitionNumbers: pointerto.Int(6),
			NewPartitionBrokers: []int{1, 2},
			TopicOtherConfigs: []topics.TopicOtherConfigs{
				{Name: "message.timestamp.type", Value: "LogAppendTime"},
				{Name: "max.message.bytes", Value: "10485760"},
			},
			Description: "This is a test topic",
		}},
	})
	th.AssertNoErr(t, err)
}

func TestUpdateTopicMinimal(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// Unset optional fields are omitted, so the server keeps current values.
	handleUpdate(t, `{"topics": [{"id": "test01"}]}`)

	err := topics.Update(dmsClient(), instanceID, topics.UpdateOpts{
		Topics: []topics.UpdateItem{{Name: "test01"}},
	})
	th.AssertNoErr(t, err)
}

func TestUpdateTopicWithoutTopics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// topics is optional according to the API reference.
	handleUpdate(t, `{}`)

	err := topics.Update(dmsClient(), instanceID, topics.UpdateOpts{})
	th.AssertNoErr(t, err)
}

func TestUpdateTopicValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	cases := map[string]struct {
		instanceID string
		opts       topics.UpdateOpts
		errContain string
	}{
		"missing instance id": {"", topics.UpdateOpts{Topics: []topics.UpdateItem{{Name: "test01"}}}, "instanceID"},
		"missing topic name":  {instanceID, topics.UpdateOpts{Topics: []topics.UpdateItem{{RetentionTime: pointerto.Int(1)}}}, "Name"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := topics.Update(dmsClient(), tc.instanceID, tc.opts)
			if err == nil || !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %v", tc.errContain, err)
			}
		})
	}
}

func TestUpdateTopicAPIError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(topicsPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error_code": "DMS.TEST", "error_msg": "mocked error"}`)
	})

	err := topics.Update(dmsClient(), instanceID, topics.UpdateOpts{
		Topics: []topics.UpdateItem{{Name: "test01"}},
	})
	var httpErr golangsdk.ErrDefault400
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault400, got %T: %v", err, err)
	}
	th.AssertEquals(t, true, strings.Contains(string(httpErr.Body), "DMS.TEST"))
}

func TestDeleteTopics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// Partial failure still returns 200 with per-topic results.
	th.Mux.HandleFunc(topicsPath+"/delete", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `{"topics": ["topic01", "topic02"]}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"topics": [{"id": "topic01", "success": true}, {"id": "topic02", "success": false}]}`)
	})

	res, err := topics.Delete(dmsClient(), instanceID, []string{"topic01", "topic02"})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &topics.DeleteResponse{Topics: []topics.TopicDelete{
		{Name: "topic01", Success: true},
		{Name: "topic02", Success: false},
	}}, res)
}

func TestDeleteTopicsValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	cases := map[string]struct {
		instanceID string
		topics     []string
		errContain string
	}{
		"missing instance id": {"", []string{"topic01"}, "instanceID"},
		"nil topics":          {instanceID, nil, "topics"},
		"empty topics":        {instanceID, []string{}, "topics"},
		"empty topic name":    {instanceID, []string{"topic01", ""}, "topics[1]"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := topics.Delete(dmsClient(), tc.instanceID, tc.topics)
			if err == nil || !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %v (response %+v)", tc.errContain, err, res)
			}
		})
	}
}

func TestDeleteTopicsAPIError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(topicsPath+"/delete", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error_code": "DMS.TEST", "error_msg": "mocked error"}`)
	})

	res, err := topics.Delete(dmsClient(), instanceID, []string{"topic01"})
	if res != nil {
		t.Fatalf("expected nil response on error, got %+v", res)
	}
	var httpErr golangsdk.ErrDefault400
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault400, got %T: %v", err, err)
	}
	th.AssertEquals(t, true, strings.Contains(string(httpErr.Body), "DMS.TEST"))
}

func TestGetTopic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/"+projectID+"/instances/"+instanceID+"/management/topics/test", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"topic": "test",
			"partitions": [{
				"partition": 0,
				"leader": 2,
				"replicas": [
					{"broker": 2, "leader": true, "in_sync": true, "size": 123971146, "lag": 0},
					{"broker": 1, "leader": false, "in_sync": false, "size": 123971146, "lag": 5}
				],
				"lso": 1,
				"leo": 13598,
				"hw": 13597,
				"last_update_timestamp": 1571477180985
			}],
			"group_subscribed": ["test-consumer-group"]
		}`)
	})

	res, err := topics.Get(dmsClient(), instanceID, "test")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &topics.TopicDetail{
		Name: "test",
		Partitions: []topics.Partition{{
			Partition: 0,
			Leader:    2,
			Replicas: []topics.Replica{
				{Broker: 2, Leader: true, InSync: true, Size: 123971146, Lag: 0},
				{Broker: 1, Leader: false, InSync: false, Size: 123971146, Lag: 5},
			},
			Lso:             1,
			Leo:             13598,
			Hw:              13597,
			UpdateTimestamp: 1571477180985,
		}},
		GroupSubscribed: []string{"test-consumer-group"},
	}, res)
}

func TestGetTopicValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	cases := map[string]struct {
		instanceID string
		topicName  string
		errContain string
	}{
		"missing instance id": {"", "test", "instanceID"},
		"missing topic name":  {instanceID, "", "topicName"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := topics.Get(dmsClient(), tc.instanceID, tc.topicName)
			if err == nil || !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %v (response %+v)", tc.errContain, err, res)
			}
		})
	}
}

func TestGetTopicAPIError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/"+projectID+"/instances/"+instanceID+"/management/topics/missing", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, `{"error_code": "DMS.TEST", "error_msg": "mocked error"}`)
	})

	res, err := topics.Get(dmsClient(), instanceID, "missing")
	if res != nil {
		t.Fatalf("expected nil response on error, got %+v", res)
	}
	var httpErr golangsdk.ErrDefault404
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault404, got %T: %v", err, err)
	}
}

const kafkaTopicPath = "/v2/" + projectID + "/kafka/instances/" + instanceID + "/topics/test"

func TestListPartitions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(kafkaTopicPath+"/partitions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		// An explicit zero offset must be sent.
		th.AssertDeepEquals(t, url.Values{"offset": {"0"}, "limit": {"10"}}, r.URL.Query())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"total": 3,
			"partitions": [{
				"partition": 0,
				"start_offset": 5,
				"last_offset": 1216303,
				"message_count": 1216298,
				"last_update_time": 1688011291458
			}]
		}`)
	})

	res, err := topics.ListPartitions(dmsClient(), instanceID, "test", topics.ListPartitionsOpts{
		Offset: pointerto.Int(0),
		Limit:  pointerto.Int(10),
	})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &topics.ListPartitionsResponse{
		Total: 3,
		Partitions: []topics.TopicPartition{{
			Partition:      0,
			StartOffset:    5,
			LastOffset:     1216303,
			MessageCount:   1216298,
			LastUpdateTime: 1688011291458,
		}},
	}, res)
}

func TestListPartitionsWithoutQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(kafkaTopicPath+"/partitions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.AssertDeepEquals(t, url.Values{}, r.URL.Query())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"total": 0, "partitions": []}`)
	})

	res, err := topics.ListPartitions(dmsClient(), instanceID, "test", topics.ListPartitionsOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(res.Partitions))
}

func TestListProducers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(kafkaTopicPath+"/producers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		// An explicit zero offset must be sent.
		th.AssertDeepEquals(t, url.Values{"offset": {"0"}, "limit": {"10"}}, r.URL.Query())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"total": 3,
			"producers": [{
				"producer_address": "192.0.0.149:40443",
				"broker_address": "192.0.0.146:9092",
				"join_time": 1687204743328
			}]
		}`)
	})

	res, err := topics.ListProducers(dmsClient(), instanceID, "test", topics.ListProducersOpts{
		Offset: pointerto.Int(0),
		Limit:  pointerto.Int(10),
	})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &topics.ListProducersResponse{
		Total: 3,
		Producers: []topics.Producer{{
			ProducerAddress: "192.0.0.149:40443",
			BrokerAddress:   "192.0.0.146:9092",
			JoinTime:        1687204743328,
		}},
	}, res)
}

func TestListProducersWithoutQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(kafkaTopicPath+"/producers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.AssertDeepEquals(t, url.Values{}, r.URL.Query())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"total": 0, "producers": []}`)
	})

	res, err := topics.ListProducers(dmsClient(), instanceID, "test", topics.ListProducersOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(res.Producers))
}

func TestListPartitionsAndProducersValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	cases := map[string]struct {
		instanceID string
		topicName  string
		errContain string
	}{
		"missing instance id": {"", "test", "instanceID"},
		"missing topic name":  {instanceID, "", "topicName"},
	}

	for name, tc := range cases {
		t.Run("partitions "+name, func(t *testing.T) {
			_, err := topics.ListPartitions(dmsClient(), tc.instanceID, tc.topicName, topics.ListPartitionsOpts{})
			if err == nil || !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %v", tc.errContain, err)
			}
		})
		t.Run("producers "+name, func(t *testing.T) {
			_, err := topics.ListProducers(dmsClient(), tc.instanceID, tc.topicName, topics.ListProducersOpts{})
			if err == nil || !strings.Contains(err.Error(), tc.errContain) {
				t.Fatalf("expected error containing %q, got %v", tc.errContain, err)
			}
		})
	}
}

func TestListPartitionsAndProducersAPIError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	for _, suffix := range []string{"/partitions", "/producers"} {
		th.Mux.HandleFunc(kafkaTopicPath+suffix, func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, `{"error_code": "DMS.TEST", "error_msg": "mocked error"}`)
		})
	}

	var httpErr golangsdk.ErrDefault400

	partitions, err := topics.ListPartitions(dmsClient(), instanceID, "test", topics.ListPartitionsOpts{})
	if partitions != nil || !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault400 and nil response, got %T: %v (response %+v)", err, err, partitions)
	}

	producers, err := topics.ListProducers(dmsClient(), instanceID, "test", topics.ListProducersOpts{})
	if producers != nil || !errors.As(err, &httpErr) {
		t.Fatalf("expected ErrDefault400 and nil response, got %T: %v (response %+v)", err, err, producers)
	}
}
