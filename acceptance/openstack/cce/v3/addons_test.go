package v3

import (
	"fmt"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/addons"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddonsLifecycle(t *testing.T) {
	t.SkipNow()
	vpcClient, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	cceClient, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)
	addonClient, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	vpc, err := vpcs.Create(vpcClient, vpcs.CreateOpts{
		Name: tools.RandomString("cce-vpc-", 4),
		CIDR: "192.168.0.0/16",
	}).Extract()
	th.AssertNoErr(t, err)
	defer vpcs.Delete(vpcClient, vpc.ID)

	subnet, err := subnets.Create(vpcClient, subnets.CreateOpts{
		Name:             tools.RandomString("cce-subnet-", 4),
		CIDR:             "192.168.0.0/24",
		DnsList:          []string{"1.1.1.1", "8.8.8.8"},
		GatewayIP:        "192.168.0.1",
		EnableDHCP:       true,
		AvailabilityZone: "eu-de-01",
		VPC_ID:           vpc.ID,
	}).Extract()
	th.AssertNoErr(t, err)
	defer subnets.Delete(vpcClient, vpc.ID, subnet.ID)
	th.AssertNoErr(t, waitForSubnetToActivate(vpcClient, subnet.ID, 60))

	cluster, err := clusters.Create(cceClient, clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name: strings.ToLower(tools.RandomString("addontest-", 4)),
		},
		Spec: clusters.Spec{
			Type:   "VirtualMachine",
			Flavor: "cce.s1.small",
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:    vpc.ID,
				SubnetId: subnet.ID,
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
	clusterID := cluster.Metadata.Id
	th.AssertNoErr(t, waitForClusterToActivate(cceClient, clusterID, 30*60))
	defer clusters.Delete(cceClient, clusterID)

	addon, err := addons.Create(addonClient, addons.CreateOpts{
		Kind:       "Addon",
		ApiVersion: "v3",
		Metadata: addons.CreateMetadata{
			Annotations: addons.CreateAnnotations{
				AddonInstallType: "install",
			},
		},
		Spec: addons.RequestSpec{
			Version:           "1.0.3",
			ClusterID:         clusterID,
			AddonTemplateName: "metrics-server",
			Values: addons.Values{
				Basic: map[string]interface{}{
					"euleros_version": "2.5",
					"rbac_enabled":    true,
					"swr_addr":        "100.125.7.25:20202",
					"swr_user":        "hwofficial",
				},
			},
		},
	}, cluster.Metadata.Id).Extract()
	th.AssertNoErr(t, err)

	addonID := addon.Metadata.Id

	defer func() {
		err := addons.Delete(addonClient, addonID, clusterID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	getAddon, err := addons.Get(addonClient, addon.Metadata.Id, cluster.Metadata.Id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getAddon.Spec.AddonTemplateName, "metrics-server")
	th.AssertEquals(t, getAddon.Spec.Version, "1.0.3")
}

func waitForSubnetToActivate(client *golangsdk.ServiceClient, subnetID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		subnet, err := subnets.Get(client, subnetID).Extract()
		if err != nil {
			return false, err
		}

		if subnet.Status == "ACTIVE" {
			return true, nil
		}

		// If subnet status is other than Active, send error
		if subnet.Status == "DOWN" || subnet.Status == "ERROR" {
			return false, fmt.Errorf("subnet status: '%s'", subnet.Status)
		}

		return false, nil
	})
}

func waitForClusterToActivate(client *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		cluster, err := clusters.Get(client, id).Extract()
		if err != nil {
			return false, err
		}
		if cluster.Status.Phase == "Available" {
			return true, nil
		}
		return false, nil
	})
}
