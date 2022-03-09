package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testKubeConfig struct {
	suite.Suite

	vpcID     string
	subnetID  string
	clusterID string
}

func TestKubeConfig(t *testing.T) {
	suite.Run(t, new(testKubeConfig))
}

func (s *testKubeConfig) SetupSuite() {
	t := s.T()
	s.vpcID = clients.EnvOS.GetEnv("VPC_ID")
	s.subnetID = clients.EnvOS.GetEnv("NETWORK_ID")
	if s.vpcID == "" || s.subnetID == "" {
		t.Skip("OS_VPC_ID and OS_NETWORK_ID are required for this test")
	}
	s.clusterID = cce.CreateCluster(t, s.vpcID, s.subnetID)
}

func (s *testKubeConfig) TearDownSuite() {
	t := s.T()
	if s.clusterID != "" {
		cce.DeleteCluster(t, s.clusterID)
		s.clusterID = ""
	}
}

func (s *testKubeConfig) TestKubeConfigReading() {
	t := s.T()

	client, err := clients.NewCceV3Client()
	th.AssertNoErr(t, err)

	kubeConfig, err := clusters.GetCert(client, s.clusterID).ExtractMap()
	th.AssertNoErr(t, err)
	require.NotEmpty(t, kubeConfig)
}
