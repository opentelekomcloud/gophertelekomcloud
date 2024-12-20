package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestScaleInCluster(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	err = clusters.ScaleInCluster(client, clusterID, []clusters.ScaleInOpts{
		{
			Type:          "ess",
			ReduceNodeNum: 1,
		},
	})
	th.AssertNoErr(t, err)

	timeout := 1200

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
