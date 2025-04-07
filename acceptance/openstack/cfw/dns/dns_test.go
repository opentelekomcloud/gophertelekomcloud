package dns

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/dns"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCFWDNSLifecycle(t *testing.T) {
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

	firewall, err := managementv1.Get(clientv1, instanceId, 0)
	th.AssertNoErr(t, err)

	groupName := tools.RandomString("test-acc-dns-group-", 3)
	createDGResp, err := dns.CreateDomainNameGroup(clientv1, dns.CreateDomainNameGroupOpts{
		FwInstanceID: instanceId,
		ObjectID:     firewall.ProtectObjects[0].ObjectID,
		Name:         groupName,
		DomainNames: []dns.DomainSetInfoDto{
			{
				DomainName: "www.aaa.com",
			},
		},
	})
	th.AssertNoErr(t, err)

	groupId := createDGResp.Id

	dns.AddDomainNames(clientv1, groupId, dns.AddDomainNameListOpts{
		FwInstanceID: instanceId,
		ObjectID:     firewall.ProtectObjects[0].ObjectID,
		DomainNames: []dns.DomainSetInfoDto{
			{
				DomainName:  "www.bbb.com",
				Description: "Test Domain",
			},
		},
	})
	th.AssertNoErr(t, err)

	domainInfo, err := dns.GetDomainNameInfo(clientv1, "www.bbb.com", groupId, instanceId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "Test Domain", domainInfo.Description)

	th.AssertNoErr(t, dns.DeleteDomainNames(clientv1, groupId, instanceId, dns.DeleteDomainNameListOpts{
		ObjectID:         firewall.ProtectObjects[0].ObjectID,
		DomainAddressIDs: []string{domainInfo.DomainAddressID},
	}))

	_, err = dns.ListDomainNameGroups(clientv1, instanceId, firewall.ProtectObjects[0].ObjectID)
	th.AssertNoErr(t, err)

	_, err = dns.UpdateDomainNameGroup(clientv1, groupId, instanceId, dns.UpdateOpts{
		Name:        groupName,
		Description: "Test Group",
	})
	th.AssertNoErr(t, err)

	groupInfo, err := dns.GetDomainNameGroup(clientv1, groupName, instanceId, firewall.ProtectObjects[0].ObjectID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, groupName, groupInfo.Name)
	th.AssertEquals(t, "Test Group", groupInfo.Description)

	th.AssertNoErr(t, dns.DeleteDomainNameGroup(clientv1, groupId, instanceId))
}
