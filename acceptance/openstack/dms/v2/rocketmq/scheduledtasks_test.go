package rocketmq

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/scheduledtasks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRocketMQScheduledTasks(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	listed, err := scheduledtasks.List(client, instanceID, scheduledtasks.ListOpts{
		Start:     1,
		Limit:     10,
		BeginTime: "19700101000000",
		EndTime:   "20991231000000",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listed)
	th.AssertEquals(t, 0, listed.JobCount)
	th.AssertEquals(t, 0, len(listed.Jobs))

	// Scheduled tasks are only created by operations scheduled from the
	// console, which have no API counterpart. DMS does not validate the task
	// ID: requests for a nonexistent task succeed.
	taskID := "ff80808288ae8b6a0188ae91f6fc000d"

	t.Logf("Attempting to cancel RocketMQ scheduled task: %s", taskID)
	th.AssertNoErr(t, scheduledtasks.Update(client, instanceID, taskID, scheduledtasks.UpdateOpts{Status: "CANCELLED"}))
	t.Logf("Attempting to reschedule RocketMQ scheduled task: %s", taskID)
	th.AssertNoErr(t, scheduledtasks.Update(client, instanceID, taskID, scheduledtasks.UpdateOpts{ExecuteAt: "20991231000000"}))
	t.Logf("Attempting to delete RocketMQ scheduled task: %s", taskID)
	th.AssertNoErr(t, scheduledtasks.Delete(client, instanceID, taskID))
}
