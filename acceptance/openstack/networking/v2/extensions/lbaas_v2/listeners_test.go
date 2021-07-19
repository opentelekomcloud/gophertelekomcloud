package lbaas_v2

import (
	"reflect"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLbaasV2ListenersLifeCycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create lbaasV2 Load Balancer
	loadBalancer := createLbaasLoadBalancer(t, client)
	defer deleteLbaasLoadBalancer(t, client, loadBalancer.ID)

	// Create lbaasV2 pool
	loadBalancerPool := createLbaasPool(t, client, loadBalancer.ID)
	defer deleteLbaasPool(t, client, loadBalancerPool.ID)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	// Get ECSv1 createOpts
	ecsCreateOpts := openstack.GetCloudServerCreateOpts(t)

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, computeClient, ecsCreateOpts)
	defer openstack.DeleteCloudServer(t, computeClient, ecs.ID)

	ecsNICs := reflect.ValueOf(ecs.Addresses).MapKeys()
	ecsPrivateIP := ecs.Addresses[ecsNICs[0].String()][0].Addr

	member := createLbaasMember(t, client, loadBalancerPool.ID, ecsPrivateIP)
	defer deleteLbaasMember(t, client, loadBalancerPool.ID, member.ID)

	// Create lbaasV2 certificate
	lbaasCertificate := createLbaasCertificate(t, client)
	defer deleteLbaasCertificate(t, client, lbaasCertificate.ID)

	// Create lbaasV2 listener
	listenerName := tools.RandomString("create-listener-", 3)
	t.Logf("Attempting to create LbaasV2 Listener")

	http2Enable := true
	createOpts := listeners.CreateOpts{
		LoadbalancerID:         loadBalancer.ID,
		Protocol:               "TERMINATED_HTTPS",
		ProtocolPort:           443,
		Name:                   listenerName,
		DefaultPoolID:          loadBalancerPool.ID,
		Http2Enable:            &http2Enable,
		DefaultTlsContainerRef: lbaasCertificate.ID,
		TlsCiphersPolicy:       "tls-1-2-strict",
	}

	listener, err := listeners.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listenerName, listener.Name)
	th.AssertEquals(t, http2Enable, listener.Http2Enable)
	t.Logf("Created LbaasV2 listener: %s", listener.ID)
	defer func() {
		updateOpts := listeners.UpdateOpts{
			DefaultPoolID: "null",
		}
		_, err = listeners.Update(client, listener.ID, updateOpts).Extract()
		th.AssertNoErr(t, listeners.Delete(client, listener.ID).ExtractErr())
	}()
}
