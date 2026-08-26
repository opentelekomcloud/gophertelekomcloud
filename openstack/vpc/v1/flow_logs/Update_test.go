package flow_logs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/flow_logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs/flow-log-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{
			"flow_log": {
				"name": "flow-log-updated",
				"description": "updated",
				"admin_state": false
			}
		}`)
		_, _ = w.Write([]byte(`{"flow_log":` + flowLogJSON + `}`))
	})

	adminState := false
	actual, err := flow_logs.Update(serviceClient(), "flow-log-id", flow_logs.UpdateOpts{
		Name:        stringPointer("flow-log-updated"),
		Description: stringPointer("updated"),
		AdminState:  &adminState,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "flow-log-id", actual.ID)
}

func TestUpdateEmptyStrings(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs/flow-log-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"flow_log":{"name":"","description":""}}`)
		_, _ = w.Write([]byte(`{"flow_log":` + flowLogJSON + `}`))
	})

	actual, err := flow_logs.Update(serviceClient(), "flow-log-id", flow_logs.UpdateOpts{
		Name:        stringPointer(""),
		Description: stringPointer(""),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "flow-log-id", actual.ID)
}

func stringPointer(value string) *string {
	return &value
}
