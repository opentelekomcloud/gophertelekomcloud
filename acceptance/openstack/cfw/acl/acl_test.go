package acl

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/acl"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestACLRuleLifecycle(t *testing.T) {
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
	one := 1
	ruleName := "test-acc-firewall-rule"
	createACLRuleOpts := acl.CreateACLRuleOpts{
		ObjectID: firewall.ProtectObjects[0].ObjectID,
		Type:     &zero,
		Rules: []acl.Rule{
			{
				Name: ruleName,
				Sequence: acl.OrderRuleAclDto{
					Top: &one,
				},
				AddressType:       &zero,
				ActionType:        &zero,
				Status:            &one,
				LongConnectEnable: &zero,
				Direction:         &zero,
				Source: acl.RuleAddressDtoRequest{
					Type:    &zero,
					Address: "1.1.1.1",
				},
				Destination: acl.RuleAddressDtoRequest{
					Type:    &zero,
					Address: "2.2.2.1",
				},
				Service: acl.RuleServiceDto{
					Type:     &zero,
					Protocol: -1,
				},
			},
		},
	}

	ruleList, err := acl.CreateACLRule(clientv1, createACLRuleOpts)
	th.AssertNoErr(t, err)

	var ruleId string
	for _, rule := range ruleList {
		if rule.Name == ruleName {
			ruleId = rule.ID
		}
	}

	rule, err := acl.GetACLRule(clientv1, firewall.ProtectObjects[0].ObjectID, ruleName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ruleId, rule.RuleId)

	updateOpts := acl.UpdateACLRuleOpts{
		Name: "test-acc-firewall-rule-updated",
		Source: &acl.RuleAddressDtoRequest{
			Type:    &zero,
			Address: "1.1.1.1",
		},
		Destination: &acl.RuleAddressDtoRequest{
			Type:    &zero,
			Address: "2.2.2.2",
		},
	}
	_, err = acl.UpdateACLRule(clientv1, ruleId, updateOpts)
	th.AssertNoErr(t, err)

	updatedRule, err := acl.GetACLRule(clientv1, firewall.ProtectObjects[0].ObjectID, "test-acc-firewall-rule-updated")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "test-acc-firewall-rule-updated", updatedRule.Name)

	err = acl.DeleteACLRule(clientv1, ruleId)
	th.AssertNoErr(t, err)
}
