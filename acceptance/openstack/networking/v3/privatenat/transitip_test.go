package privatenat

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v3/privatenat/transitip"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPrivateNatTransitIpLifecycle(t *testing.T) {
	client, err := clients.NewNatV3Client()
	th.AssertNoErr(t, err)
	networkId := clients.EnvOS.GetEnv("NETWORK_ID")
	if networkId == "" {
		t.Skip("OS_NETWORK_ID is missing but test requires using existing network")
	}

	createOpts := transitip.CreateTransitIpOpts{
		VirSubnetID: networkId,
	}
	createResp, err := transitip.Create(client, createOpts)
	th.AssertNoErr(t, err)
	transitIpId := createResp.TransitIp.Id
	t.Logf("Created Transit IP: %s", transitIpId)
	t.Cleanup(func() {
		t.Logf("Deleting Transit IP: %s", transitIpId)
		th.AssertNoErr(t, transitip.Delete(client, transitIpId))
		t.Logf("Deleted Transit IP: %s", transitIpId)
	})

	getResponse, err := transitip.Get(client, transitIpId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, networkId, getResponse.TransitIp.VirSubnetID)

	listResonse, err := transitip.List(client, transitip.ListTransitIpsQueryParams{
		Id: []string{transitIpId},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listResonse.TransitIps))
}
