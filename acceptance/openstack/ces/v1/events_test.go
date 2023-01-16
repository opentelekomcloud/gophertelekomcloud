package v1

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/events"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEvents(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	event, err := events.CreateEvents(client, []events.EventItem{
		{
			EventName:   "test",
			EventSource: "SYS.ECS",
			Time:        time.Now().Unix() * 1000,
			Detail: events.EventItemDetail{
				Content:      "The financial system was invaded",
				ResourceId:   "9d3bc7be-5181-4c5a-9d15-26aac9da91b7",
				ResourceName: "ecs-CES-testing",
				EventState:   "normal",
				EventLevel:   "Major",
				EventUser:    "zuul",
			},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, event[0].EventName, "test")

	eventsRes, err := events.ListEvents(client, events.ListEventsOpts{
		Limit: 1,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(eventsRes.Events), 1)

	detailRes, err := events.ListEventDetail(client, events.ListEventDetailOpts{
		EventName: eventsRes.Events[0].EventName,
		EventType: eventsRes.Events[0].EventType,
		Limit:     1,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(detailRes.EventInfo), 1)
}
