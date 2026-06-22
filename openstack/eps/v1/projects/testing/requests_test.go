package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/projects"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const epID = "0"

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"enterprise_projects":[{"id":"0","name":"default","description":"Default project","status":1,"created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-02T00:00:00Z","type":"prod","delete_flag":false}]}`)
	})

	pages := 0
	err := projects.List(client.ServiceClient(), projects.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := projects.ExtractProjects(page)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, 1, len(actual))
		th.AssertEquals(t, epID, actual[0].ID)
		th.AssertEquals(t, "default", actual[0].Name)
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, pages)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/"+epID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"enterprise_project":{"id":"0","name":"default","description":"Default project","status":1,"created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-02T00:00:00Z","type":"prod","delete_flag":false}}`)
	})

	actual, err := projects.Get(client.ServiceClient(), epID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, epID, actual.ID)
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
	}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "project-a", actual.Name)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/"+epID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"name":"default","description":"updated by test"}`)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"enterprise_project":{"id":"0","name":"default","description":"updated by test","status":1,"type":"prod","delete_flag":false}}`)
	})

	actual, err := projects.Update(client.ServiceClient(), epID, projects.UpdateOpts{Name: "default", Description: "updated by test"}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated by test", actual.Description)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/"+epID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	err := projects.Delete(client.ServiceClient(), epID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestAction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/"+epID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"action":"enable"}`)
		w.WriteHeader(http.StatusNoContent)
	})

	err := projects.Action(client.ServiceClient(), epID, projects.ActionOpts{Action: "enable"}).ExtractErr()
	th.AssertNoErr(t, err)
}
