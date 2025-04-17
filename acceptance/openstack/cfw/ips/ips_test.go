package ips

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/ips"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestIPSLifecycle(t *testing.T) {
	t.Skip("Too long. Non reproducible in CI")
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
	th.AssertNoErr(t, common.WaitForJobCompleted(clientv3, 600, 5, createResp.JobID))
	instanceId := createResp.JobID
	t.Cleanup(func() {
		_, err = managementv2.Delete(clientv2, instanceId)
		th.AssertNoErr(t, err)
	})

	firewall, err := managementv1.Get(clientv1, instanceId, "0")
	th.AssertNoErr(t, err)

	zero := 0
	th.AssertNoErr(t, ips.SetIPSFeatureStatus(clientv1, ips.SetFeatureStatusOpts{
		ObjectID: firewall.ProtectObjects[0].ObjectID,
		IpsType:  2,
		Status:   &zero,
	}))
	_, err = ips.GetIPSFeatureStatus(clientv1, firewall.ProtectObjects[0].ObjectID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, ips.SetProtectionMode(clientv1, ips.SetProtectionModeOpts{
		ObjectID: firewall.ProtectObjects[0].ObjectID,
		Mode:     &zero,
	}))
	_, err = ips.GetProtectionMode(clientv1, firewall.ProtectObjects[0].ObjectID)
	th.AssertNoErr(t, err)
}
