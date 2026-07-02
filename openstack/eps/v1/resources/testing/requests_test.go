package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/resources"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestFilter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/project-id-1/resources/filter", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"projects":["project-id-1"],"resource_types":["ecs"],"limit":10,"matches":[{"key":"resource_name","value":"my-server"}]}`)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"resources":[{"project_id":"project-id-1","project_name":"default","resource_id":"res-1","resource_name":"my-server","resource_type":"ecs"}],"errors":[],"total_count":1}`)
	})

	result, err := resources.Filter(client.ServiceClient(), "project-id-1", resources.FilterOpts{
		Projects:      []string{"project-id-1"},
		ResourceTypes: []string{"ecs"},
		Limit:         10,
		Matches:       []resources.Match{{Key: "resource_name", Value: "my-server"}},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, result.TotalCount)
	th.AssertEquals(t, 1, len(result.Resources))
	th.AssertEquals(t, "res-1", result.Resources[0].ResourceID)
	th.AssertEquals(t, "ecs", result.Resources[0].ResourceType)
}

func TestMigrate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/enterprise-projects/source-project/resources-migrate", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"project_id":"target-project","resources":[{"resource_id":"res-1","resource_type":"ecs"}]}`)
		w.WriteHeader(http.StatusNoContent)
	})

	err := resources.Migrate(client.ServiceClient(), "source-project", resources.MigrateOpts{
		ProjectID: "target-project",
		Resources: []resources.MigrateResource{
			{ResourceID: "res-1", ResourceType: "ecs"},
		},
	})
	th.AssertNoErr(t, err)
}
