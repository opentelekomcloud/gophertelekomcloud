package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListCluster(t *testing.T) {
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	_, err = clusters.List(client, clusters.ListOpts{})
	th.AssertNoErr(t, err)
}

func TestCluster(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	eniSubnetID := clients.EnvOS.GetEnv("ENI_SUBNET_ID")
	eniCidr := clients.EnvOS.GetEnv("ENI_SUBNET_CIDR")
	if vpcID == "" || subnetID == "" || eniSubnetID == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID and OS_ENI_SUBNET_ID are required for this test")
	}
	if eniCidr == "" {
		eniCidr = "192.168.0.0/24"
	}

	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	cluster := cce.CreateTurboCluster(t, vpcID, subnetID, eniSubnetID, eniCidr)

	clusterID := cluster.Metadata.Id

	clusterGet, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, cluster.Metadata.Name, clusterGet.Metadata.Name)

	if clusterID != "" {
		cce.DeleteCluster(t, clusterID)
		clusterID = ""
	}
}
