package v1

import (
	"reflect"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

var emptyFlavor = flavors.Flavor{}

func TestFlavorList(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	versions, err := flavors.List(client)
	th.AssertNoErr(t, err)

	if len(versions) == 0 {
		t.Fatalf("empty version list")
	}
	for _, version := range versions {
		t.Logf("%+v", version)
		if version.Version == "" {
			t.Error("version object has no object")
		}
		if version.Type == "" {
			t.Error("version object has no type")
		}
		if len(version.Flavors) == 0 {
			t.Errorf("empty flavor list in version")
		}
		for _, flavor := range version.Flavors {
			if reflect.DeepEqual(emptyFlavor, flavor) {
				t.Error("flavor is empty")
			}
		}
	}
	expectedVersion := versions[0]
	expectedFlavor := expectedVersion.Flavors[0]

	filterOpts := flavors.FilterOpts{
		Version: expectedVersion.Version,
		Type:    expectedVersion.Type,
		DiskMin: &flavors.Limit{Min: expectedFlavor.DiskMin, Max: expectedFlavor.DiskMin},
		DiskMax: &flavors.Limit{Min: expectedFlavor.DiskMax, Max: expectedFlavor.DiskMax},
		Region:  expectedFlavor.Region,
		CPU:     &flavors.Limit{Min: expectedFlavor.CPU, Max: expectedFlavor.CPU},
		RAM:     &flavors.Limit{Min: expectedFlavor.RAM, Max: expectedFlavor.RAM},
	}

	filteredVersions := flavors.FilterVersions(versions, filterOpts)

	th.AssertEquals(t, 1, len(filteredVersions))
	th.AssertEquals(t, 1, len(filteredVersions[0].Flavors))
	th.AssertEquals(t, expectedVersion.Type, filteredVersions[0].Type)
	th.AssertEquals(t, expectedVersion.Version, filteredVersions[0].Version)
	th.AssertDeepEquals(t, expectedFlavor, filteredVersions[0].Flavors[0])

	foundFlavor := flavors.FindFlavor(versions, filterOpts)
	th.AssertDeepEquals(t, expectedFlavor, *foundFlavor)

	foundByName := flavors.FindFlavor(versions, flavors.FilterOpts{FlavorName: expectedFlavor.Name})
	th.AssertDeepEquals(t, expectedFlavor, *foundByName)
}
