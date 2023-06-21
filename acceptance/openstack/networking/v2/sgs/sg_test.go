package sgs

import (
	"log"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/security/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestICMPSecurityGroupRules(t *testing.T) {
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

	optsEchoReply := rules.CreateOpts{
		Description:  "ICMP echo reply",
		SecGroupID:   sg.ID,
		PortRangeMin: pointerto.Int(0),
		PortRangeMax: pointerto.Int(0),
		Direction:    "ingress",
		EtherType:    "IPv4",
		Protocol:     "ICMP",
	}
	log.Print("[DEBUG] Create OpenTelekomCloud Neutron ICMP echo reply Security Group Rule")
	echoReply, err := rules.Create(clientNetworking, optsEchoReply).Extract()

	getEchoReply, err := rules.Get(clientNetworking, echoReply.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, *getEchoReply.PortRangeMin, 0)
	th.AssertEquals(t, *getEchoReply.PortRangeMax, 0)

	optsAll := rules.CreateOpts{
		Description:  "ICMP all",
		SecGroupID:   sg.ID,
		PortRangeMin: nil,
		PortRangeMax: nil,
		Direction:    "ingress",
		EtherType:    "IPv4",
		Protocol:     "ICMP",
	}
	log.Print("[DEBUG] Create OpenTelekomCloud Neutron ICMP All Security Group Rule")
	all, err := rules.Create(clientNetworking, optsAll).Extract()

	getAll, err := rules.Get(clientNetworking, all.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getAll.PortRangeMin, (*int)(nil))
	th.AssertEquals(t, getAll.PortRangeMax, (*int)(nil))

	optsEcho := rules.CreateOpts{
		Description:  "ICMP echo",
		SecGroupID:   sg.ID,
		PortRangeMin: pointerto.Int(8),
		PortRangeMax: pointerto.Int(0),
		Direction:    "ingress",
		EtherType:    "IPv4",
		Protocol:     "ICMP",
	}
	log.Print("[DEBUG] Create OpenTelekomCloud Neutron ICMP Echo Security Group Rule")
	echo, err := rules.Create(clientNetworking, optsEcho).Extract()

	getEcho, err := rules.Get(clientNetworking, echo.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, *getEcho.PortRangeMin, 8)
	th.AssertEquals(t, *getEcho.PortRangeMax, 0)

	optsFragment := rules.CreateOpts{
		Description:  "ICMP Fragment need DF set",
		SecGroupID:   sg.ID,
		PortRangeMin: pointerto.Int(3),
		PortRangeMax: pointerto.Int(4),
		Direction:    "ingress",
		EtherType:    "IPv4",
		Protocol:     "ICMP",
	}
	log.Print("[DEBUG] Create OpenTelekomCloud Neutron ICMP Fragment need DF set Security Group Rule")
	fragment, err := rules.Create(clientNetworking, optsFragment).Extract()

	getFragment, err := rules.Get(clientNetworking, fragment.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, *getFragment.PortRangeMin, 3)
	th.AssertEquals(t, *getFragment.PortRangeMax, 4)

	t.Cleanup(func() {
		secgroups.Delete(clientCompute, sg.ID)
	})
}

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
		portRangeMin := startIndex*1000 + i
		portRangeMax := startIndex*5000 + i
		opts := rules.CreateOpts{
			Description:  "description",
			SecGroupID:   sgID,
			PortRangeMin: &portRangeMin,
			PortRangeMax: &portRangeMax,
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
