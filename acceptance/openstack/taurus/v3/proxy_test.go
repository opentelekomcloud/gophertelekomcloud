package v3

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/proxy"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTaurusProxyLifecycle(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	createResp := createTaurusInstance(t, client, vpcID, subnetID, secGroupID)
	instanceID := createResp.Instance.Id

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db")
		_, err = instance.Delete(client, instanceID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Waiting for instance to become available")
	err = waitForInstanceAvailable(client, 1200, instanceID)
	th.AssertNoErr(t, err)

	t.Logf("Testing GetFlavors")
	flavors, err := proxy.GetFlavors(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, flavors[0].ProxyFlavors[0].Vcpus, "2")
	th.AssertEquals(t, flavors[0].ProxyFlavors[0].Ram, "4")
	th.AssertEquals(t, flavors[0].ProxyFlavors[0].DbType, "Proxy")

	t.Logf("Testing EnableProxy")
	enableOpts := proxy.EnableProxyOpts{
		FlavorRef: "gaussdb.proxy.large.x86.2",
		NodeNum:   2,
		ProxyName: "test-proxy",
		ProxyMode: "readwrite",
	}
	jobID, err := proxy.EnableProxy(client, instanceID, enableOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for enable proxy job to complete")
	// job query always shows status as "running" with 0 progress
	th.AssertNoErr(t, waitForProxyAvailable(client, 600, instanceID))

	t.Cleanup(func() {
		t.Logf("Attempting to disable taurus db proxy")
		jobID, err = proxy.DisableProxy(client, instanceID, nil)
		th.AssertNoErr(t, err)

		t.Logf("Waiting for disable proxy job to complete")
		th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, *jobID))
	})

	t.Logf("Testing List proxies")
	proxies, err := proxy.List(client, instanceID, proxy.ListOpts{
		Offset: 0,
		Limit:  10,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(proxies) > 0, true)

	proxyID := proxies[0].Proxy.PoolId

	t.Logf("Testing UpdateWeight")
	updateWeightOpts := proxy.UpdateWeightOpts{
		MasterWeight: pointerto.Int(100),
	}
	jobID, err = proxy.UpdateWeight(client, instanceID, proxyID, updateWeightOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for update weight job to complete")
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 300, *jobID))

	t.Logf("Testing Resize")
	jobID, err = proxy.Resize(client, instanceID, proxyID, "gaussdb.proxy.xlarge.x86.2")
	th.AssertNoErr(t, err)
	t.Logf("Waiting for resize job to complete")
	th.AssertNoErr(t, waitForProxyAvailable(client, 1200, instanceID))

	t.Logf("Testing Enlarge")
	enlargeOpts := proxy.EnlargeOpts{
		InstanceID: instanceID,
		NodeNum:    1,
		ProxyId:    proxyID,
	}
	_, err = proxy.Enlarge(client, enlargeOpts)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, waitForProxyAvailable(client, 1200, instanceID))
}

func waitForProxyAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		proxies, err := proxy.List(client, instanceID, proxy.ListOpts{
			Offset: 0,
			Limit:  10,
		})
		if err != nil {
			return false, err
		}
		if proxies[0].Proxy.Status == "ACTIVE" && proxies[0].Proxy.PoolStatus == "ACTIVE" {
			return true, nil
		}
		return false, nil
	})
}
