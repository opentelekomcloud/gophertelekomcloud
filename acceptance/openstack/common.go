// Package openstack contains common functions that can be used
// across all OpenStack components for acceptance testing.
package openstack

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/extensions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// PrintExtension prints an extension and all of its attributes.
func PrintExtension(t *testing.T, extension *extensions.Extension) {
	t.Logf("Name: %s", extension.Name)
	t.Logf("Namespace: %s", extension.Namespace)
	t.Logf("Alias: %s", extension.Alias)
	t.Logf("Description: %s", extension.Description)
	t.Logf("Updated: %s", extension.Updated)
	t.Logf("Links: %v", extension.Links)
}

func DefaultSecurityGroup(t *testing.T) string {
	computeClient, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	securityGroupPages, err := secgroups.List(computeClient).AllPages()
	th.AssertNoErr(t, err)
	securityGroups, err := secgroups.ExtractSecurityGroups(securityGroupPages)
	th.AssertNoErr(t, err)
	var sgId string
	for _, val := range securityGroups {
		if val.Name == "default" {
			sgId = val.ID
			break
		}
	}
	if sgId == "" {
		t.Fatalf("Unable to find default secgroup")
	}
	return sgId
}
