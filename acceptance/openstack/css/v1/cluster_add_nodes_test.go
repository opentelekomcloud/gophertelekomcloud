package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddClusterClientNodes(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	nodeType := "ess-client"
	nodeSize := 1
	volumeType := "HIGH"

	cluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	for _, instance := range cluster.Instances {
		if instance.Type == nodeType {
			t.Skip("Cluster already contains ess-client nodes.")
		}
	}

	Versions, err := flavors.List(client)
	th.AssertNoErr(t, err)
	filteredVersions := flavors.FilterVersions(Versions, flavors.FilterOpts{
		Version: cluster.Datastore.Version,
		Type:    nodeType,
	})

	_, err = clusters.AddClusterNodes(client, clusterID, nodeType, clusters.AddNodesOpts{
		Flavor:     filteredVersions[0].Flavors[0].FlavorID,
		NodeSize:   nodeSize,
		VolumeType: volumeType,
	})
	th.AssertNoErr(t, err)

	timeout := 1200

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestAddClusterMasterNodes(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	nodeType := "ess-master"
	nodeSize := 3
	volumeType := "HIGH"

	cluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	for _, instance := range cluster.Instances {
		if instance.Type == nodeType {
			t.Skip("Cluster already contains ess-master nodes.")
		}
	}

	Versions, err := flavors.List(client)
	th.AssertNoErr(t, err)
	filteredVersions := flavors.FilterVersions(Versions, flavors.FilterOpts{
		Version: cluster.Datastore.Version,
		Type:    nodeType,
	})

	_, err = clusters.AddClusterNodes(client, clusterID, nodeType, clusters.AddNodesOpts{
		Flavor:     filteredVersions[0].Flavors[0].FlavorID,
		NodeSize:   nodeSize,
		VolumeType: volumeType,
	})
	th.AssertNoErr(t, err)

	timeout := 1200

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
