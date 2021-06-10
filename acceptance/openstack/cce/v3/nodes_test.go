package v3

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/keypairs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

type testNodes struct {
	suite.Suite

	vpcID     string
	subnetID  string
	clusterID string
	kmsID     string
}

func TestNodes(t *testing.T) {
	suite.Run(t, new(testNodes))
}

func (s *testNodes) SetupSuite() {
	t := s.T()
	s.vpcID = clients.EnvOS.GetEnv("VPC_ID")
	s.subnetID = clients.EnvOS.GetEnv("NETWORK_ID")
	s.kmsID = clients.EnvOS.GetEnv("KMS_ID")
	if s.vpcID == "" || s.subnetID == "" {
		t.Skip("OS_VPC_ID and OS_NETWORK_ID are required for this test")
	}
	s.clusterID = createCluster(t, s.vpcID, s.subnetID)
}

func (s *testNodes) TearDownSuite() {
	t := s.T()
	if s.clusterID != "" {
		deleteCluster(t, s.clusterID)
		s.clusterID = ""
	}
}

func (s *testNodes) TestNodeLifecycle() {
	t := s.T()
	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	privateIP := "192.168.1.12" // suppose used subnet is 192.168.0.0/16

	kp := createKeypair(t)
	defer deleteKeypair(t, kp)

	var encryption string
	if s.kmsID != "" {
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
					Metadata: map[string]interface{}{
						"__system__encrypted": encryption,
						"__system__cmkid":     s.kmsID,
					},
				},
			},
			Count: 1,
			NodeNicSpec: nodes.NodeNicSpec{
				PrimaryNic: nodes.PrimaryNic{
					SubnetId: s.subnetID,
					FixedIPs: []string{privateIP},
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

	th.AssertNoErr(t, nodes.Delete(client, s.clusterID, nodeID).ExtractErr())

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
}

func createKeypair(t *testing.T) string {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	opts := keypairs.CreateOpts{
		Name: tools.RandomString("cce-nodes-", 4),
	}
	_, err = keypairs.Create(client, opts).Extract()
	th.AssertNoErr(t, err)
	return opts.Name
}

func deleteKeypair(t *testing.T, kp string) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, keypairs.Delete(client, kp).ExtractErr())
}
