package v3

import (
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodepools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func (s *testNodes) TestNodePoolLifecycle() {
	t := s.T()
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	clusterId := s.clusterID

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
						ExtendParam: map[string]any{
							"useType": "docker",
						},
					},
					{
						Size:       100,
						VolumeType: "SSD",
					},
				},
				Count: 1,
			},
			InitialNodeCount: 1,
		},
	}

	nodePool, err := nodepools.Create(client, clusterId, createOpts).Extract()
	th.AssertNoErr(t, err)

	nodeId := nodePool.Metadata.Id

	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		n, err := nodepools.Get(client, clusterId, nodeId).Extract()
		if err != nil {
			return false, err
		}
		if n.Status.Phase == "" {
			return true, nil
		}
		time.Sleep(10 * time.Second)
		return false, nil
	}))

	_, err = nodepools.Get(client, clusterId, nodeId).Extract()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, nodepools.Delete(client, clusterId, nodeId).ExtractErr())

	err = golangsdk.WaitFor(1800, func() (bool, error) {
		_, err := nodepools.Get(client, clusterId, nodeId).Extract()
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
