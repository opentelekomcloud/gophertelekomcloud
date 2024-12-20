package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateSecurityGroup(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	sgID := openstack.DefaultSecurityGroup(t)
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	err = clusters.UpdateSecurityGroup(client, clusterID, clusters.SecurityGroupOpts{
		SecurityGroupID: sgID,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
