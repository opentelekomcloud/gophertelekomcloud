package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rts/v1/stackresources"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListResources(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, ListOutput)

	//count := 0
	actual, err := stackresources.List(fake.ServiceClient(), "hello_world", stackresources.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract resources: %v", err)
	}
	th.AssertDeepEquals(t, ListExpected, actual)
}
