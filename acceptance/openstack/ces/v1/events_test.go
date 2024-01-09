package v1

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/events"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEvents(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)
	name := tools.RandomString("event_test_", 3)
	currentTime := time.Now().Unix() * 1000
	t.Logf("Attempting to create CES Event: %s", name)
	event, err := events.CreateEvents(client, []events.EventItem{
		{
			EventName:   name,
			EventSource: "SYS.ECS",
			Time:        currentTime,
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
	th.AssertEquals(t, event[0].EventName, name)

	t.Log("List CES Events")
	eventsRes, err := events.ListEvents(client, events.ListEventsOpts{
		From:  currentTime,
		To:    currentTime + 100000,
		Limit: 10,
	})
	th.AssertNoErr(t, err)

	t.Log("List CES Event Details")
	_, err = events.ListEventDetail(client, events.ListEventDetailOpts{
		EventName: eventsRes.Events[0].EventName,
		EventType: eventsRes.Events[0].EventType,
		Limit:     10,
		From:      currentTime,
		To:        currentTime + 100000,
	})
	th.AssertNoErr(t, err)
}
