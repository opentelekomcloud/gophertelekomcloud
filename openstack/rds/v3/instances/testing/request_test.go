package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestUpgradeDescription(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	instanceID := "5b409baece064984a1b3eef6addae50cin01"
	expectedAlias := "alias-test"

	th.Mux.HandleFunc("/instances/"+instanceID+"/alias", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")

		th.TestJSONRequest(t, r, `
{
	"alias": "alias-test"
}
		`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	_, err := instances.UpgradeDescription(client.ServiceClient(), instances.UpgradeDescriptionOpts{
		InstanceId: instanceID,
		Alias:      &expectedAlias,
	})

	th.AssertNoErr(t, err)
}
