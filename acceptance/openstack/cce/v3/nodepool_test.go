package v3

import (
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
	th.AssertEquals(t, numExistingNodepools+1, len(nodepoolList))

	pool, err := nodepools.Get(client, clusterId, nodeId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 55, pool.Spec.NodeTemplate.ExtendParam.MaxPods)
	// Not supported params by now
	// th.AssertEquals(t, "false", pool.Spec.NodeTemplate.ExtendParam.IsAutoPay)
	// th.AssertEquals(t, "false", pool.Spec.NodeTemplate.ExtendParam.IsAutoRenew)

	updateOpts := nodepools.UpdateOpts{
		Metadata: nodepools.UpdateMetaData{
			Name: "nodepool-test-updated",
		},
		Spec: nodepools.UpdateSpec{
			InitialNodeCount: 1,
			NodeTemplate:     nodepools.UpdateNodeTemplate{},
		},
	}
	updatedPool, err := nodepools.Update(client, clusterId, nodeId, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "nodepool-test-updated", updatedPool.Metadata.Name)
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
