package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
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

	listener, err := listeners.Create(client, createOpts)
	defer func() {
		t.Logf("Attempting to delete ELBv3 Listener: %s", listener.ID)
		err := listeners.Delete(client, listener.ID)
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
	_, err = listeners.Update(client, listener.ID, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Listener: %s", listener.ID)

	newListener, err := listeners.Get(client, listener.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listenerName, newListener.Name)
	th.AssertEquals(t, emptyDescription, newListener.Description)
	th.AssertEquals(t, "longest_suffix", newListener.SniMatchAlgo)

	listOpts := listeners.ListOpts{LoadBalancerID: []string{loadbalancerID}}
	listenerSlice, err := listeners.List(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listenerSlice))
	th.AssertDeepEquals(t, *newListener, listenerSlice[0])
}

func TestListenerForceDelete(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	listenerID := createListener(t, client, loadbalancerID)
	listenerGone := false
	poolGone := false
	poolID := ""
	defer func() {
		if !poolGone && poolID != "" {
			deletePool(t, client, poolID)
		}
		if !listenerGone {
			deleteListener(t, client, listenerID)
		}
	}()

	pool, err := pools.Create(client, pools.CreateOpts{
		LBMethod:                 "LEAST_CONNECTIONS",
		Protocol:                 "HTTP",
		ListenerID:               listenerID,
		Name:                     tools.RandomString("force-delete-pool-", 3),
		VpcId:                    clients.EnvOS.GetEnv("VPC_ID"),
		Type:                     "instance",
		DeletionProtectionEnable: pointerto.Bool(false),
	}).Extract()
	th.AssertNoErr(t, err)
	poolID = pool.ID

	err = listeners.ForceDelete(client, listenerID)
	th.AssertNoErr(t, err)

	if _, err = listeners.Get(client, listenerID); err == nil {
		t.Fatal("expected force-deleted listener to be absent")
	}
	listenerGone = true
	if _, err = pools.Get(client, poolID).Extract(); err == nil {
		t.Fatal("expected associated pool to be deleted")
	}
	poolGone = true
}
