package v2

import (
	"os"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/trigger"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionTriggerLifecycle(t *testing.T) {
	agency := os.Getenv("AGENCY")
	if agency == "" {
		t.Skip("`AGENCY`needs to be defined to run this test")
	}

	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	clientLts, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraphAgency(t, client, agency)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}

	logId, err := groups.CreateLogGroup(clientLts, createOpts)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, id string) {
		err = groups.DeleteLogGroup(clientLts, logId)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	streamId, err := streams.CreateLogStream(clientLts, streams.CreateOpts{
		GroupId:       logId,
		LogStreamName: tools.RandomString("test-stream-", 3),
	})
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, id string) {
		err = streams.DeleteLogStream(clientLts, streams.DeleteOpts{
			GroupId:  logId,
			StreamId: streamId,
		})
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	createTriggerOpts := trigger.CreateOpts{
		FuncUrn:         funcUrn,
		TriggerTypeCode: "LTS",
		TriggerStatus:   "ACTIVE",
		EventData: map[string]interface{}{
			"log_group_id": logId,
			"log_topic_id": streamId,
		},
	}

	t.Logf("Attempting to CREATE FUNCGRAPH TRIGGER")
	createTriggerResp, err := trigger.Create(client, createTriggerOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, createTriggerResp)

	defer func(client *golangsdk.ServiceClient, urn, triggerType, id string) {
		t.Logf("Attempting to DELETE FUNCGRAPH TRIGGER")
		err = trigger.Delete(client, urn, triggerType, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn, "LTS", createTriggerResp.TriggerId)

	t.Logf("Attempting to GET FUNCGRAPH TRIGGER")
	getTriggerResp, err := trigger.Get(client, funcUrn, "LTS", createTriggerResp.TriggerId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createTriggerResp.TriggerId, getTriggerResp.TriggerId)
	th.AssertEquals(t, createTriggerResp.TriggerStatus, getTriggerResp.TriggerStatus)
	th.AssertEquals(t, createTriggerResp.CreatedTime, getTriggerResp.CreatedTime)

	updateTriggerOpts := trigger.UpdateOpts{
		FuncUrn:         funcUrn,
		TriggerId:       getTriggerResp.TriggerId,
		TriggerTypeCode: getTriggerResp.TriggerTypeCode,
		TriggerStatus:   "DISABLED",
	}

	t.Logf("Attempting to UPDATE FUNCGRAPH TRIGGER")
	updateTriggerResp, err := trigger.Update(client, updateTriggerOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateTriggerResp.TriggerStatus, "DISABLED")

}

func createFunctionGraphAgency(t *testing.T, client *golangsdk.ServiceClient, agency string) (*function.FuncGraph, string) {
	funcName := "funcgraph-" + tools.RandomString("acctest", 4)

	createOpts := function.CreateOpts{
		Name:       funcName,
		Package:    "default",
		Runtime:    "Python3.9",
		Timeout:    200,
		Handler:    "index.py",
		MemorySize: 512,
		CodeType:   "zip",
		Xrole:      agency,
		CodeURL:    "https://regr-func-graph.obs.eu-de.otc.t-systems.com/index.py",
	}

	createResp, err := function.Create(client, createOpts)
	th.AssertNoErr(t, err)

	return createResp, funcName
}
