package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/stretchr/testify/require"
)

func TestKubeConfig(t *testing.T) {
	routerID := clients.EnvOS.GetEnv("VPC_ID", "ROUTER_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if routerID == "" || subnetID == "" {
		t.Skip("OS_ROUTER_ID and OS_NETWORK_ID are required for this test")
	}
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	clusterID := cce.CreateCluster(t, routerID, subnetID)
	t.Cleanup(func() {
		cce.DeleteCluster(t, clusterID)
	})

	kubeConfig, err := clusters.GetCert(client, clusterID)
	th.AssertNoErr(t, err)
	require.NotEmpty(t, kubeConfig)

	kubeConfigExp, err := clusters.GetCertWithExpiration(client, clusterID, clusters.ExpirationOpts{
		Duration: 5,
	})
	th.AssertNoErr(t, err)
	require.NotEmpty(t, kubeConfigExp)
}
