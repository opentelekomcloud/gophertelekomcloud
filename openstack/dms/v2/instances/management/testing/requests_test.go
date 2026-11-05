package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "8959ab1c-7n1a-yyb1-a05t-93dfc361b32d"
)

const reassignPath = "/v2/" + projectID + "/kafka/instances/" + instanceID + "/reassign"

// dmsClient mirrors the real dmsv2 catalog endpoint: https://dms.{region}.otc.t-systems.com/v2/{project_id}/
func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func TestInitPartitionReassigningAutomatic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(reassignPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"reassignments": [{"topic": "topic-1", "brokers": [0, 1, 2], "replication_factor": 3}],
			"throttle": 10000000,
			"time_estimate": true
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"reassignment_time": 10}`)
	})

	opts := management.InitPartitionReassigningOpts{
		Reassignments: []management.PartitionReassign{{
			Topic:             "topic-1",
			Brokers:           []int{0, 1, 2},
			ReplicationFactor: 3,
		}},
		Throttle:     10000000,
		TimeEstimate: true,
	}
	res, err := management.InitPartitionReassigning(dmsClient(), instanceID, &opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", res.JobId)
	th.AssertEquals(t, 10, res.ReassignmentTime)
}

func TestInitPartitionReassigningManual(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(reassignPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		// Partition 0 must be sent; unset partition must be omitted.
		th.TestJSONRequest(t, r, `{
			"reassignments": [{
				"topic": "topic-1",
				"assignment": [
					{"partition": 0, "partition_brokers": [0, 1, 2]},
					{"partition_brokers": [1, 2, 0]}
				]
			}]
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"job_id": "8a2c259182ab0e9d0182ab1882560009"}`)
	})

	opts := management.InitPartitionReassigningOpts{
		Reassignments: []management.PartitionReassign{{
			Topic: "topic-1",
			Assignment: []*management.TopicAssignment{
				{Partition: pointerto.Int(0), PartitionBrokers: []int{0, 1, 2}},
				{PartitionBrokers: []int{1, 2, 0}},
			},
		}},
	}
	res, err := management.InitPartitionReassigning(dmsClient(), instanceID, &opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "8a2c259182ab0e9d0182ab1882560009", res.JobId)
	th.AssertEquals(t, 0, res.ReassignmentTime)
}

func TestReassignReplicas(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	const topic = "topic-1"
	th.Mux.HandleFunc("/v2/"+projectID+"/instances/"+instanceID+"/management/topics/"+topic+"/replicas-reassignment",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestJSONRequest(t, r, `{
				"partitions": [
					{"partition": 1, "replicas": [1, 2]},
					{"partition": 0, "replicas": [0, 1]}
				]
			}`)
			w.WriteHeader(http.StatusNoContent)
		})

	opts := management.ReassignReplicasOpts{
		Partitions: []*management.Partition{
			{PartitionID: pointerto.Int(1), Replicas: []int{1, 2}},
			{PartitionID: pointerto.Int(0), Replicas: []int{0, 1}},
		},
	}
	th.AssertNoErr(t, management.ReassignReplicas(dmsClient(), instanceID, topic, opts))
}

func TestGetMonitoringDimensions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/"+projectID+"/instances/"+instanceID+"/ces-hierarchy",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `
{
  "dimensions" : [ {
    "name" : "kafka_instance_id",
    "metrics" : [ "current_partitions", "current_topics" ],
    "key_name" : [ "instance_ids" ],
    "dim_router" : [ "kafka_instance_id" ]
  }, {
    "name" : "kafka_topics",
    "metrics" : [ "topic_data_size" ],
    "key_name" : [ "queues" ],
    "dim_router" : [ "kafka_instance_id", "kafka_topics" ],
    "children" : [ {
      "name" : "kafka_partitions",
      "metrics" : [ "produced_messages", "partition_messages" ],
      "key_name" : [ "queues", "partitions" ],
      "dim_router" : [ "kafka_instance_id", "kafka_topics", "kafka_partitions" ]
    } ]
  } ],
  "instance_ids" : [ { "name" : "68f3f6a0-3741-453b-bda9-a6ff6b5bb6f7" } ],
  "nodes" : [ { "name" : "broker-0" }, { "name" : "broker-1" } ],
  "queues" : [ {
    "name" : "mytest",
    "partitions" : [ { "name" : "0" }, { "name" : "1" } ]
  } ],
  "groups" : [ {
    "name" : "test-consumer-group",
    "queues" : [ {
      "name" : "mytest",
      "partitions" : [ { "name" : "0" } ]
    } ]
  } ]
}`)
		})

	res, err := management.GetMonitoringDimensions(dmsClient(), instanceID)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &management.GetMonitoringDimensionsResp{
		Dimensions: []management.Dimension{
			{
				Name:      "kafka_instance_id",
				Metrics:   []string{"current_partitions", "current_topics"},
				KeyName:   []string{"instance_ids"},
				DimRouter: []string{"kafka_instance_id"},
			},
			{
				Name:      "kafka_topics",
				Metrics:   []string{"topic_data_size"},
				KeyName:   []string{"queues"},
				DimRouter: []string{"kafka_instance_id", "kafka_topics"},
				Children: []management.DimensionChild{
					{
						Name:      "kafka_partitions",
						Metrics:   []string{"produced_messages", "partition_messages"},
						KeyName:   []string{"queues", "partitions"},
						DimRouter: []string{"kafka_instance_id", "kafka_topics", "kafka_partitions"},
					},
				},
			},
		},
		InstanceIDs: []management.DimensionObject{{Name: "68f3f6a0-3741-453b-bda9-a6ff6b5bb6f7"}},
		Nodes:       []management.DimensionObject{{Name: "broker-0"}, {Name: "broker-1"}},
		Queues: []management.DimensionQueue{
			{Name: "mytest", Partitions: []management.DimensionObject{{Name: "0"}, {Name: "1"}}},
		},
		Groups: []management.DimensionGroup{
			{
				Name: "test-consumer-group",
				Queues: []management.DimensionQueue{
					{Name: "mytest", Partitions: []management.DimensionObject{{Name: "0"}}},
				},
			},
		},
	}, res)
}
