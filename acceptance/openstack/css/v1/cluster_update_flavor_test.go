package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateClusterFlavor(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	cluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	var (
		currentFlavor string
		newFlavorID   string
		instanceType  string = "ess"
	)

	for _, instance := range cluster.Instances {
		if instance.Type == instanceType {
			currentFlavor = instance.SpecCode
			break
		}
	}

	filterOpts := flavors.FilterOpts{
		Version: "7.10.2",
		Type:    instanceType,
	}

	versions, err := flavors.List(client)
	th.AssertNoErr(t, err)

	filteredVersions := flavors.FilterVersions(versions, filterOpts)
	if filteredVersions[0].Flavors[0].Name != currentFlavor {
		newFlavorID = filteredVersions[0].Flavors[0].FlavorID
	} else {
		newFlavorID = filteredVersions[0].Flavors[1].FlavorID
	}
	needCheckReplica := false
	err = clusters.UpdateClusterFlavor(client, clusterID, clusters.ClusterFlavorOpts{
		NeedCheckReplica: &needCheckReplica,
		NewFlavorID:      newFlavorID,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestUpdateClusterNodeFlavor(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	cluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	var (
		currentFlavor string
		newFlavorID   string
		instanceType  string
	)

	for _, instance := range cluster.Instances {
		if instance.Type == "ess-cold" || instance.Type == "ess-client" || instance.Type == "ess-master" {
			currentFlavor = instance.SpecCode
			instanceType = instance.Type
			break
		}
	}

	if instanceType == "" {
		t.Skip("There are no client, cold or master nodes to change the flavor.")
	}

	filterOpts := flavors.FilterOpts{
		Version: "7.10.2",
		Type:    instanceType,
	}

	versions, err := flavors.List(client)
	th.AssertNoErr(t, err)

	filteredVersions := flavors.FilterVersions(versions, filterOpts)
	if filteredVersions[0].Flavors[0].Name != currentFlavor {
		newFlavorID = filteredVersions[0].Flavors[0].FlavorID
	} else {
		newFlavorID = filteredVersions[0].Flavors[1].FlavorID
	}

	needCheckReplica := false
	err = clusters.UpdateClusterFlavor(client, clusterID, clusters.ClusterFlavorOpts{
		NeedCheckReplica: &needCheckReplica,
		NewFlavorID:      newFlavorID,
		NodeType:         instanceType,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
