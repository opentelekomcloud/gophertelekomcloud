package sgs

import (
	"log"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/security/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestThrottlingSgs(t *testing.T) {
	t.Skip("please run only manually, long test")
	clientNetworking, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a networking client: %v", err)
	}
	clientCompute, err := clients.NewComputeV2Client()
	if err != nil {
		t.Fatalf("Unable to create a networking client: %v", err)
	}

	createSGOpts := secgroups.CreateOpts{
		Name:        "sg-test-01",
		Description: "desc",
	}
	t.Logf("Attempting to create sg: %s", createSGOpts.Name)

	sg, err := secgroups.Create(clientCompute, createSGOpts).Extract()
	th.AssertNoErr(t, err)

	size := 15
	q := make(chan []string, size)
	for i := 0; i < size; i++ {
		go CreateMultipleSgsRules(clientNetworking, sg.ID, 47, i, q) // nolint
	}
	for i := 0; i < size; i++ {
		sgs := <-q
		t.Log(sgs)
	}

	rulesOpts := rules.ListOpts{
		SecGroupID: sg.ID,
	}
	allPages, err := rules.List(clientNetworking, rulesOpts).AllPages()
	th.AssertNoErr(t, err)

	rls, err := rules.ExtractRules(allPages)
	th.AssertNoErr(t, err)
	if len(rls) == 0 {
		t.Fatalf("empty rules list")
	}

	t.Cleanup(func() {
		secgroups.Delete(clientCompute, sg.ID)
	})
}

func CreateMultipleSgsRules(clientV2 *golangsdk.ServiceClient, sgID string, count int, startIndex int, output chan<- []string) error {
	i := 0
	createdSgs := make([]string, count)
	for i < count {
		opts := rules.CreateOpts{
			Description:  "description",
			SecGroupID:   sgID,
			PortRangeMin: startIndex*1000 + i,
			PortRangeMax: startIndex*5000 + i,
			Direction:    "ingress",
			EtherType:    "IPv4",
			Protocol:     "TCP",
		}
		log.Printf("[DEBUG] Create OpenTelekomCloud Neutron security group: %#v", opts)
		securityGroupRule, err := rules.Create(clientV2, opts).Extract()
		if err != nil {
			output <- createdSgs
			return err
		}
		createdSgs[i] = securityGroupRule.ID
		i += 1
	}
	output <- createdSgs
	return nil
}
