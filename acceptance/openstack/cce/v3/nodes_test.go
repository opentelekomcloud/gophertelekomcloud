package v3

import (
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestNodeLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("OS_VPC_ID and OS_NETWORK_ID are required for this test")
	}
	clusterID := cce.CreateCluster(t, vpcID, subnetID)
	t.Cleanup(func() {
		cce.DeleteCluster(t, clusterID)
	})

	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	nodesListInit, err := nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	initNodeNum := len(nodesListInit)

	privateIP := "192.168.0.12" // suppose used subnet is 192.168.0.0/16
	kp := cce.CreateKeypair(t)
	defer cce.DeleteKeypair(t, kp)
	var encryption string
	if kmsID != "" {
		encryption = "1"
	} else {
		encryption = "0"
	}
	opts := nodes.CreateOpts{
		Kind:       "Node",
		ApiVersion: "v3",
		Metadata: nodes.CreateMetaData{
			Name: "nodes-test",
		},
		Spec: nodes.Spec{
			Flavor: "s2.xlarge.2",
			Az:     "eu-de-01",
			Os:     "EulerOS 2.9",
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
					Metadata: map[string]interface{}{
						"__system__encrypted": encryption,
						"__system__cmkid":     kmsID,
					},
				},
			},
			Count: 1,
			NodeNicSpec: nodes.NodeNicSpec{
				PrimaryNic: nodes.PrimaryNic{
					SubnetId: subnetID,
					FixedIPs: []string{privateIP},
				},
			},
			Runtime: nodes.RuntimeSpec{
				Name: "containerd",
			},
			ExtendParam: nodes.ExtendParam{
				MaxPods:        16,
				DockerBaseSize: 20,
			},
		},
	}
	if v := os.Getenv("OS_AGENCY_NAME"); v != "" {
		opts.Spec.ExtendParam.AgencyName = v
	}
	node, err := nodes.Create(client, clusterID, opts)
	th.AssertNoErr(t, err)

	nodeID := node.Metadata.Id

	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		n, err := nodes.Get(client, clusterID, nodeID)
		if err != nil {
			return false, err
		}
		if n.Status.Phase == "Active" {
			return true, nil
		}
		time.Sleep(10 * time.Second)
		return false, nil
	}))

	nodesList, err := nodes.List(client, clusterID, nodes.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, initNodeNum+1, len(nodesList))

	state, err := nodes.Get(client, clusterID, nodeID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, privateIP, state.Status.PrivateIP)

	updatedNode, err := nodes.Update(client, clusterID, nodeID, nodes.UpdateOpts{
		Metadata: nodes.UpdateMetadata{
			Name: "node-updated",
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "node-updated", updatedNode.Metadata.Name)

	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		n, err := nodes.Get(client, clusterID, nodeID)
		if err != nil {
			return false, err
		}
		if n.Status.Phase == "Active" {
			return true, nil
		}
		time.Sleep(10 * time.Second)
		return false, nil
	}))

	th.AssertNoErr(t, nodes.Delete(client, clusterID, nodeID))

	err = golangsdk.WaitFor(1800, func() (bool, error) {
		_, err := nodes.Get(client, clusterID, nodeID)
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
