package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/defsecrules"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const ruleID = "{ruleID}"

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	mockListRulesResponse(t)

	actual, err := defsecrules.List(client.ServiceClient())
	th.AssertNoErr(t, err)

	expected := []defsecrules.DefaultRule{
		{
			FromPort:   80,
			ID:         ruleID,
			IPProtocol: "TCP",
			IPRange:    secgroups.IPRange{CIDR: "10.10.10.0/24"},
			ToPort:     80,
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockCreateRuleResponse(t)

	opts := defsecrules.CreateOpts{
		IPProtocol: "TCP",
		FromPort:   80,
		ToPort:     80,
		CIDR:       "10.10.12.0/24",
	}

	group, err := defsecrules.Create(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)

	expected := &defsecrules.DefaultRule{
		ID:         ruleID,
		FromPort:   80,
		ToPort:     80,
		IPProtocol: "TCP",
		IPRange:    secgroups.IPRange{CIDR: "10.10.12.0/24"},
	}
	th.AssertDeepEquals(t, expected, group)
}

func TestCreateICMPZero(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockCreateRuleResponseICMPZero(t)

	opts := defsecrules.CreateOpts{
		IPProtocol: "ICMP",
		FromPort:   0,
		ToPort:     0,
		CIDR:       "10.10.12.0/24",
	}

	group, err := defsecrules.Create(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)

	expected := &defsecrules.DefaultRule{
		ID:         ruleID,
		FromPort:   0,
		ToPort:     0,
		IPProtocol: "ICMP",
		IPRange:    secgroups.IPRange{CIDR: "10.10.12.0/24"},
	}
	th.AssertDeepEquals(t, expected, group)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockGetRuleResponse(t, ruleID)

	group, err := defsecrules.Get(client.ServiceClient(), ruleID)
	th.AssertNoErr(t, err)

	expected := &defsecrules.DefaultRule{
		ID:         ruleID,
		FromPort:   80,
		ToPort:     80,
		IPProtocol: "TCP",
		IPRange:    secgroups.IPRange{CIDR: "10.10.12.0/24"},
	}

	th.AssertDeepEquals(t, expected, group)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockDeleteRuleResponse(t, ruleID)

	err := defsecrules.Delete(client.ServiceClient(), ruleID)
	th.AssertNoErr(t, err)
}
