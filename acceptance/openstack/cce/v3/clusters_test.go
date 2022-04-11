package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cce"
	"github.com/stretchr/testify/suite"
)

type testCluster struct {
	suite.Suite

	vpcID       string
	subnetID    string
	clusterID   string
	eniSubnetID string
	eniCidr     string
}

func TestCluster(t *testing.T) {
	suite.Run(t, new(testCluster))
}

func (s *testCluster) SetupSuite() {
	t := s.T()
	s.vpcID = clients.EnvOS.GetEnv("VPC_ID")
	s.subnetID = clients.EnvOS.GetEnv("NETWORK_ID")
	s.eniSubnetID = clients.EnvOS.GetEnv("ENI_SUBNET_ID")
	s.eniCidr = "10.0.0.0/14"
	if s.vpcID == "" || s.subnetID == "" || s.eniSubnetID == "" {
		t.Skip("OS_VPC_ID, OS_NETWORK_ID and OS_ENI_SUBNET_ID are required for this test")
	}
	s.clusterID = cce.CreateTurboCluster(t, s.vpcID, s.subnetID, s.eniSubnetID, s.eniCidr)
}

func (s *testCluster) TearDownSuite() {
	t := s.T()
	if s.clusterID != "" {
		cce.DeleteCluster(t, s.clusterID)
		s.clusterID = ""
	}
}
