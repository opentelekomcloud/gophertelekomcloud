package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	quick_search "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/quick-search"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsQuickSearchLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	clientV10, err := clients.NewLtsV10Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}
	t.Logf("Attempting to Create Log Group")
	group, err := groups.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(client, group)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Log Stream")
	sname := tools.RandomString("test-stream-", 3)
	stream, err := streams.Create(client, streams.CreateOpts{
		GroupId:       group,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(client, streams.DeleteOpts{
			GroupId:  group,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Quick Search")
	qsname := tools.RandomString("test-qs-", 3)
	qs, err := quick_search.Create(clientV10, group, stream, quick_search.CreateOpts{
		Criteria:   "content : 1234567891234567891234567891234567891234567891234567891234567894",
		Name:       qsname,
		SearchType: "ORIGINALLOG",
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Quick Search")
		err = quick_search.Delete(clientV10, group, stream, quick_search.DeleteOpts{
			ID: qs,
		})
		th.AssertNoErr(t, err)
	})

	criterias, err := quick_search.ListCriterias(clientV10, group, stream, quick_search.ListOpts{
		SearchType: "ORIGINALLOG",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(criterias))

	groupCriterias, err := quick_search.ListGroupCriterias(clientV10, group)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(groupCriterias))
}
