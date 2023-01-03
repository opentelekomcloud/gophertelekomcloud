package v1

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/snapshot"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDWS(t *testing.T) {
	if os.Getenv("RUN_DWS_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	client, err := clients.NewDWSV1Client()
	th.AssertNoErr(t, err)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID is missing but DWS test requires using existing network")
	}

	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if subnetID == "" {
		t.Skip("OS_SUBNET_ID env var is missing but DWS test requires using existing network")
	}

	// clusterId := "e3ececc2-e2ae-4f5c-8ab5-fc96f4f20f4f"

	t.Log("Creating cluster")

	name := tools.RandomString("dws-test-", 3)
	clusterId, err := cluster.CreateCluster(client, cluster.CreateClusterOpts{
		NodeType:        "dws.m3.xlarge",
		NumberOfNode:    3,
		SubnetId:        subnetID,
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
		VpcId:           vpcID,
		Name:            name,
		UserName:        "dbadmin",
		UserPwd:         "#dbadmin123",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = golangsdk.WaitFor(1000, func() (bool, error) {
			err = cluster.DeleteCluster(client, cluster.DeleteClusterOpts{
				ClusterId:              clusterId,
				KeepLastManualSnapshot: pointerto.Int(0),
			})
			if err != nil {
				t.Error(err)
				return false, nil
			}
			return true, nil
		})
		th.AssertNoErr(t, err)
	})

	err = cluster.WaitForCluster(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	err = cluster.ResetPassword(client, cluster.ResetPasswordOpts{
		ClusterId:   clusterId,
		NewPassword: "#SomePassword123",
	})
	th.AssertNoErr(t, err)

	err = cluster.WaitForCluster(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	err = cluster.ResizeCluster(client, cluster.ResizeClusterOpts{
		ClusterId: clusterId,
		Count:     1,
	})

	err = cluster.WaitForCluster(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	err = cluster.RestartCluster(client, cluster.RestartClusterOpts{
		ClusterId: clusterId,
		Restart:   struct{}{},
	})
	th.AssertNoErr(t, err)

	err = cluster.WaitForCluster(client, clusterId, 1000)
	th.AssertNoErr(t, err)

	list, err := cluster.ListClusters(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)

	snapId, err := snapshot.CreateSnapshot(client, snapshot.Snapshot{
		Name:      name,
		ClusterId: clusterId,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = snapshot.DeleteSnapshot(client, snapId)
		th.AssertNoErr(t, err)
	})

	err = snapshot.WaitForSnapshot(client, snapId, 1000)
	th.AssertNoErr(t, err)

	newName := tools.RandomString("dws-test-", 3)
	resCId, err := snapshot.RestoreCluster(client, snapshot.RestoreClusterOpts{
		SnapshotId: snapId,
		Name:       newName,
	})
	t.Cleanup(func() {
		err = golangsdk.WaitFor(1000, func() (bool, error) {
			err = cluster.DeleteCluster(client, cluster.DeleteClusterOpts{
				ClusterId:              resCId,
				KeepLastManualSnapshot: pointerto.Int(0),
			})
			if err != nil {
				t.Error(err)
				return false, nil
			}
			return true, nil
		})
		th.AssertNoErr(t, err)
	})

	err = cluster.WaitForCluster(client, resCId, 1000)
	th.AssertNoErr(t, err)

	snaps, err := snapshot.ListSnapshot(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, snaps)
}
