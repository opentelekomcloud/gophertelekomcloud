package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	nwv1 "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/networking/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/mrs/v1/cluster"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMrsClusterLifecycle(t *testing.T) {
	client, err := clients.NewMrsV1()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	vpc, subnet := nwv1.CreateNetwork(t, "mrs", az)
	defer nwv1.DeleteNetwork(t, subnet)

	name := tools.RandomString("css-create", 3)
	createOpts := cluster.CreateOpts{
		BillingType:        12,
		DataCenter:         cc.RegionName,
		MasterNodeNum:      2,
		MasterNodeSize:     "s2.4xlarge.2.linux.mrs",
		CoreNodeNum:        3,
		CoreNodeSize:       "s2.4xlarge.2.linux.mrs",
		AvailableZoneID:    az,
		ClusterName:        name,
		Vpc:                vpc.Name,
		VpcID:              vpc.ID,
		SubnetID:           subnet.NetworkID,
		SubnetName:         subnet.Name,
		ClusterVersion:     "MRS 2.1.0",
		ClusterType:        0,
		VolumeType:         "SATA",
		VolumeSize:         100,
		SafeMode:           1,
		ClusterAdminSecret: "Qwerty123!",
		LoginMode:          0,
		ComponentList: []cluster.ComponentOpts{
			{
				ComponentName: "Hadoop",
			},
			{
				ComponentName: "Spark",
			},
			{
				ComponentName: "Hive",
			},
		},
	}

	clResponse, err := cluster.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = waitForClusterToBeActive(client, clResponse.ClusterID, 600)
	th.AssertNoErr(t, err)

	defer func() {
		err = cluster.Delete(client, clResponse.ClusterID).ExtractErr()
		th.AssertNoErr(t, err)
		err = waitForClusterToBeDeleted(client, clResponse.ClusterID, 600)
		th.AssertNoErr(t, err)
	}()

	newCluster, err := cluster.Get(client, clResponse.ClusterID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(newCluster.ComponentList), 3)
}

func waitForClusterToBeActive(client *golangsdk.ServiceClient, clusterID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		n, err := cluster.Get(client, clusterID).Extract()
		if err != nil {
			return false, err
		}

		if n.ClusterState == "running" {
			return true, nil
		}

		return false, nil
	})
}

func waitForClusterToBeDeleted(client *golangsdk.ServiceClient, clusterID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := cluster.Get(client, clusterID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}
