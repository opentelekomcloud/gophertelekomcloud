package security_groups

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/security/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestThrottlingSgs(t *testing.T) {
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

	sg1, err := secgroups.Create(clientCompute, createSGOpts).Extract()
	th.AssertNoErr(t, err)
	sgs1, err := CreateMultipleSgsRules(clientNetworking, sg1.ID, 100)
	th.AssertNoErr(t, err)
	if len(sgs1) == 0 {
		t.Fatalf("empty rules list for: %s", sg1.ID)
	}

	createSGOpts2 := secgroups.CreateOpts{
		Name:        "sg-test-02",
		Description: "desc",
	}
	t.Logf("Attempting to create sg: %s", createSGOpts2.Name)
	sg2, err := secgroups.Create(clientCompute, createSGOpts2).Extract()
	th.AssertNoErr(t, err)
	sgs2, err := CreateMultipleSgsRules(clientNetworking, sg2.ID, 100)
	th.AssertNoErr(t, err)
	if len(sgs2) == 0 {
		t.Fatalf("empty rules list for: %s", sg2.ID)
	}

	rulesOpts := rules.ListOpts{
		SecGroupID: sg1.ID,
	}
	allPages, err := rules.List(clientNetworking, rulesOpts).AllPages()
	th.AssertNoErr(t, err)

	rls, err := rules.ExtractRules(allPages)
	th.AssertNoErr(t, err)
	if len(rls) == 0 {
		t.Fatalf("empty rules list")
	}

	t.Cleanup(func() {
		secgroups.Delete(clientCompute, sg1.ID)
		secgroups.Delete(clientCompute, sg2.ID)
	})
}
