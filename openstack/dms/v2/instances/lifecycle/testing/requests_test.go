package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const projectID = "0123456789abcdef0123456789abcdef"

// dmsClient mirrors the real dmsv2 catalog endpoint: https://dms.{region}.otc.t-systems.com/v2/{project_id}/
func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

const actionPath = "/v2/" + projectID + "/instances/action"

func TestBatchRestartDeleteInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(actionPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{
			"action": "restart",
			"instances": ["54602a9d-5e22-4239-9123-77e350df4a34", "7166cdea-dbad-4d79-9610-7163e6f8b640"]
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"results": [{"result": "success", "instance": "54602a9d-5e22-4239-9123-77e350df4a34"}]}`)
	})

	res, err := lifecycle.BatchRestartDelete(dmsClient(), lifecycle.BatchRestartDeleteOpts{
		Action:    "restart",
		Instances: []string{"54602a9d-5e22-4239-9123-77e350df4a34", "7166cdea-dbad-4d79-9610-7163e6f8b640"},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(res.Results))
	th.AssertEquals(t, "success", res.Results[0].Result)
	th.AssertEquals(t, "54602a9d-5e22-4239-9123-77e350df4a34", res.Results[0].Instance)
}

func TestBatchDeleteAllFailure(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(actionPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{"action": "delete", "all_failure": "kafka"}`)
		w.WriteHeader(http.StatusNoContent)
	})

	res, err := lifecycle.BatchRestartDelete(dmsClient(), lifecycle.BatchRestartDeleteOpts{
		Action:     "delete",
		AllFailure: "kafka",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(res.Results))
}

func TestUpdateInstanceConf(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	const instanceID = "8959ab1c-7n1a-yyb1-a05t-93dfc361b32d"
	th.Mux.HandleFunc("/v2/"+projectID+"/instances/"+instanceID+"/configs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestJSONRequest(t, r, `{
"kafka_configs": [
{"name": "connections.max.idle.ms", "value": "500000"},
{"name": "log.retention.hours", "value": "66"}
]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"job_id": "8abfa7b38ba79a20018ba9afc550576a", "dynamic_config": 0, "static_config": 2}`)
	})

	res, err := lifecycle.UpdateInstanceConf(dmsClient(), instanceID, lifecycle.UpdateInstanceConfOpts{
		KafkaConfigs: []lifecycle.KafkaConfig{
			{Name: "connections.max.idle.ms", Value: "500000"},
			{Name: "log.retention.hours", Value: "66"},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "8abfa7b38ba79a20018ba9afc550576a", res.JobID)
	th.AssertEquals(t, 0, res.DynamicConfig)
	th.AssertEquals(t, 2, res.StaticConfig)
}
