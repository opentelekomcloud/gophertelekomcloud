package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestChangePasswordWorkflow(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	opts := clusters.ChangePasswordOpts{

		NewPassword: tools.RandomString("newpass-css-", 4),
	}
	err = clusters.ChangePassword(client, clusterID, opts)

	th.AssertNoErr(t, err)
}
