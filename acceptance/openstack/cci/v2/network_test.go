package v2

import (
	"os"
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	ns "github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/namespace"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/network"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestNetworkLifecycle(t *testing.T) {
	t.Skip("Tenant not whitelisted to run CCI")
	client, err := clients.NewCCIClient()
	th.AssertNoErr(t, err)

	subnetID := os.Getenv("OS_SUBNET_ID")

	nsName := "cci-namespace-" + strconv.Itoa(tools.RandomInt(1, 1000))
	networkName := "test-net-" + strconv.Itoa(tools.RandomInt(1, 1000))

	createNsOpts := ns.CreateOpts{
		APIVersion: "cci/v2",
		Kind:       "Namespace",
		Metadata: ns.Metadata{
			Name: nsName,
		},
	}

	t.Logf("Attempting to create namespace")
	_, err = ns.Create(client, createNsOpts)
	th.AssertNoErr(t, err)

	err = waitForStatusActive(client, 600, nsName)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete namespace")
		nsDeleteOpts := ns.DeleteOpts{
			Name: nsName,
		}
		_, err = ns.Delete(client, nsDeleteOpts)
		th.AssertNoErr(t, err)
		err = waitForStatusDeleted(client, 500, nsName)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create network")
	networkClient, err := clients.NewCCINetworkClient()
	th.AssertNoErr(t, err)

	createOpts := network.CreateOpts{
		Namespace:  nsName,
		APIVersion: "yangtse/v2",
		Kind:       "Network",
		Metadata: &network.ObjectMeta{
			Name: networkName,
			Annotations: map[string]string{
				"yangtse.io/domain-id":                  client.DomainID,
				"yangtse.io/project-id":                 client.ProjectID,
				"yangtse.io/warm-pool-recycle-interval": "1",
				"yangtse.io/warm-pool-size":             "10",
			},
		},
		Spec: &network.NetworkSpec{
			NetworkType: "underlay_neutron",
			SecurityGroups: []string{
				"default",
			},
			Subnets: []network.SubnetConf{
				{
					SubnetID: subnetID,
				},
			},
		},
	}

	cm, err := network.Create(networkClient, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, networkName, cm.Metadata.Name)

	t.Cleanup(func() {
		t.Logf("Attempting to delete network")
		deleteOpts := network.DeleteOpts{
			Namespace: nsName,
			Name:      networkName,
		}
		_, err = network.Delete(networkClient, deleteOpts)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to list network")
	listOpts := network.ListOpts{
		Namespace: nsName,
	}
	netList, err := network.List(networkClient, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, netList[0].Spec.SecurityGroups[0], "default")

	t.Logf("Attempting to get network")
	getNetwork, err := network.Get(networkClient, nsName, networkName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getNetwork.Spec.NetworkType, "underlay_neutron")

	t.Logf("Attempting to update network")
	updateOpts := network.UpdateOpts{
		APIVersion: "yangtse/v2",
		Kind:       "Network",
		Namespace:  nsName,
		Name:       networkName,
		Metadata: &network.ObjectMeta{
			Name: networkName,
			Annotations: map[string]string{
				"yangtse.io/domain-id":                  client.DomainID,
				"yangtse.io/project-id":                 client.ProjectID,
				"yangtse.io/warm-pool-recycle-interval": "2",
				"yangtse.io/warm-pool-size":             "20",
			},
			ResourceVersion: getNetwork.Metadata.ResourceVersion,
		},
		Spec: &network.NetworkSpec{
			NetworkType: "underlay_neutron",
			SecurityGroups: []string{
				"default",
			},
			Subnets: []network.SubnetConf{
				{
					SubnetID: subnetID,
				},
			},
		},
		Status: &network.NetworkStatus{},
	}

	updatedCm, err := network.Update(networkClient, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedCm.Metadata.Annotations["yangtse.io/warm-pool-recycle-interval"], updateOpts.Metadata.Annotations["yangtse.io/warm-pool-recycle-interval"])
	th.AssertEquals(t, updatedCm.Metadata.Annotations["yangtse.io/warm-pool-size"], updateOpts.Metadata.Annotations["yangtse.io/warm-pool-size"])
}
