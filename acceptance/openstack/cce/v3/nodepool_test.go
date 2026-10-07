package v3

import (
	"encoding/base64"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodepools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestNodePoolLifecycle(t *testing.T) {

	clusterId := clients.EnvOS.GetEnv("CLUSTER_ID")
	if clusterId == "" {
		t.Skip("OS_CLUSTER_ID is required for this test")
	}

	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	kp := cce.CreateKeypair(t)
	defer cce.DeleteKeypair(t, kp)

	postInstallScript := `#!/bin/bash echo "New postinstall"`
	postInstallEncoded := base64.StdEncoding.EncodeToString([]byte(postInstallScript))

	createOpts := nodepools.CreateOpts{
		Kind:       "NodePool",
		ApiVersion: "v3",
		Metadata: nodepools.CreateMetaData{
			Name: "nodepool-test",
		},
		Spec: nodepools.CreateSpec{
			Type: "vm",
			NodeTemplate: nodes.Spec{
				ExtendParam: nodes.ExtendParam{
					MaxPods:     55,
					IsAutoRenew: "false",
					IsAutoPay:   "false",
					PostInstall: postInstallEncoded,
				},
				Flavor: "s2.large.2",
				Az:     "eu-de-01",
				Os:     "EulerOS 2.5",
				Login: nodes.LoginSpec{
					SshKey: kp,
				},
				RootVolume: nodes.VolumeSpec{
					Size:       40,
					VolumeType: "SSD",
				},
				DataVolumes: []nodes.VolumeSpec{
					{
						Size:       100,
						VolumeType: "SSD",
						ExtendParam: map[string]interface{}{
							"useType": "docker",
						},
					},
					{
						Size:       100,
						VolumeType: "SSD",
					},
				},
				Count: 1,
				Storage: &nodes.Storage{
					StorageSelectors: []nodes.StorageSelector{
						{
							Name:        "cceUse",
							StorageType: "evs",
							MatchLabels: &nodes.MatchLabels{
								Size:       "100",
								VolumeType: "SSD",
								Count:      "1",
							},
						},
					},
					StorageGroups: []nodes.StorageGroup{
						{
							Name:       "vgpaas",
							CceManaged: true,
							SelectorNames: []string{
								"cceUse",
							},
							VirtualSpaces: []nodes.VirtualSpace{
								{
									Name: "runtime",
									Size: "90%",
								},
								{
									Name: "kubernetes",
									Size: "10%",
								},
							},
						},
					},
				},
			},
			InitialNodeCount: 1,
			ExtensionScaleGroups: []nodepools.ExtensionScaleGroup{
				{
					Metadata: &nodepools.ExtensionScaleGroupMetadata{
						Name: "nodepool-test-extension",
					},
					Spec: &nodepools.ExtensionScaleGroupSpec{
						Flavor: "s2.xlarge.2",
						Az:     "eu-de-02",
					},
				},
			},
		},
	}

	existingNodepools, err := nodepools.List(client, clusterId, nodepools.ListOpts{})
	th.AssertNoErr(t, err)
	numExistingNodepools := len(existingNodepools)

	nodePool, err := nodepools.Create(client, clusterId, createOpts)
	th.AssertNoErr(t, err)

	nodeId := nodePool.Metadata.Id

	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		n, err := nodepools.Get(client, clusterId, nodeId)
		if err != nil {
			return false, err
		}
		if n.Status.Phase == "" {
			return true, nil
		}
		time.Sleep(10 * time.Second)
		return false, nil
	}))

	nodepoolList, err := nodepools.List(client, clusterId, nodepools.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, numExistingNodepools <= len(nodepoolList))

	pool, err := nodepools.Get(client, clusterId, nodeId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 55, pool.Spec.NodeTemplate.ExtendParam.MaxPods)
	th.AssertEquals(t, postInstallEncoded, pool.Spec.NodeTemplate.ExtendParam.PostInstall)
	extensionScaleGroup := assertExtensionScaleGroup(t, pool.Spec.ExtensionScaleGroups)
	if extensionScaleGroup.Metadata.Uid == "" {
		t.Fatal("missing extension scale group UID")
	}
	// Not supported params by now
	// th.AssertEquals(t, "false", pool.Spec.NodeTemplate.ExtendParam.IsAutoPay)
	// th.AssertEquals(t, "false", pool.Spec.NodeTemplate.ExtendParam.IsAutoRenew)

	updatedPostInstallScript := `#!/bin/bash echo "Updated postinstall"`
	updatedPostInstallEncoded := base64.StdEncoding.EncodeToString([]byte(updatedPostInstallScript))
	updatedExtensionScaleGroups := []nodepools.ExtensionScaleGroup{
		{
			Metadata: &nodepools.ExtensionScaleGroupMetadata{
				Uid:  extensionScaleGroup.Metadata.Uid,
				Name: extensionScaleGroup.Metadata.Name,
			},
			Spec: &nodepools.ExtensionScaleGroupSpec{
				Flavor: extensionScaleGroup.Spec.Flavor,
				Az:     extensionScaleGroup.Spec.Az,
			},
		},
	}

	updateOpts := nodepools.UpdateOpts{
		Metadata: nodepools.UpdateMetaData{
			Name: "nodepool-test-updated",
		},
		Spec: nodepools.UpdateSpec{
			InitialNodeCount:     1,
			ExtensionScaleGroups: &updatedExtensionScaleGroups,
		},
	}

	updateOpts.Spec.NodeTemplate = createOpts.Spec.NodeTemplate
	updateOpts.Spec.NodeTemplate.ExtendParam.PostInstall = updatedPostInstallEncoded

	updatedPool, err := nodepools.Update(client, clusterId, nodeId, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "nodepool-test-updated", updatedPool.Metadata.Name)
	th.AssertEquals(t, updatedPostInstallEncoded, updatedPool.Spec.NodeTemplate.ExtendParam.PostInstall)
	updatedExtensionScaleGroup := assertExtensionScaleGroup(t, updatedPool.Spec.ExtensionScaleGroups)
	th.AssertEquals(t, extensionScaleGroup.Metadata.Uid, updatedExtensionScaleGroup.Metadata.Uid)

	getUpdatedPool, err := nodepools.Get(client, clusterId, nodeId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedPostInstallEncoded, getUpdatedPool.Spec.NodeTemplate.ExtendParam.PostInstall)
	getUpdatedExtensionScaleGroup := assertExtensionScaleGroup(t, getUpdatedPool.Spec.ExtensionScaleGroups)
	th.AssertEquals(t, extensionScaleGroup.Metadata.Uid, getUpdatedExtensionScaleGroup.Metadata.Uid)

	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		n, err := nodepools.Get(client, clusterId, nodeId)
		if err != nil {
			return false, err
		}
		if n.Status.Phase == "" {
			return true, nil
		}
		time.Sleep(10 * time.Second)
		return false, nil
	}))

	th.AssertNoErr(t, nodepools.Delete(client, clusterId, nodeId))

	err = golangsdk.WaitFor(1800, func() (bool, error) {
		_, err := nodepools.Get(client, clusterId, nodeId)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
	th.AssertNoErr(t, err)
}

func assertExtensionScaleGroup(t *testing.T, extensionScaleGroups []nodepools.ExtensionScaleGroup) nodepools.ExtensionScaleGroup {
	t.Helper()

	if len(extensionScaleGroups) != 1 {
		t.Fatalf("expected one extension scale group, got %d", len(extensionScaleGroups))
	}
	extensionScaleGroup := extensionScaleGroups[0]
	if extensionScaleGroup.Metadata == nil {
		t.Fatal("missing extension scale group metadata")
	}
	if extensionScaleGroup.Spec == nil {
		t.Fatal("missing extension scale group specification")
	}
	th.AssertEquals(t, "nodepool-test-extension", extensionScaleGroup.Metadata.Name)
	th.AssertEquals(t, "s2.xlarge.2", extensionScaleGroup.Spec.Flavor)
	th.AssertEquals(t, "eu-de-02", extensionScaleGroup.Spec.Az)

	return extensionScaleGroup
}
