package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/compliance"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAllPoliciesList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	// Test ListAllPolicies
	policies, err := compliance.ListAllPolicies(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(policies) > 0)

	policy, err := compliance.GetPolicy(client, policies[0].ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.PolicyType, "builtin")
}

func TestComplianceList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	ruleName := tools.RandomString("rule-", 4)
	rule := createRule(t, client, ruleName)

	// Test ListAllCompliance
	listOpts := compliance.ListAllComplianceOpts{
		DomainId: client.DomainID,
		PolicyId: rule.ID,
	}

	allCompliance, err := compliance.ListAllRuleCompliance(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(allCompliance) > 0)

	// Test ListAllUserCompliance
	userComplianceOpts := compliance.ListAllUserComplianceOpts{
		DomainId: client.DomainID,
	}

	userCompliance, err := compliance.ListAllUserCompliance(client, userComplianceOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(userCompliance) > 0)

	// Test ListResCompliance
	resComplianceOpts := compliance.ListResComplianceOpts{
		DomainId:   client.DomainID,
		ResourceId: "test_resource_id",
	}

	resCompliance, err := compliance.ListResCompliance(client, resComplianceOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(resCompliance) >= 0)
}

func TestComplianceLifecycle(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	ruleName := tools.RandomString("rule-", 4)
	rule := createRule(t, client, ruleName)

	// Test GetRule
	retrievedRule, err := compliance.GetRule(client, client.DomainID, rule.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, rule.ID, retrievedRule.ID)

	// Test UpdateRule
	updateOpts := compliance.UpdateRuleOpts{
		DomainId:             client.DomainID,
		PolicyAssignmentId:   rule.ID,
		PolicyDefinitionID:   "5f8d549bffeecc14f1fb522a",
		Name:                 ruleName + "-new",
		PolicyAssignmentType: "builtin",
		Parameters: map[string]compliance.PolicyParameter{
			"listOfAllowedFlavors": {
				Value: []string{"second"},
			},
		},
	}

	updatedRule, err := compliance.UpdateRule(client, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ruleName+"-new", updatedRule.Name)

	th.AssertNoErr(t, waitForRuleAvailable(client, 20, client.DomainID, rule.ID))

	// Test DisableRule
	err = compliance.DisableRule(client, client.DomainID, rule.ID)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForRuleAvailable(client, 20, client.DomainID, rule.ID))

	// Test EnableRule
	err = compliance.EnableRule(client, client.DomainID, rule.ID)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForRuleAvailable(client, 20, client.DomainID, rule.ID))

	// Test RunEval
	err = compliance.RunEval(client, client.DomainID, rule.ID)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForRuleAvailable(client, 20, client.DomainID, rule.ID))

	// Test ListRules
	listRulesOpts := compliance.ListRulesOpts{
		DomainId: client.DomainID,
	}

	rules, err := compliance.ListRules(client, listRulesOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(rules) > 0)

	// Test GetPolicy
	// policy, err := compliance.GetPolicy(client, rule.PolicyDefinitionID)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, rule.PolicyDefinitionID, policy.ID)
	//
	// // Test UpdateCompliance
	// updateComplianceOpts := compliance.UpdateComplianceOpts{
	// 	DomainId: client.DomainID,
	// 	PolicyResource: compliance.PolicyResource{
	// 		ResourceID:   "test_resource_id",
	// 		ResourceName: "TestResource",
	// 	},
	// 	TriggerType:          "resource",
	// 	ComplianceState:      "NonCompliant",
	// 	PolicyAssignmentID:   rule.ID,
	// 	PolicyAssignmentName: rule.Name,
	// 	EvaluationTime:       1667374060248,
	// 	EvaluationHash:       "89342b8f338165651991afb8bd471396",
	// }
	//
	// updatedCompliance, err := compliance.UpdateCompliance(client, updateComplianceOpts)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, "Compliant", updatedCompliance.ComplianceState)
}

func createRule(t *testing.T, client *golangsdk.ServiceClient, name string) *compliance.PolicyRule {
	addOpts := compliance.AddRuleOpts{
		Name:                 name,
		DomainId:             client.DomainID,
		PolicyAssignmentType: "builtin",
		PolicyDefinitionID:   "5f8d549bffeecc14f1fb522a",
		Description:          "Test compliance rule",
		Parameters: map[string]compliance.PolicyParameter{
			"listOfAllowedFlavors": {
				Value: []string{"first"},
			},
		},
	}

	rule, err := compliance.AddRule(client, addOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, rule.Name)

	t.Cleanup(func() {
		th.AssertNoErr(t, compliance.DisableRule(client, client.DomainID, rule.ID))
		th.AssertNoErr(t, compliance.Delete(client, client.DomainID, rule.ID))
	})

	return rule
}

func waitForRuleAvailable(client *golangsdk.ServiceClient, secs int, domainId, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		jobStatus, err := compliance.GetRuleStatus(client, domainId, id)
		if err != nil {
			return false, err
		}
		if jobStatus == nil {
			return false, nil
		}
		if jobStatus.State == "Succeeded" {
			return true, nil
		}
		return false, nil
	})
}
