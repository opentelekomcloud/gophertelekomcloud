package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/queues"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDmsQueueList(t *testing.T) {
	client, err := clients.NewDmsV1Client()
	th.AssertNoErr(t, err)

	queueAllPages, err := queues.List(client, true).AllPages()
	th.AssertNoErr(t, err)
	queueInstances, err := queues.ExtractQueues(queueAllPages)
	th.AssertNoErr(t, err)
	for _, val := range queueInstances {
		tools.PrintResource(t, val)
	}
}

func TestDmsQueueLifeCycle(t *testing.T) {
	client, err := clients.NewDmsV1Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create DMSv1 queue")
	createOpts := queues.CreateOpts{
		Name: "test-queue",
	}
	queue, err := queues.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created DMSv1 queue: %s", queue.ID)
	defer func() {
		t.Logf("Attempting to delete DMSv1 queue: %s", queue.ID)
		err := queues.Delete(client, queue.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted DMSv1 queue: %s", queue.ID)
	}()

	newQueue, err := queues.Get(client, queue.ID, true).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Name, newQueue.Name)
}
