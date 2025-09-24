package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v3/security/group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v3/security/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVPCSecurityGroupRuleV3Lifecycle(t *testing.T) {
	client, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("sec-grp-acc-", 3)
	secGroup, err := group.Create(client, group.CreateOpts{
		SecurityGroup: group.SecurityGroupOptions{
			Name: name,
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, secGroup.ID))
	})

	createOpts := rules.CreateOpts{
		SecurityGroupRule: rules.SecurityGroupRuleOptions{
			SecurityGroupID: secGroup.ID,
			Direction:       "ingress",
			Protocol:        "tcp",
			Description:     "create-rule",
			Action:          "allow",
			Priority:        1,
			Multiport:       "8080",
			RemoteIPPrefix:  "10.10.0.0/16",
		},
	}

	rule, err := rules.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, rules.Delete(client, rule.ID))
	})

	ruleList, err := rules.List(client, rules.ListQueryParams{
		SecurityGroupId: []string{secGroup.ID},
	})
	th.AssertNoErr(t, err)
	th.AssertNotEquals(t, 0, len(ruleList.SecurityGroupRules))

	secRule, err := rules.Get(client, rule.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "create-rule", secRule.Description)
	th.AssertEquals(t, 1, secRule.Priority)
}
