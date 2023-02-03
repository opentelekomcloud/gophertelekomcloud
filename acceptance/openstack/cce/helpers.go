package cce

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/keypairs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateCluster(t *testing.T, vpcID, subnetID string) string {
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	cluster, err := clusters.Create(client, clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name: strings.ToLower(tools.RandomString("cce-gopher-", 4)),
		},
		Spec: clusters.Spec{
			Type:   "VirtualMachine",
			Flavor: "cce.s1.small",
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:    vpcID,
				SubnetId: subnetID,
			},
			ContainerNetwork: clusters.ContainerNetworkSpec{
				Mode: "overlay_l2",
			},
			Authentication: clusters.AuthenticationSpec{
				Mode:                "rbac",
				AuthenticatingProxy: make(map[string]string),
			},
			KubernetesSvcIpRange: "10.247.0.0/16",
		},
	}).Extract()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, waitForClusterToActivate(client, cluster.Metadata.Id, 30*60))
	return cluster.Metadata.Id
}

func CreateTurboCluster(t *testing.T, vpcID, subnetID string, eniSubnetID string, eniCidr string) string {
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	cluster, err := clusters.Create(client, clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name: strings.ToLower(tools.RandomString("cce-gopher-turbo-", 4)),
		},
		Spec: clusters.Spec{
			Category: "Turbo",
			Type:     "VirtualMachine",
			Flavor:   "cce.s1.small",
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:    vpcID,
				SubnetId: subnetID,
			},
			ContainerNetwork: clusters.ContainerNetworkSpec{
				Mode: "eni",
			},
			EniNetwork: &clusters.EniNetworkSpec{
				SubnetId: eniSubnetID,
				Cidr:     eniCidr,
			},
			Authentication: clusters.AuthenticationSpec{
				Mode:                "rbac",
				AuthenticatingProxy: make(map[string]string),
			},
			KubernetesSvcIpRange: "10.247.0.0/16",
		},
	}).Extract()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, waitForClusterToActivate(client, cluster.Metadata.Id, 30*60))
	return cluster.Metadata.Id
}

func DeleteCluster(t *testing.T, clusterID string) {
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)
	clusters.Delete(client, clusterID)
	th.AssertNoErr(t, waitForClusterToDelete(client, clusterID, 20*60))
}

func waitForClusterToActivate(client *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		cluster, err := clusters.Get(client, id).Extract()
		if err != nil {
			return false, err
		}
		if cluster == nil {
			return false, nil
		}
		if cluster.Status.Phase == "Available" {
			return true, nil
		}
		return false, nil
	})
}

func waitForClusterToDelete(client *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := clusters.Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}

func CreateKeypair(t *testing.T) string {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	opts := keypairs.CreateOpts{
		Name: tools.RandomString("cce-nodes-", 4),
	}
	_, err = keypairs.Create(client, opts)
	th.AssertNoErr(t, err)
	return opts.Name
}

func DeleteKeypair(t *testing.T, kp string) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, keypairs.Delete(client, kp))
}
