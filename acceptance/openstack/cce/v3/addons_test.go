package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/addons"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddonsLifecycle(t *testing.T) {

	clusterID := clients.EnvOS.GetEnv("CLUSTER_ID")
	tenantID := clients.EnvOS.GetEnv("TENANT_ID")
	if clusterID == "" || tenantID == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID, and OS_CLUSTER_ID are required for this test")
	}

	client, err := clients.NewCceV3AddonClient()
	th.AssertNoErr(t, err)

	createOpts := addons.CreateOpts{
		Kind:       "Addon",
		ApiVersion: "v3",
		Metadata: addons.CreateMetadata{
			Annotations: addons.CreateAnnotations{
				AddonInstallType: "install",
			},
		},
		Spec: addons.RequestSpec{
			Version:           "1.19.1",
			ClusterID:         clusterID,
			AddonTemplateName: "npd",
			Values: addons.Values{
				Basic: map[string]interface{}{
					"image_version": "1.19.1",
					"swr_addr":      "100.125.7.25:20202",
					"swr_user":      "cce-addons",
				},
				Advanced: map[string]interface{}{
					"multiAZBalance":         false,
					"multiAZEnabled":         false,
					"feature_gates":          "",
					"node_match_expressions": []interface{}{},
					"npc": map[string]interface{}{
						"maxTaintedNode": "10%",
					},
					"tolerations": []map[string]interface{}{
						{
							"effect":            "NoExecute",
							"key":               "node.kubernetes.io/not-ready",
							"operator":          "Exists",
							"tolerationSeconds": 60,
						},
						{
							"effect":            "NoExecute",
							"key":               "node.kubernetes.io/unreachable",
							"operator":          "Exists",
							"tolerationSeconds": 60,
						},
					},
				},
			},
		},
	}

	listExistingAddons, err := addons.ListAddonInstances(client, clusterID)
	th.AssertNoErr(t, err)

	existingAddons := len(listExistingAddons.Items)

	addon, err := addons.Create(client, createOpts, clusterID)
	th.AssertNoErr(t, err)

	addonID := addon.Metadata.Id

	defer func() {
		err := addons.Delete(client, addonID, clusterID)
		th.AssertNoErr(t, err)

		th.AssertNoErr(t, addons.WaitForAddonDeleted(client, addonID, clusterID, 1200))
	}()

	listAddons, err := addons.ListAddonInstances(client, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, existingAddons+1, len(listAddons.Items))

	getAddon, err := addons.Get(client, addonID, clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "npd", getAddon.Spec.AddonTemplateName)
	th.AssertEquals(t, "1.19.1", getAddon.Spec.Version)

	waitErr := addons.WaitForAddonRunning(client, addonID, clusterID, 1200)
	th.AssertNoErr(t, waitErr)

	updateOpts := addons.UpdateOpts{
		Kind:       "Addon",
		ApiVersion: "v3",
		Metadata: addons.UpdateMetadata{
			Annotations: addons.UpdateAnnotations{
				AddonUpdateType: "upgrade",
			},
		},
		Spec: addons.RequestSpec{
			Version:           createOpts.Spec.Version,
			ClusterID:         createOpts.Spec.ClusterID,
			AddonTemplateName: createOpts.Spec.AddonTemplateName,
			Values: addons.Values{
				Basic:    createOpts.Spec.Values.Basic,
				Advanced: createOpts.Spec.Values.Advanced,
			},
		},
	}

	updateOpts.Spec.Values.Basic["rbac_enabled"] = true

	getAddon2, err := addons.Update(client, addonID, clusterID, updateOpts)
	th.AssertNoErr(t, err)

	// USE THIS TO DEBUG
	// addonJson, _ := json.MarshalIndent(getAddon2, "", "  ")
	// t.Logf("existing addon templates:\n%s", string(addonJson))
	th.AssertEquals(t, true, getAddon2.Spec.Values.Basic["rbac_enabled"])
	waitErr = addons.WaitForAddonRunning(client, addonID, clusterID, 1200)
	th.AssertNoErr(t, waitErr)
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

	// jsonList, _ := json.MarshalIndent(list.Items, "", "  ")
	// t.Logf("existing addon templates:\n%s", string(jsonList))
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
