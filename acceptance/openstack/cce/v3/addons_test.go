package v3

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/addons"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

type testAddons struct {
	suite.Suite

	vpcID     string
	subnetID  string
	clusterID string
}

func TestAddons(t *testing.T) {
	suite.Run(t, new(testAddons))
}

func (a *testAddons) SetupSuite() {
	t := a.T()
	a.vpcID = clients.EnvOS.GetEnv("VPC_ID")
	a.subnetID = clients.EnvOS.GetEnv("NETWORK_ID")
	if a.vpcID == "" || a.subnetID == "" {
		t.Skip("OS_VPC_ID and OS_NETWORK_ID are required for this test")
	}
	a.clusterID = createCluster(t, a.vpcID, a.subnetID)
}

func (a *testAddons) TearDownSuite() {
	t := a.T()
	if a.clusterID != "" {
		deleteCluster(t, a.clusterID)
		a.clusterID = ""
	}
}

func (a *testAddons) TestAddonsLifecycle() {
	t := a.T()

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	addon, err := addons.Create(client, addons.CreateOpts{
		Kind:       "Addon",
		ApiVersion: "v3",
		Metadata: addons.CreateMetadata{
			Annotations: addons.CreateAnnotations{
				AddonInstallType: "install",
			},
		},
		Spec: addons.RequestSpec{
			Version:           "1.0.3",
			ClusterID:         a.clusterID,
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
	}, a.clusterID).Extract()
	th.AssertNoErr(t, err)

	addonID := addon.Metadata.Id

	defer func() {
		err := addons.Delete(client, addonID, a.clusterID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	getAddon, err := addons.Get(client, addon.Metadata.Id, a.clusterID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getAddon.Spec.AddonTemplateName, "metrics-server")
	th.AssertEquals(t, getAddon.Spec.Version, "1.0.3")
}

func (a *testAddons) TestListAddonTemplates() {
	t := a.T()

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	list, err := addons.ListTemplates(client, a.clusterID, nil).Extract()
	th.AssertNoErr(t, err)

	if len(list.Items) == 0 {
		t.Fatal("empty addon template list")
	}

	jsonList, _ := json.MarshalIndent(list.Items, "", "  ")

	t.Logf("existing addon templates:\n%s", string(jsonList))
}

func createCluster(t *testing.T, vpcID, subnetID string) string {
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	cluster, err := clusters.Create(client, clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name: strings.ToLower(tools.RandomString("addontest-", 4)),
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

func deleteCluster(t *testing.T, clusterID string) {
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
