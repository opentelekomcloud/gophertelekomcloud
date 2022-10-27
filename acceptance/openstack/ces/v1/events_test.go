package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/events"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEvents(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	eventsRes, err := events.ListEvents(client, events.ListEventsOpts{
		Limit: 1,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, eventsRes.MetaData.Total, 1)

	detailRes, err := events.ListEventDetail(client, events.ListEventDetailOpts{
		EventName: eventsRes.Events[0].EventName,
		EventType: eventsRes.Events[0].EventType,
		Limit:     1,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, detailRes.MetaData.Total, 1)
}
