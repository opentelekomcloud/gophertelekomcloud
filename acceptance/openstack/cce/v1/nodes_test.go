package v1

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	nodesv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v1/nodes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/stretchr/testify/suite"
)

type testNodes struct {
	suite.Suite

	vpcID     string
	subnetID  string
	clusterID string
}

func TestNodes(t *testing.T) {
	suite.Run(t, new(testNodes))
}

func (s *testNodes) SetupSuite() {
	t := s.T()
	s.vpcID = clients.EnvOS.GetEnv("VPC_ID")
	s.subnetID = clients.EnvOS.GetEnv("NETWORK_ID")
	if s.vpcID == "" || s.subnetID == "" {
		t.Skip("OS_VPC_ID and OS_NETWORK_ID are required for this test")
	}
	s.clusterID = cce.CreateCluster(t, s.vpcID, s.subnetID)
}

func (s *testNodes) TearDownSuite() {
	t := s.T()
	if s.clusterID != "" {
		cce.DeleteCluster(t, s.clusterID)
		s.clusterID = ""
	}
}

func (s *testNodes) TestNodeLifecycle() {
	t := s.T()
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	privateIP := "192.168.1.12" // suppose used subnet is 192.168.0.0/16

	kp := cce.CreateKeypair(t)
	defer cce.DeleteKeypair(t, kp)

	opts := nodes.CreateOpts{
		Kind:       "Node",
		ApiVersion: "v3",
		Metadata: nodes.CreateMetaData{
			Name: "nodes-test",
		},
		Spec: nodes.Spec{
			Flavor: "s2.xlarge.2",
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
				},
			},
			Count: 1,
			NodeNicSpec: nodes.NodeNicSpec{
				PrimaryNic: nodes.PrimaryNic{
					SubnetId: s.subnetID,
					FixedIPs: []string{privateIP},
				},
			},
			K8sTags: map[string]string{
				"app": "sometag",
			},
			Taints: []nodes.TaintSpec{
				{
					Key:    "dedicated",
					Value:  "database",
					Effect: "NoSchedule",
				},
			},
		},
	}

	node, err := nodes.Create(client, s.clusterID, opts).Extract()
	th.AssertNoErr(t, err)

	nodeID := node.Metadata.Id

	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		n, err := nodes.Get(client, s.clusterID, nodeID).Extract()
		if err != nil {
			return false, err
		}
		if n.Status.Phase == "Active" {
			return true, nil
		}
		time.Sleep(10 * time.Second)
		return false, nil
	}))

	state, err := nodes.Get(client, s.clusterID, nodeID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, privateIP, state.Status.PrivateIP)

	defer func() {
		th.AssertNoErr(t, nodes.Delete(client, s.clusterID, nodeID).Err)
		err = golangsdk.WaitFor(1800, func() (bool, error) {
			_, err := nodes.Get(client, s.clusterID, nodeID).Extract()
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault404); ok {
					return true, nil
				}
				return false, err
			}
			return false, nil
		})
		th.AssertNoErr(t, err)
	}()

	clientV1, err := clients.NewCceV1Client()
	th.AssertNoErr(t, err)

	k8Name := state.Status.PrivateIP
	k8Node, err := nodesv1.Get(clientV1, s.clusterID, k8Name).Extract()
	th.AssertNoErr(t, err)
	val, ok := k8Node.Metadata.Labels["app"]
	th.AssertEquals(t, true, ok)
	th.AssertEquals(t, val, opts.Spec.K8sTags["app"])
}
