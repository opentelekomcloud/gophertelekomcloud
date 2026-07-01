package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/projects"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"enterprise_projects":[{"id":"0","name":"default","description":"Default project","status":1,"created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-02T00:00:00Z","type":"prod","delete_flag":false}],"total_count":1}`)
	})

	actual, err := projects.List(client.ServiceClient(), projects.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, actual.TotalCount)
	th.AssertEquals(t, 1, len(actual.EnterpriseProjects))
	th.AssertEquals(t, "0", actual.EnterpriseProjects[0].ID)
	th.AssertEquals(t, "default", actual.EnterpriseProjects[0].Name)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"enterprise_project":{"id":"0","name":"default","description":"Default project","status":1,"created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-02T00:00:00Z","type":"prod","delete_flag":false}}`)
	})

	actual, err := projects.Get(client.ServiceClient(), "0")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "0", actual.ID)
	th.AssertEquals(t, "default", actual.Name)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"name":"project-a","description":"created by test"}`)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, `{"enterprise_project":{"id":"0","name":"project-a","description":"created by test","status":1,"type":"prod","delete_flag":false}}`)
	})

	actual, err := projects.Create(client.ServiceClient(), projects.CreateOpts{
		Name:        "project-a",
		Description: "created by test",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "project-a", actual.Name)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"name":"default","description":"updated by test"}`)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"enterprise_project":{"id":"0","name":"default","description":"updated by test","status":1,"type":"prod","delete_flag":false}}`)
	})

	actual, err := projects.Update(client.ServiceClient(), "0", projects.UpdateOpts{
		Name:        "default",
		Description: "updated by test",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated by test", actual.Description)
}

func TestAction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/0/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"action":"enable"}`)
		w.WriteHeader(http.StatusNoContent)
	})

	err := projects.Action(client.ServiceClient(), "0", projects.ActionOpts{Action: "enable"})
	th.AssertNoErr(t, err)
}
