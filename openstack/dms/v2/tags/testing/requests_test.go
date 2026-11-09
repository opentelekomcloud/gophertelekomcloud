package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID   = "0123456789abcdef0123456789abcdef"
	instanceID  = "54602a9d-5e22-4239-9123-77e350df4a34"
	projectPath = "/v2/" + projectID + "/kafka/tags"
	tagsPath    = "/v2/" + projectID + "/kafka/" + instanceID + "/tags"
	actionPath  = tagsPath + "/action"
)

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func handleAction(t *testing.T, expectedBody string) {
	th.Mux.HandleFunc(actionPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, expectedBody)
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleGet(t *testing.T, path, body string) {
	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, body)
	})
}

func TestActionCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleAction(t, `{"action": "create", "tags": [{"key": "key1", "value": "value1"}, {"key": "key2", "value": "value2"}]}`)

	err := tags.Action(dmsClient(), instanceID, tags.ActionOpts{
		Action: "create",
		Tags: []tags.TagEntity{
			{Key: "key1", Value: pointerto.String("value1")},
			{Key: "key2", Value: pointerto.String("value2")},
		},
	})
	th.AssertNoErr(t, err)
}

func TestActionEmptyValueVsOmittedValue(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleAction(t, `{"action": "delete", "tags": [{"key": "key1", "value": ""}, {"key": "key2"}]}`)

	err := tags.Action(dmsClient(), instanceID, tags.ActionOpts{
		Action: "delete",
		Tags: []tags.TagEntity{
			{Key: "key1", Value: pointerto.String("")},
			{Key: "key2"},
		},
	})
	th.AssertNoErr(t, err)
}

func TestActionOmitsEmptyFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleAction(t, `{}`)

	th.AssertNoErr(t, tags.Action(dmsClient(), instanceID, tags.ActionOpts{}))
}

func TestActionRejectsUndocumentedStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(actionPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	err := tags.Action(dmsClient(), instanceID, tags.ActionOpts{Action: "create"})
	if err == nil {
		t.Fatal("expected error for status 200")
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, tagsPath, `
{
  "tags" : [ {
    "key" : "key1",
    "value" : "value1"
  }, {
    "key" : "key2",
    "value" : ""
  } ]
}`)

	res, err := tags.Get(dmsClient(), instanceID)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []tags.Tag{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: ""},
	}, res.Tags)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, projectPath, `
{
  "tags" : [ {
    "key" : "key1",
    "values" : [ "value-test", "value1" ]
  }, {
    "key" : "key2",
    "values" : [ "value2" ]
  } ]
}`)

	res, err := tags.List(dmsClient())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []tags.ProjectTag{
		{Key: "key1", Values: []string{"value-test", "value1"}},
		{Key: "key2", Values: []string{"value2"}},
	}, res.Tags)
}
