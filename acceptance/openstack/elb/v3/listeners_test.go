package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListenerLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	certificateID := createCertificate(t, client)
	defer deleteCertificate(t, client, certificateID)

	t.Logf("Attempting to create ELBv3 Listener")
	listenerName := tools.RandomString("create-listener-", 3)

	createOpts := listeners.CreateOpts{
		DefaultTlsContainerRef: certificateID,
		Description:            "some interesting description",
		LoadbalancerID:         loadbalancerID,
		Name:                   listenerName,
		Protocol:               "HTTPS",
		ProtocolPort:           443,
		Tags: []tags.ResourceTag{
			{
				Key:   "gophertelekomcloud",
				Value: "listener",
			},
		},
	}

	listener, err := listeners.Create(client, createOpts).Extract()
	defer func() {
		t.Logf("Attempting to delete ELBv3 Listener: %s", listener.ID)
		err := listeners.Delete(client, listener.ID).Err
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 Listener: %s", listener.ID)
	}()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Name, listener.Name)
	th.AssertEquals(t, createOpts.Description, listener.Description)
	t.Logf("Created ELBv3 Listener: %s", listener.ID)

	t.Logf("Attempting to update ELBv3 Listener: %s", listener.ID)
	listenerName = tools.RandomString("update-listener-", 3)
	emptyDescription := ""
	updateOpts := listeners.UpdateOpts{
		Description:  &emptyDescription,
		Name:         &listenerName,
		SniMatchAlgo: "longest_suffix",
	}
	_, err = listeners.Update(client, listener.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Listener: %s", listener.ID)

	newListener, err := listeners.Get(client, listener.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listenerName, newListener.Name)
	th.AssertEquals(t, emptyDescription, newListener.Description)
	th.AssertEquals(t, "longest_suffix", newListener.SniMatchAlgo)

	listOpts := listeners.ListOpts{LoadBalancerID: []string{loadbalancerID}}
	pages, err := listeners.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	listenerSlice, err := listeners.ExtractListeners(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listenerSlice))
	th.AssertDeepEquals(t, *newListener, listenerSlice[0])
}
