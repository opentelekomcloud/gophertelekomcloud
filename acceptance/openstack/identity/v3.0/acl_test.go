package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3.0/acl"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestACLConsoleLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	createOpts := acl.ACLPolicy{
		DomainId: client.DomainID,
		AllowAddressNetmasks: []acl.AllowAddressNetmasks{
			{
				AddressNetmask: "192.168.0.1/24",
				Description:    "test-netmask-description",
			},
		},
		AllowIPRanges: []acl.AllowIPRanges{
			{
				IPRange:     "0.0.0.0-255.255.255.255",
				Description: "test ip range description",
			},
		},
	}

	resp, err := acl.ConsoleACLPolicyUpdate(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.AllowAddressNetmasks[0].AddressNetmask, resp.AllowAddressNetmasks[0].AddressNetmask)
	th.AssertEquals(t, createOpts.AllowAddressNetmasks[0].Description, resp.AllowAddressNetmasks[0].Description)
	th.AssertEquals(t, createOpts.AllowIPRanges[0].IPRange, resp.AllowIPRanges[0].IPRange)
	th.AssertEquals(t, createOpts.AllowIPRanges[0].Description, resp.AllowIPRanges[0].Description)

	getAcl, err := acl.ConsoleACLPolicyGet(client, client.DomainID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, getAcl.AllowIPRanges[0].IPRange, resp.AllowIPRanges[0].IPRange)
	th.AssertEquals(t, getAcl.AllowIPRanges[0].Description, resp.AllowIPRanges[0].Description)

	t.Cleanup(func() {
		netmasksList := make([]acl.AllowAddressNetmasks, 0, 1)
		netmask := acl.AllowAddressNetmasks{
			AddressNetmask: "0.0.0.0-255.255.255.255",
		}
		netmasksList = append(netmasksList, netmask)

		deleteOpts := acl.ACLPolicy{
			DomainId:             client.DomainID,
			AllowAddressNetmasks: netmasksList,
		}
		_, err = acl.ConsoleACLPolicyUpdate(client, deleteOpts)
		th.AssertNoErr(t, err)
	})
}

func TestACLAPILifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	createOpts := acl.ACLPolicy{
		DomainId: client.DomainID,
		AllowAddressNetmasks: []acl.AllowAddressNetmasks{
			{
				AddressNetmask: "192.168.0.1/24",
				Description:    "test-netmask-description",
			},
		},
		AllowIPRanges: []acl.AllowIPRanges{
			{
				IPRange:     "0.0.0.0-255.255.255.255",
				Description: "test ip range description",
			},
		},
	}

	resp, err := acl.APIACLPolicyUpdate(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.AllowAddressNetmasks[0].AddressNetmask, resp.AllowAddressNetmasks[0].AddressNetmask)
	th.AssertEquals(t, createOpts.AllowAddressNetmasks[0].Description, resp.AllowAddressNetmasks[0].Description)
	th.AssertEquals(t, createOpts.AllowIPRanges[0].IPRange, resp.AllowIPRanges[0].IPRange)
	th.AssertEquals(t, createOpts.AllowIPRanges[0].Description, resp.AllowIPRanges[0].Description)

	getAcl, err := acl.APIACLPolicyGet(client, client.DomainID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, getAcl.AllowIPRanges[0].IPRange, resp.AllowIPRanges[0].IPRange)
	th.AssertEquals(t, getAcl.AllowIPRanges[0].Description, resp.AllowIPRanges[0].Description)

	t.Cleanup(func() {
		netmasksList := make([]acl.AllowAddressNetmasks, 0, 1)
		netmask := acl.AllowAddressNetmasks{
			AddressNetmask: "0.0.0.0-255.255.255.255",
		}
		netmasksList = append(netmasksList, netmask)

		deleteOpts := acl.ACLPolicy{
			DomainId:             client.DomainID,
			AllowAddressNetmasks: netmasksList,
		}
		_, err = acl.APIACLPolicyUpdate(client, deleteOpts)
		th.AssertNoErr(t, err)
	})
}
