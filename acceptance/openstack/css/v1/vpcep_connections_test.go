package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/vpc_endpoint"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVpcepConnection(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID must be defined")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	enablePriv := true
	err = vpc_endpoint.Enable(client, clusterID, vpc_endpoint.EnableOpts{
		EndpointWithDnsName: &enablePriv,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, 1200))

	connections, err := vpc_endpoint.ListConnections(client, clusterID)
	th.AssertNoErr(t, err)

	if len(connections) == 0 {
		t.Error("no VPC endpoint connections found")
		return
	}

	if connections[0].ID == "" {
		t.Error("first VPC endpoint connection has empty ID")
		return
	}

	t.Logf("First VPC endpoint connection ID: %s\n", connections[0].ID)

	err = vpc_endpoint.Action(client, clusterID, vpc_endpoint.ActionOpts{
		Action:    "reject",
		Endpoints: []string{connections[0].ID},
	})
	th.AssertNoErr(t, err)

	err = vpc_endpoint.Action(client, clusterID, vpc_endpoint.ActionOpts{
		Action:    "receive",
		Endpoints: []string{connections[0].ID},
	})
	th.AssertNoErr(t, err)

	// UpdateWhitelist
	err = vpc_endpoint.UpdateWhitelist(client, clusterID, vpc_endpoint.UpdateWhitelistOpts{
		Permissions: []string{"682024355ec84ae89cf872c42c25de76"},
	})
	th.AssertNoErr(t, err)

	err = vpc_endpoint.Disable(client, clusterID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, 1200))
}
