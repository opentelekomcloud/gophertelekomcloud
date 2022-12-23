package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDWS(t *testing.T) {
	client, err := clients.NewDWSV1Client()
	th.AssertNoErr(t, err)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID is missing but DWS test requires using existing network")
	}

	subnetID := clients.EnvOS.GetEnv("SUBNET_ID")
	if subnetID == "" {
		t.Skip("OS_SUBNET_ID env var is missing but DWS test requires using existing network")
	}

	name := tools.RandomString("dws-test-", 3)
	newCluster, err := cluster.CreateCluster(client, cluster.CreateClusterOpts{
		NodeType:        "dws.m3.xlarge",
		NumberOfNode:    3,
		SubnetId:        subnetID,
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
		VpcId:           vpcID,
		Name:            name,
		UserName:        "dbAdmin",
		UserPwd:         "#dbAdmin123",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = cluster.DeleteCluster(client, cluster.DeleteClusterOpts{
			ClusterId:              newCluster,
			KeepLastManualSnapshot: 0,
		})
		th.AssertNoErr(t, err)
	})

	err = cluster.WaitForCluster(client, newCluster, 600)
	th.AssertNoErr(t, err)
}
