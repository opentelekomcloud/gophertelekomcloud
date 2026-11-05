package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/tasks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "54602a9d-5e22-4239-9123-77e350df4a34"
	taskID     = "8abfa7b272adc5b40172b73130065ae7"
	tasksPath  = "/v2/" + projectID + "/instances/" + instanceID + "/tasks"
)

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func handleGet(t *testing.T, path, expectedQuery, body string) {
	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertEquals(t, expectedQuery, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, body)
	})
}

const listResponse = `
{
  "task_count" : "2",
  "tasks" : [ {
    "id" : "8abfa7b372160bfd0172165864064079",
    "name" : "modifyAutoTopic",
    "user_name" : "paas_dms",
    "user_id" : "3df5acbc24a54fadb62a043c9000a307",
    "params" : "{\"old_auto_status\":true,\"new_auto_status\":false}",
    "status" : "EXECUTING",
    "created_at" : "2020-05-15T03:19:51.046Z",
    "updated_at" : "2020-05-15T03:19:51.065Z"
  }, {
    "id" : "8abfa7b372160bfd017216560af83e6e",
    "name" : "changeRetentionPolicy",
    "user_name" : "paas_dms",
    "user_id" : "3df5acbc24a54fadb62a043c9000a307",
    "params" : "{\"new_retention_policy\":\"produce_reject\",\"origin_retention_policy\":\"time_base\"}",
    "status" : "SUCCESS",
    "created_at" : "2020-05-15T03:17:17.176Z",
    "updated_at" : "2020-05-15T03:17:22.162Z"
  } ]
}`

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, tasksPath, "begin_time=20200515000000&end_time=20200516000000&limit=10&start=1", listResponse)

	res, err := tasks.List(dmsClient(), instanceID, tasks.ListOpts{
		Start:     1,
		Limit:     10,
		BeginTime: "20200515000000",
		EndTime:   "20200516000000",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "2", res.TaskCount)
	th.AssertEquals(t, 2, len(res.Tasks))
	th.AssertDeepEquals(t, tasks.Task{
		ID:        "8abfa7b372160bfd0172165864064079",
		Name:      "modifyAutoTopic",
		UserName:  "paas_dms",
		UserID:    "3df5acbc24a54fadb62a043c9000a307",
		Params:    `{"old_auto_status":true,"new_auto_status":false}`,
		Status:    "EXECUTING",
		CreatedAt: "2020-05-15T03:19:51.046Z",
		UpdatedAt: "2020-05-15T03:19:51.065Z",
	}, res.Tasks[0])
	th.AssertEquals(t, "SUCCESS", res.Tasks[1].Status)
}

func TestListOmitsEmptyQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, tasksPath, "", `{"task_count": "0", "tasks": []}`)

	res, err := tasks.List(dmsClient(), instanceID, tasks.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "0", res.TaskCount)
	th.AssertEquals(t, 0, len(res.Tasks))
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, tasksPath+"/"+taskID, "", `
{
  "task_count" : "1",
  "tasks" : [ {
    "id" : "8abfa7b272adc5b40172b73130065ae7",
    "name" : "bindInstancePublicIp",
    "user_name" : "paas_dms",
    "user_id" : "3df5acbc24a54fadb62a043c9000a307",
    "params" : "{\"enable_public_ip\":true}",
    "status" : "SUCCESS",
    "created_at" : "2020-06-15T08:55:53.606Z",
    "updated_at" : "2020-06-15T08:55:56.600Z"
  } ]
}`)

	res, err := tasks.Get(dmsClient(), instanceID, taskID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "1", res.TaskCount)
	th.AssertEquals(t, 1, len(res.Tasks))
	th.AssertEquals(t, taskID, res.Tasks[0].ID)
	th.AssertEquals(t, "bindInstancePublicIp", res.Tasks[0].Name)
	th.AssertEquals(t, `{"enable_public_ip":true}`, res.Tasks[0].Params)
	th.AssertEquals(t, "2020-06-15T08:55:56.600Z", res.Tasks[0].UpdatedAt)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(tasksPath+"/"+taskID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, tasks.Delete(dmsClient(), instanceID, taskID))
}

func TestDeleteRejectsUndocumentedStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(tasksPath+"/"+taskID, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	err := tasks.Delete(dmsClient(), instanceID, taskID)
	if err == nil {
		t.Fatal("expected error for status 200")
	}
}
