package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/smart_connect"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "8959ab1c-7n1a-yyb1-a05t-93dfc361b32d"
	taskID     = "3c0ac4d1-5a1e-4c1b-8f6d-0123456789ab"

	base      = "/v2/" + projectID + "/instances/" + instanceID + "/connector"
	tasksPath = base + "/tasks"
	taskPath  = tasksPath + "/" + taskID
	// These two endpoints are documented with "kafka" right after the project ID.
	kafkaBase = "/v2/" + projectID + "/kafka/instances/" + instanceID
)

// dmsClient mirrors the real dmsv2 catalog endpoint: https://dms.{region}.otc.t-systems.com/v2/{project_id}/
func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

const obsTaskJSON = `{
	"task_name": "smart-connect-1",
	"topics_regex": "topic-obs*",
	"source_type": "NONE",
	"sink_type": "OBS_SINK",
	"sink_task": {
		"consumer_strategy": "earliest",
		"destination_file_type": "TEXT",
		"deliver_time_interval": 300,
		"obs_bucket_name": "obs_bucket",
		"obs_path": "obsTransfer-1810125534",
		"partition_format": "yyyy/MM/dd/HH/mm",
		"record_delimiter": "\\n",
		"store_keys": false,
		"obs_part_size": 5242880,
		"flush_size": 1000000,
		"timezone": "Asia/Chongqing",
		"schema_generator_class": "io.confluent.connect.storage.hive.schema.DefaultSchemaGenerator",
		"partitioner_class": "io.confluent.connect.storage.partitioner.TimeBasedPartitioner",
		"value_converter": "org.apache.kafka.connect.converters.ByteArrayConverter",
		"key_converter": "org.apache.kafka.connect.converters.ByteArrayConverter",
		"kv_delimiter": ":"
	},
	"id": "` + taskID + `",
	"status": "RUNNING",
	"create_time": 1708565483911
}`

func assertObsTask(t *testing.T, name, sinkType, status string, sink *smart_connect.SmartConnectTaskRespSinkConfig) {
	t.Helper()
	th.AssertEquals(t, "smart-connect-1", name)
	th.AssertEquals(t, "OBS_SINK", sinkType)
	th.AssertEquals(t, "RUNNING", status)
	th.AssertEquals(t, "obs_bucket", sink.ObsBucketName)
	th.AssertEquals(t, 300, sink.DeliverTimeInterval)
	// Fields returned by the API only.
	th.AssertEquals(t, 5242880, sink.ObsPartSize)
	th.AssertEquals(t, 1000000, sink.FlushSize)
	th.AssertEquals(t, "Asia/Chongqing", sink.Timezone)
	th.AssertEquals(t, "io.confluent.connect.storage.partitioner.TimeBasedPartitioner", sink.PartitionerClass)
	th.AssertEquals(t, ":", sink.KvDelimiter)
}

func TestEnable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(base, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{"specification": "100MB", "spec_code": "kafka.c3.mini.connector"}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"job_id": "7c3ec20c-11de-4df9-acc0-7ef1dea25dfe", "connector_id": "55b78880-9077-4c74-ad5a-6868555f76a4"}`)
	})

	res, err := smart_connect.Enable(dmsClient(), smart_connect.EnableOpts{
		InstanceId:    instanceID,
		Specification: "100MB",
		SpecCode:      "kafka.c3.mini.connector",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "7c3ec20c-11de-4df9-acc0-7ef1dea25dfe", res.JobId)
	th.AssertEquals(t, "55b78880-9077-4c74-ad5a-6868555f76a4", res.ConnectorId)
}

func TestDisable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(kafkaBase+"/delete-connector", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"job_id": "d366178c-29ea-4d5c-a344-fa2ece4a1836"}`)
	})

	res, err := smart_connect.Disable(dmsClient(), instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "d366178c-29ea-4d5c-a344-fa2ece4a1836", res.JobId)
}

func TestCreateTaskOBS(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(tasksPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{
			"task_name": "smart-connect-1",
			"start_later": false,
			"source_type": "NONE",
			"topics_regex": "topic-obs*",
			"sink_type": "OBS_SINK",
			"sink_task": {
				"consumer_strategy": "earliest",
				"destination_file_type": "TEXT",
				"deliver_time_interval": 300,
				"access_key": "ak",
				"secret_key": "sk",
				"obs_bucket_name": "obs_bucket",
				"obs_path": "obsTransfer-1810125534",
				"partition_format": "yyyy/MM/dd/HH/mm",
				"record_delimiter": "\\n",
				"store_keys": false
			}
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, obsTaskJSON)
	})

	res, err := smart_connect.CreateTask(dmsClient(), instanceID, smart_connect.CreateTaskOpts{
		TaskName:    "smart-connect-1",
		StartLater:  pointerto.Bool(false),
		SourceType:  "NONE",
		TopicsRegex: "topic-obs*",
		SinkType:    "OBS_SINK",
		SinkTask: &smart_connect.SmartConnectTaskSinkConfig{
			ConsumerStrategy:    "earliest",
			DestinationFileType: "TEXT",
			DeliverTimeInterval: 300,
			AccessKey:           "ak",
			SecretKey:           "sk",
			ObsBucketName:       "obs_bucket",
			ObsPath:             "obsTransfer-1810125534",
			PartitionFormat:     "yyyy/MM/dd/HH/mm",
			RecordDelimiter:     "\\n",
			StoreKeys:           pointerto.Bool(false),
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, taskID, res.ID)
	assertObsTask(t, res.TaskName, res.SinkType, res.Status, res.SinkTask)
}

func TestCreateTaskKafkaReplicator(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(tasksPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{
			"task_name": "smart-connect-2",
			"start_later": true,
			"source_type": "KAFKA_REPLICATOR_SOURCE",
			"source_task": {
				"current_cluster_name": "A",
				"cluster_name": "B",
				"user_name": "user1",
				"password": "pwd",
				"sasl_mechanism": "SCRAM-SHA-512",
				"instance_id": "b54c9dd8-0000-0000-0000-000000000000",
				"direction": "two-way",
				"sync_consumer_offsets_enabled": false,
				"replication_factor": 3,
				"task_num": 2,
				"rename_topic_enabled": false,
				"provenance_header_enabled": true,
				"consumer_strategy": "latest",
				"compression_type": "snappy",
				"topics_mapping": "topic-sc-1:topic-sc-3,topic-sc-2:topic-sc-4"
			}
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"task_name": "smart-connect-2", "source_type": "KAFKA_REPLICATOR_SOURCE", "sink_task": null,
			"sink_type": "NONE", "id": "`+taskID+`", "status": "RUNNING", "create_time": 1708427753133}`)
	})

	res, err := smart_connect.CreateTask(dmsClient(), instanceID, smart_connect.CreateTaskOpts{
		TaskName:   "smart-connect-2",
		StartLater: pointerto.Bool(true),
		SourceType: "KAFKA_REPLICATOR_SOURCE",
		SourceTask: &smart_connect.SmartConnectTaskSourceConfig{
			CurrentClusterName:         "A",
			ClusterName:                "B",
			UserName:                   "user1",
			Password:                   "pwd",
			SaslMechanism:              "SCRAM-SHA-512",
			InstanceId:                 "b54c9dd8-0000-0000-0000-000000000000",
			Direction:                  "two-way",
			SyncConsumerOffsetsEnabled: pointerto.Bool(false),
			ReplicationFactor:          3,
			TaskNum:                    2,
			RenameTopicEnabled:         pointerto.Bool(false),
			ProvenanceHeaderEnabled:    pointerto.Bool(true),
			ConsumerStrategy:           "latest",
			CompressionType:            "snappy",
			TopicsMapping:              "topic-sc-1:topic-sc-3,topic-sc-2:topic-sc-4",
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, taskID, res.ID)
	th.AssertEquals(t, "KAFKA_REPLICATOR_SOURCE", res.SourceType)
}

func TestListTasks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(tasksPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.CheckEquals(t, "2", r.URL.Query().Get("offset"))
		th.CheckEquals(t, "10", r.URL.Query().Get("limit"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, `{"tasks": [%s], "total_number": 1, "max_tasks": 18, "quota_tasks": 18}`, obsTaskJSON)
	})

	res, err := smart_connect.ListTasks(dmsClient(), smart_connect.QueryTasksOpts{
		InstanceId: instanceID,
		Offset:     pointerto.Int(2),
		Limit:      pointerto.Int(10),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, res.TotalNumber)
	th.AssertEquals(t, 18, res.MaxTasks)
	th.AssertEquals(t, 18, res.QuotaTasks)
	th.AssertEquals(t, 1, len(res.Tasks))
	assertObsTask(t, res.Tasks[0].TaskName, res.Tasks[0].SinkType, res.Tasks[0].Status, res.Tasks[0].SinkTask)
}

func TestGetTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(taskPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, obsTaskJSON)
	})

	res, err := smart_connect.GetTask(dmsClient(), instanceID, taskID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, taskID, res.ID)
	assertObsTask(t, res.TaskName, res.SinkType, res.Status, res.SinkTask)
}

func TestGetTaskRedisTarget(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(taskPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"task_name": "smart-connect-3", "topics": "topic-sc",
			"source_type": "REDIS_REPLICATOR_SOURCE",
			"source_task": {"redis_address": "192.168.91.179:6379", "redis_type": "standalone", "sync_mode": "RDB_ONLY",
				"full_sync_wait_ms": 13000, "full_sync_max_retry": 4, "ratelimit": -1},
			"sink_type": "REDIS_REPLICATOR_SINK",
			"sink_task": {"redis_address": "192.168.119.51:6379", "redis_type": "standalone", "target_db": -1},
			"id": "`+taskID+`", "status": "RUNNING", "create_time": 1708427753133}`)
	})

	res, err := smart_connect.GetTask(dmsClient(), instanceID, taskID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, -1, res.SourceTask.Ratelimit)
	// target_db is -1 here and must not be lost.
	th.AssertEquals(t, true, res.SinkTask.TargetDB != nil)
	th.AssertEquals(t, -1, *res.SinkTask.TargetDB)
}

func TestDeleteTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(taskPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, smart_connect.DeleteTask(dmsClient(), instanceID, taskID))
}

func TestPauseTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(taskPath+"/pause", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, smart_connect.PauseTask(dmsClient(), instanceID, taskID))
}

func TestRestartTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(taskPath+"/resume", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, smart_connect.RestartTask(dmsClient(), instanceID, taskID))
}

func TestStartOrRestartTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(kafkaBase+"/connector/tasks/"+taskID+"/restart", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, smart_connect.StartOrRestartTask(dmsClient(), instanceID, taskID))
}
