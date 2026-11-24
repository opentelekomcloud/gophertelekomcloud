package dns

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/endpoints"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEndpointsLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but DNS endpoint test requires using existing network")
	}

	client, err := clients.NewDNSV2Client()
	th.AssertNoErr(t, err)

	ep := createEndpoint(t, client, "inbound", subnetID)

	got, err := endpoints.Get(client, ep.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)
	th.AssertEquals(t, "inbound", got.Direction)
	th.AssertEquals(t, vpcID, got.VpcID)
	th.AssertEquals(t, 2, got.IPAddressCount)

	list, err := endpoints.List(client, endpoints.ListOpts{Direction: "inbound", Name: ep.Name})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(list.Endpoints))
	th.AssertEquals(t, ep.ID, list.Endpoints[0].ID)

	newName := ep.Name + "-upd"
	updated, err := endpoints.Update(client, ep.ID, endpoints.UpdateOpts{Name: newName})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, newName, updated.Name)

	before, err := waitForEndpointIPs(client, ep.ID, 2)
	th.AssertNoErr(t, err)

	_, err = endpoints.AddIP(client, ep.ID, endpoints.IPAddressInfo{SubnetID: subnetID})
	th.AssertNoErr(t, err)
	after, err := waitForEndpointIPs(client, ep.ID, 3)
	th.AssertNoErr(t, err)

	known := map[string]bool{}
	for _, ip := range before {
		known[ip.ID] = true
	}
	var addedID string
	for _, ip := range after {
		if !known[ip.ID] {
			addedID = ip.ID
		}
	}
	th.AssertEquals(t, true, addedID != "")

	th.AssertNoErr(t, endpoints.RemoveIP(client, ep.ID, addedID))
	_, err = waitForEndpointIPs(client, ep.ID, 2)
	th.AssertNoErr(t, err)

	vpcs, err := endpoints.ListVPCs(client, endpoints.ListVPCsOpts{VpcID: vpcID})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, vpcs)
	th.AssertEquals(t, 1, len(vpcs.VPCs))
	th.AssertEquals(t, vpcID, vpcs.VPCs[0].ID)
	th.AssertEquals(t, true, vpcs.VPCs[0].InboundEndpointCount >= 1)
}

// createEndpoint creates an endpoint with two IP addresses in the subnet, waits for ACTIVE
// and registers deletion in t.Cleanup.
func createEndpoint(t *testing.T, client *golangsdk.ServiceClient, direction, subnetID string) *endpoints.Endpoint {
	cloud, err := clients.EnvOS.Cloud()
	th.AssertNoErr(t, err)

	opts := endpoints.CreateOpts{
		Name:        tools.RandomString("acc-dns-ep-", 3),
		Direction:   direction,
		Region:      cloud.RegionName,
		IPAddresses: []endpoints.IPAddressInfo{{SubnetID: subnetID}, {SubnetID: subnetID}},
	}
	ep, err := endpoints.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created DNS %s endpoint: %s", direction, ep.ID)

	t.Cleanup(func() {
		th.AssertNoErr(t, endpoints.Delete(client, ep.ID))
		th.AssertNoErr(t, waitForDeleted(func() error {
			_, err := endpoints.Get(client, ep.ID)
			return err
		}))
		t.Logf("Deleted DNS endpoint: %s", ep.ID)
	})

	th.AssertEquals(t, opts.Name, ep.Name)
	th.AssertEquals(t, direction, ep.Direction)

	err = golangsdk.WaitFor(600, func() (bool, error) {
		cur, err := endpoints.Get(client, ep.ID)
		if err != nil {
			return false, err
		}
		return cur.Status == "ACTIVE", nil
	})
	th.AssertNoErr(t, err)
	return ep
}

// waitForEndpointIPs waits until the endpoint has exactly count IP addresses, all ACTIVE.
func waitForEndpointIPs(client *golangsdk.ServiceClient, endpointID string, count int) ([]endpoints.IPAddress, error) {
	var ips []endpoints.IPAddress
	err := golangsdk.WaitFor(600, func() (bool, error) {
		var err error
		ips, err = endpoints.ListIPs(client, endpointID, endpoints.ListIPsOpts{})
		if err != nil || len(ips) != count {
			return false, err
		}
		for _, ip := range ips {
			if ip.Status != "ACTIVE" {
				return false, nil
			}
		}
		return true, nil
	})
	return ips, err
}

// waitForDeleted waits until get returns 404.
func waitForDeleted(get func() error) error {
	return golangsdk.WaitFor(600, func() (bool, error) {
		err := get()
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			return true, nil
		}
		return false, err
	})
}
