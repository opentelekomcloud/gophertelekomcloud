package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v1/apiversions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := apiversions.List(client.ServiceClient())
	th.AssertNoErr(t, err)

	expected := []apiversions.APIVersion{
		{
			ID:      "v1.0",
			Status:  "CURRENT",
			Updated: "2012-01-04T11:33:21Z",
		},
		{
			ID:      "v2.0",
			Status:  "CURRENT",
			Updated: "2012-11-21T11:33:21Z",
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestAPIInfo(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := apiversions.Get(client.ServiceClient(), "v1")
	th.AssertNoErr(t, err)

	expected := apiversions.APIVersion{
		ID:      "v1.0",
		Status:  "CURRENT",
		Updated: "2012-01-04T11:33:21Z",
	}

	th.AssertEquals(t, actual.ID, expected.ID)
	th.AssertEquals(t, actual.Status, expected.Status)
	th.AssertEquals(t, actual.Updated, expected.Updated)
}
