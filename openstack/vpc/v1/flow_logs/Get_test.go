package flow_logs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/flow_logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/fl/flow_logs/flow-log-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"flow_log":` + flowLogJSON + `}`))
	})

	actual, err := flow_logs.Get(serviceClient(), "flow-log-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "flow-log-id", actual.ID)
	th.AssertEquals(t, "project-id", actual.TenantID)
}
