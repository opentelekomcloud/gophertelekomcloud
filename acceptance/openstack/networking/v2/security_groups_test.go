package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestThrottlingSgs(t *testing.T) {
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

	createSGOpts2 := secgroups.CreateOpts{
		Name:        "sg-test-02",
		Description: "desc",
	}
	t.Logf("Attempting to create sg: %s", createSGOpts2.Name)
	sg2, err := secgroups.Create(clientCompute, createSGOpts2).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		secgroups.Delete(clientCompute, sg1.ID)
		secgroups.Delete(clientCompute, sg2.ID)
	})
}
