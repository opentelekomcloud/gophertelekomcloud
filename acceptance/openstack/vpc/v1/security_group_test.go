package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygrouprules"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSecurityGroupLifecycle(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	group, err := securitygroups.Create(client, securitygroups.CreateOpts{
		Name:                tools.RandomString("security-group-acc-", 3),
		EnterpriseProjectID: "0",
	})
	th.AssertNoErr(t, err)
	groupID := group.ID
	t.Cleanup(func() {
		if groupID != "" {
			th.AssertNoErr(t, securitygroups.Delete(client, groupID))
		}
	})

	actualGroup, err := securitygroups.Get(client, groupID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, groupID, actualGroup.ID)

	groups, err := securitygroups.List(client, securitygroups.ListOpts{EnterpriseProjectID: "0"})
	th.AssertNoErr(t, err)
	groupFound := false
	for _, candidate := range groups {
		if candidate.ID == groupID {
			groupFound = true
			break
		}
	}
	th.AssertEquals(t, true, groupFound)

	rule, err := securitygrouprules.Create(client, securitygrouprules.CreateOpts{
		SecurityGroupID: groupID,
		Description:     "acceptance rule",
		Direction:       "ingress",
		EtherType:       "IPv4",
		Protocol:        "tcp",
		PortRangeMin:    pointerto.Int(8080),
		PortRangeMax:    pointerto.Int(8080),
		RemoteIPPrefix:  "192.0.2.0/24",
	})
	th.AssertNoErr(t, err)
	ruleID := rule.ID
	t.Cleanup(func() {
		if ruleID != "" {
			th.AssertNoErr(t, securitygrouprules.Delete(client, ruleID))
		}
	})

	actualRule, err := securitygrouprules.Get(client, ruleID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ruleID, actualRule.ID)
	th.AssertEquals(t, groupID, actualRule.SecurityGroupID)

	rules, err := securitygrouprules.List(client, securitygrouprules.ListOpts{
		SecurityGroupID: groupID,
		RemoteIPPrefix:  "192.0.2.0/24",
	})
	th.AssertNoErr(t, err)
	ruleFound := false
	for _, candidate := range rules {
		if candidate.ID == ruleID {
			ruleFound = true
			break
		}
	}
	th.AssertEquals(t, true, ruleFound)

	th.AssertNoErr(t, securitygrouprules.Delete(client, ruleID))
	ruleID = ""
	th.AssertNoErr(t, securitygroups.Delete(client, groupID))
	groupID = ""
}
