package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
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
	if vpcID == "" {
		t.Skip("OS_VPC_ID is required for this test")
	}

	clientNet, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	listOpts := subnets.ListOpts{
		VpcID: vpcID,
	}
	subnetsList, err := subnets.List(clientNet, listOpts)
	th.AssertNoErr(t, err)

	if len(subnetsList) < 1 {
		t.Skip("no subnets found in selected VPC")
	}

	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	cluster := cce.CreateTurboCluster(t, vpcID, subnetsList[0].NetworkID, subnetsList[0].SubnetID, subnetsList[0].CIDR)

	clusterID := cluster.Metadata.Id

	clusterGet, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, cluster.Metadata.Name, clusterGet.Metadata.Name)

	if clusterID != "" {
		cce.DeleteCluster(t, clusterID)
	}
}
