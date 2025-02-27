package management

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCFWList(t *testing.T) {
	clientv1, err := clients.NewCFWV1Client()
	th.AssertNoErr(t, err)

	queryOpts := managementv1.ListOpts{
		Limit: 1024,
	}
	_, err = managementv1.List(clientv1, queryOpts)
	th.AssertNoErr(t, err)
}

func TestFirewallLifecycle(t *testing.T) {
	clientv1, err := clients.NewCFWV1Client()
	th.AssertNoErr(t, err)
	clientv2, err := clients.NewCFWV2Client()
	th.AssertNoErr(t, err)
	clientv3, err := clients.NewCFWV3Client()
	th.AssertNoErr(t, err)

	instanceName := tools.RandomString("test-acc-firewall-", 3)
	createOpts := managementv2.CreateOpts{
		Name: instanceName,
		Flavor: managementv2.CreateFlavor{
			Version: "standard",
		},
		ChargeInfo: managementv2.ChargeInfo{
			ChargeMode: "postPaid",
		},
	}

	createResp, err := managementv2.Create(clientv2, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, instanceName, createResp.Data.Name)
	th.AssertNoErr(t, common.WaitForJobCompleted(clientv3, 600, 5, createResp.JobID))
	instanceId := createResp.JobID
	t.Cleanup(func() {
		_, err = managementv2.Delete(clientv2, instanceId)
		th.AssertNoErr(t, err)
	})

	firewall, err := managementv1.Get(clientv1, instanceId, 0)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, instanceName, firewall.FwInstanceName)
}
