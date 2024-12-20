package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestChangeClusterNameWorkflow(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	opts := clusters.ChangeClusterNameOpts{
		DisplayName: tools.RandomString("changed-css-", 4),
	}
	err = clusters.ChangeClusterName(client, clusterID, opts)
	th.AssertNoErr(t, err)
}
