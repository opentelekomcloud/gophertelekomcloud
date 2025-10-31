package v3

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/floatingips"
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
	th.AssertEquals(t, cluster.Metadata.Timezone, clusterGet.Metadata.Timezone)
	th.AssertEquals(t, cluster.Spec.PublicAccess.Cidrs[0], "192.168.45.0/24")

	computeClient, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	eip, err := floatingips.Create(computeClient, floatingips.CreateOpts{
		Pool: "admin_external_net",
	}).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err := floatingips.Delete(computeClient, eip.ID).ExtractErr()
		th.CheckNoErr(t, err)
	})

	updateIpOpts := clusters.UpdateIpOpts{
		Action:    "bind",
		ElasticIp: eip.ID,
	}
	updateIpOpts.Spec.ID = eip.ID
	err = clusters.UpdateMasterIp(client, clusterID, updateIpOpts)
	th.AssertNoErr(t, err)

	if clusterID != "" {
		cce.DeleteCluster(t, clusterID)
	}
}

func TestTurboClusterWithCillium(t *testing.T) {
	t.Skip("Only available in whitelisted tenants")
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

	cluster, err := clusters.Create(client, clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name:     strings.ToLower(tools.RandomString("cce-gopher-turbo-", 4)),
			Timezone: "Pacific/Auckland",
		},
		Spec: clusters.Spec{
			Category: "Turbo",
			Type:     "VirtualMachine",
			Flavor:   "cce.s1.small",
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:    vpcID,
				SubnetId: subnetsList[0].NetworkID,
			},
			ContainerNetwork: clusters.ContainerNetworkSpec{
				Mode: "eni",
			},
			EniNetwork: &clusters.EniNetworkSpec{
				SubnetId: subnetsList[0].SubnetID,
				Cidr:     subnetsList[0].CIDR,
			},
			Authentication: clusters.AuthenticationSpec{
				Mode:                "rbac",
				AuthenticatingProxy: make(map[string]string),
			},
			KubernetesSvcIpRange: "10.247.0.0/16",
			Masters: []clusters.MasterSpec{
				{
					AvailabilityZone: "eu-de-01",
				},
			},
			PublicAccess: &clusters.PublicAccess{
				Cidrs: []string{
					"192.168.45.0/24",
					"10.234.128.0/20",
				},
			},
			ConfigurationsOverride: []clusters.PackageConfiguration{
				{
					Name: "kube-apiserver",
					Configurations: []clusters.Configuration{
						{
							Name:  "support-overload",
							Value: true,
						},
					},
				},
				{
					Name: "eni",
					Configurations: []clusters.Configuration{
						{
							Name:  "dataplane-v2",
							Value: true,
						},
					},
				},
			},
		},
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, cce.WaitForClusterToActivate(client, cluster.Metadata.Id, 30*60))
	clusterID := cluster.Metadata.Id

	clusterGet, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, cluster.Metadata.Name, clusterGet.Metadata.Name)
	th.AssertEquals(t, cluster.Metadata.Timezone, clusterGet.Metadata.Timezone)
	th.AssertEquals(t, cluster.Spec.PublicAccess.Cidrs[0], "192.168.45.0/24")

	if clusterID != "" {
		cce.DeleteCluster(t, clusterID)
	}
}
