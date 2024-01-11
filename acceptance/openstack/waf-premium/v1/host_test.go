package v1

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/hosts"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestWafPremiumHostWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	hostId := createHost(t, client, vpcID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium host: %s", hostId)
		th.AssertNoErr(t, hosts.Delete(client, hostId, hosts.DeleteOpts{}))
		t.Logf("Deleted WAF Premium host: %s", hostId)
	})

	t.Logf("Attempting to Get WAF Premium host: %s", hostId)
	h, err := hosts.Get(client, hostId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, h.ID, hostId)
	th.AssertEquals(t, h.ProtectStatus, 1)
	th.AssertEquals(t, h.Description, "description")

	t.Logf("Attempting to List WAF Premium hosts")
	hostsList, err := hosts.List(client, hosts.ListOpts{})
	th.AssertNoErr(t, err)

	if len(hostsList) < 1 {
		t.Fatal("empty WAF hosts list")
	}

	hostUpdated, err := hosts.Update(client, hostId, hosts.UpdateOpts{
		Proxy:       pointerto.Bool(true),
		Description: "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, hostUpdated.Proxy, true)
	th.AssertEquals(t, hostUpdated.Description, "updated")

	t.Logf("Attempting to Update WAF Premium host protect status: %s", hostId)
	err = hosts.UpdateProtectStatus(client, hostId, hosts.ProtectUpdateOpts{
		ProtectStatus: 0,
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get updated WAF Premium host: %s", hostId)
	hu, err := hosts.Get(client, hostId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, hu.ProtectStatus, 0)
}

func createHost(t *testing.T, client *golangsdk.ServiceClient, vpcID string) string {
	t.Logf("Attempting to create WAF premium host")

	server := hosts.PremiumWafServer{
		FrontProtocol: "HTTP",
		BackProtocol:  "HTTP",
		Address:       "10.10.11.11",
		Port:          80,
		Type:          "ipv4",
		VpcId:         vpcID,
	}
	opts := hosts.CreateOpts{
		Hostname:    tools.RandomString("www.waf-demo.com", 3),
		Proxy:       pointerto.Bool(false),
		Server:      []hosts.PremiumWafServer{server},
		Description: "description",
	}
	h, err := hosts.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created WAF host: %s", h.ID)
	return h.ID
}
