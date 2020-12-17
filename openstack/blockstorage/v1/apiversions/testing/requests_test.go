package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v1/apiversions"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	count := 0

	_ = apiversions.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := apiversions.ExtractAPIVersions(page)
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

		return true, nil
	})

	th.AssertEquals(t, 1, count)
}

func TestAPIInfo(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := apiversions.Get(client.ServiceClient(), "v1").Extract()
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
