package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/schedulerstats"

	"github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListStoragePoolsDetail(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleStoragePoolsListSuccessfully(t)

	actual, err := schedulerstats.List(client.ServiceClient(), schedulerstats.ListOpts{Detail: true})
	testhelper.AssertNoErr(t, err)

	if len(actual) != 2 {
		t.Fatalf("Expected 2 backends, got %d", len(actual))
	}
	testhelper.CheckDeepEquals(t, StoragePoolFake1, actual[0])
	testhelper.CheckDeepEquals(t, StoragePoolFake2, actual[1])
}
