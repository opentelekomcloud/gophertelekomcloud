package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/mrs/v1/cluster"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMrsClusterLifecycle(t *testing.T) {
	client, err := clients.NewMrsV1()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-02"
	}

	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID", "ROUTER_ID")
	keyPairName := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	if networkID == "" || vpcID == "" || keyPairName == "" {
		t.Skip("OS_NETWORK_ID, OS_VPC_ID or OS_KEYPAIR_NAME env vars are missing but MRS Cluster test requires")
	}

	nwV1Client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	vpc, err := vpcs.Get(nwV1Client, vpcID).Extract()
	th.AssertNoErr(t, err)

	subnet, err := subnets.Get(nwV1Client, networkID).Extract()
	th.AssertNoErr(t, err)

	name := tools.RandomString("mrs-create-", 3)
	createOpts := cluster.CreateOpts{
		BillingType:        12,
		DataCenter:         cc.RegionName,
		MasterNodeNum:      2,
		MasterNodeSize:     "c3.xlarge.4.linux.mrs",
		CoreNodeNum:        3,
		CoreNodeSize:       "c3.xlarge.4.linux.mrs",
		AvailableZoneId:    az,
		ClusterName:        name,
		Vpc:                vpc.Name,
		VpcId:              vpc.ID,
		SubnetId:           subnet.NetworkID,
		SubnetName:         subnet.Name,
		ClusterVersion:     "MRS 2.1.0",
		ClusterType:        pointerto.Int(0),
		VolumeType:         "SATA",
		VolumeSize:         100,
		SafeMode:           1,
		ClusterAdminSecret: "Qwerty123!",
		LoginMode:          pointerto.Int(1),
		NodePublicCertName: keyPairName,
		LogCollection:      pointerto.Int(1),
		ComponentList: cluster.ExpandComponent(
			[]string{"Presto", "Hadoop", "Spark", "HBase", "Hive", "Hue", "Loader", "Tez", "Flink"},
		),
	}

	clResponse, err := cluster.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = waitForClusterToBeActive(client, clResponse.ClusterId, 3000)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = cluster.Delete(client, clResponse.ClusterId)
		th.AssertNoErr(t, err)
		err = waitForClusterToBeDeleted(client, clResponse.ClusterId, 3000)
		th.AssertNoErr(t, err)
	})

	tagOpts := []tags.ResourceTag{
		{
			Key:   "muh",
			Value: "lala",
		},
		{
			Key:   "kuh",
			Value: "lala",
		},
	}

	err = tags.Create(client, "clusters", clResponse.ClusterId, tagOpts).ExtractErr()
	th.AssertNoErr(t, err)

	newCluster, err := cluster.Get(client, clResponse.ClusterId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(newCluster.ComponentList), 9)

	tagList, err := tags.Get(client, "clusters", clResponse.ClusterId).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(tagList), len(tagOpts))
}

func waitForClusterToBeActive(client *golangsdk.ServiceClient, clusterID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		n, err := cluster.Get(client, clusterID)
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
		n, err := cluster.Get(client, clusterID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		if n.ClusterState == "terminated" {
			return true, nil
		}

		return false, nil
	})
}
