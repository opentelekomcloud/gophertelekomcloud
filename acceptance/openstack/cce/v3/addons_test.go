package v3

import (
	"encoding/json"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/addons"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddonsLifecycle(t *testing.T) {

	clusterID := clients.EnvOS.GetEnv("CLUSTER_ID")
	if clusterID == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID, and OS_CLUSTER_ID are required for this test")
	}

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	custom := map[string]interface{}{
		"coresTotal":                     32000,
		"maxEmptyBulkDeleteFlag":         10,
		"maxNodesTotal":                  1000,
		"memoryTotal":                    128000,
		"scaleDownDelayAfterAdd":         10,
		"scaleDownDelayAfterDelete":      10,
		"scaleDownDelayAfterFailure":     3,
		"scaleDownEnabled":               true,
		"scaleDownUnneededTime":          10,
		"scaleDownUtilizationThreshold":  0.25,
		"scaleUpCpuUtilizationThreshold": 0.8,
		"scaleUpMemUtilizationThreshold": 0.8,
		"scaleUpUnscheduledPodEnabled":   true,
		"scaleUpUtilizationEnabled":      true,
		"unremovableNodeRecheckTimeout":  5,
	}
	cOpts := addons.CreateOpts{
		Kind:       "Addon",
		ApiVersion: "v3",
		Metadata: addons.CreateMetadata{
			Annotations: addons.CreateAnnotations{
				AddonInstallType: "install",
			},
		},
		Spec: addons.RequestSpec{
			Version:           "1.17.2",
			ClusterID:         clusterID,
			AddonTemplateName: "autoscaler",
			Values: addons.Values{
				Basic: map[string]interface{}{
					"cceEndpoint":     "https://cce.eu-de.otc.t-systems.com",
					"ecsEndpoint":     "https://ecs.eu-de.otc.t-systems.com",
					"euleros_version": "2.5",
					"region":          "eu-de",
					"swr_addr":        "100.125.7.25:20202",
					"swr_user":        "hwofficial",
				},
				Advanced: custom,
			},
		},
	}

	addon, err := addons.Create(client, cOpts, clusterID)
	th.AssertNoErr(t, err)

	addonID := addon.Metadata.Id

	defer func() {
		err := addons.Delete(client, addonID, clusterID)
		th.AssertNoErr(t, err)

		th.AssertNoErr(t, addons.WaitForAddonDeleted(client, addonID, clusterID, 600))
	}()

	getAddon, err := addons.Get(client, addonID, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "autoscaler", getAddon.Spec.AddonTemplateName)
	th.AssertEquals(t, "1.17.2", getAddon.Spec.Version)
	th.AssertEquals(t, true, getAddon.Spec.Values.Advanced["scaleDownEnabled"])

	waitErr := addons.WaitForAddonRunning(client, addonID, clusterID, 600)
	th.AssertNoErr(t, waitErr)

	uOpts := addons.UpdateOpts{
		Kind:       "Addon",
		ApiVersion: "v3",
		Metadata: addons.UpdateMetadata{
			Annotations: addons.UpdateAnnotations{
				AddonUpdateType: "upgrade",
			},
		},
		Spec: addons.RequestSpec{
			Version:           cOpts.Spec.Version,
			ClusterID:         cOpts.Spec.ClusterID,
			AddonTemplateName: cOpts.Spec.AddonTemplateName,
			Values: addons.Values{
				Basic:    cOpts.Spec.Values.Basic,
				Advanced: cOpts.Spec.Values.Advanced,
			},
		},
	}
	uOpts.Spec.Values.Advanced["scaleDownEnabled"] = false
	uOpts.Spec.Values.Advanced["scaleDownDelayAfterAdd"] = 11

	_, err = addons.Update(client, addonID, clusterID, uOpts)
	th.AssertNoErr(t, err)

	getAddon2, err := addons.Get(client, addonID, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, getAddon2.Spec.Values.Advanced["scaleDownEnabled"])
	th.AssertEquals(t, 11.0, getAddon2.Spec.Values.Advanced["scaleDownDelayAfterAdd"])
}

func TestAddonsListTemplates(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CLUSTER_ID")
	if clusterID == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID, and OS_CLUSTER_ID are required for this test")
	}

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	list, err := addons.ListTemplates(client, clusterID, addons.ListOpts{})
	th.AssertNoErr(t, err)

	if len(list.Items) == 0 {
		t.Fatal("empty addon template list")
	}

	jsonList, _ := json.MarshalIndent(list.Items, "", "  ")

	t.Logf("existing addon templates:\n%s", string(jsonList))
}

func TestAddonsListInstances(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CLUSTER_ID")
	if clusterID == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID, and OS_CLUSTER_ID are required for this test")
	}

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	list, err := addons.ListAddonInstances(client, clusterID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(list.Items), 3)
	// check if listed addon exists
	_, err = addons.Get(client, list.Items[0].Metadata.ID, clusterID)
	th.AssertNoErr(t, err)
}

func TestAddonsGetTemplates(t *testing.T) {

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	templates, err := addons.GetTemplates(client)
	th.AssertNoErr(t, err)
	if len(templates.Items) == 0 {
		t.Fatal("empty addon templates list")
	}
}
