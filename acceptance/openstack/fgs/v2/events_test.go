package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/events"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphEventsLifecycle(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	t.Logf("Attempting to CREATE TEST EVENT")

	eventName := tools.RandomString("funcgraph-event", 4)

	createOpts := events.CreateOpts{
		FuncUrn: funcUrn,
		Name:    eventName,
		Content: "eyJrIjoidiJ9",
	}

	eventResp, err := events.Create(client, createOpts)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, funcUrn, eventId string) {
		t.Logf("Attempting to DELETE TEST EVENT")
		err = events.Delete(client, funcUrn, eventId)
		th.AssertNoErr(t, err)
	}(client, funcUrn, eventResp.Id)

	th.AssertEquals(t, eventResp.Name, eventName)

	t.Logf("Attempting to LIST TEST EVENTS")
	listResp, err := events.List(client, funcUrn)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, eventResp.Name, listResp.Events[0].Name)
	th.AssertEquals(t, eventResp.Id, listResp.Events[0].Id)

	newContent := "ewogICAgImJvZHkiOiAiIiwKICAgICJy"
	t.Logf("Attempting to UPDATE TEST EVENT")
	updateOpts := events.UpdateOpts{
		FuncUrn: funcUrn,
		EventId: eventResp.Id,
		Content: newContent,
	}

	updateResp, err := events.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to GET TEST EVENT")
	getResp, err := events.Get(client, funcUrn, eventResp.Id)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, getResp.Name, updateResp.Name)
	th.AssertEquals(t, getResp.Id, updateResp.Id)
	th.AssertEquals(t, getResp.Content, updateOpts.Content)

	tools.PrintResource(t, getResp)

}
