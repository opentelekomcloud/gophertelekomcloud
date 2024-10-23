package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	v3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gaussdb/v3/proxy"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGaussDBProxyLifecycle(t *testing.T) {
	if os.Getenv("RUN_GAUSS") == "" {
		t.Skip("long test")
	}
	client, err := clients.NewGaussDBClient()
	th.AssertNoErr(t, err)
	createOpts := openstack.GetCloudServerCreateOpts(t)

	name := tools.RandomString("gaussdb-test-", 5)

	ins, err := instance.CreateInstance(client, instance.CreateInstanceOpts{
		Region:                 "eu-de",
		Name:                   name,
		AvailabilityZoneMode:   "multi",
		MasterAvailabilityZone: "eu-de-01",
		Datastore: instance.Datastore{
			Type:    "gaussdb-mysql",
			Version: "8.0",
		},
		Mode:            "Cluster",
		FlavorRef:       "gaussdb.mysql.large.x86.4",
		VpcId:           createOpts.VpcId,
		SubnetId:        createOpts.Nics[0].SubnetId,
		Password:        "gaussdb1!-test",
		SlaveCount:      pointerto.Int(1),
		SecurityGroupId: openstack.DefaultSecurityGroup(t),
	})
	th.AssertNoErr(t, err)
	id := ins.Instance.Id

	t.Cleanup(func() {
		_, err = instance.DeleteInstance(client, ins.Instance.Id)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, ins.JobId, 600)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list GaussDB proxy flavors")
	flavors, err := proxy.ListFlavors(client, id)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, flavors)

	t.Logf("Attempting to enable GaussDB proxy")
	jobId, err := proxy.EnableProxy(client, proxy.EnableProxyOpts{
		InstanceID: id,
		NodeNum:    1,
		FlavorRef:  flavors[0].ProxyFlavors[0].SpecCode,
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, jobId, 600)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list proxy instances")

	proxyNode, err := proxy.ListProxyInstances(client, proxy.ListProxyInstancesOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, proxyNode)

	t.Logf("Attempting to enlarge proxy")

	jobId, err = proxy.EnlargeProxy(client, proxy.EnlargeProxyOpts{
		InstanceID: id,
		NodeNum:    2,
	})
	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, jobId, 600)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to change proxy flavor")

	jobId, err = proxy.ChangeFlavor(client, proxy.ChangeFlavorOpts{
		InstanceID: id,
		ProxyID:    proxyNode.ProxyList[0].Proxy.Nodes[0].ID,
		FlavorRef:  flavors[0].ProxyFlavors[1].SpecCode,
	})

	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, jobId, 600)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to disable proxy flavor")
	jobId, err = proxy.Disable(client, proxy.DisableProxyOpts{
		InstanceID: id,
		ProxyIDs: &[]string{
			proxyNode.ProxyList[0].Proxy.Nodes[0].ID,
			proxyNode.ProxyList[0].Proxy.Nodes[1].ID,
		},
	})

	th.AssertNoErr(t, err)
	_, err = v3.WaitForGaussJob(client, jobId, 600)
	th.AssertNoErr(t, err)
}
