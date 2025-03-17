package blackwhitelist

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/blackwhitelist"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBlacklistWhitelistRuleLifecycle(t *testing.T) {
	// t.Skip("Too long. Non reproducible in CI")
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

	firewall, err := managementv1.Get(clientv1, instanceId, 0)
	th.AssertNoErr(t, err)

	zero := 0
	createWhitelistRuleOpts := blackwhitelist.CreateOpts{
		ObjectID:    firewall.ProtectObjects[0].ObjectID,
		ListType:    5,
		Direction:   &zero,
		AddressType: &zero,
		Address:     "1.1.1.1",
		Protocol:    6,
		Port:        "1",
	}

	_, err = blackwhitelist.CreateBlacklistOrWhitelistRule(clientv1, createWhitelistRuleOpts)
	th.AssertNoErr(t, err)

	rule, err := blackwhitelist.GetBlacklistOrWhitelistRule(clientv1, firewall.ProtectObjects[0].ObjectID, 5, "1.1.1.1")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "1.1.1.1", rule.Address)

	updateOpts := blackwhitelist.UpdateOpts{
		Address:     "1.1.1.1",
		Port:        "8500",
		Description: "test",
	}
	_, err = blackwhitelist.UpdateBlacklistOrWhitelistRule(clientv1, rule.ListId, updateOpts)
	th.AssertNoErr(t, err)

	updatedRule, err := blackwhitelist.GetBlacklistOrWhitelistRule(clientv1, firewall.ProtectObjects[0].ObjectID, 5, "1.1.1.1")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "8500", updatedRule.Port)
	th.AssertEquals(t, "test", updatedRule.Description)

	err = blackwhitelist.DeleteBlacklistOrWhitelistRule(clientv1, rule.ListId)
	th.AssertNoErr(t, err)
}
