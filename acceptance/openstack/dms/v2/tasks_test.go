package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/tasks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackgroundTasks(t *testing.T) {
	name := kafkaInstanceName(t)

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := existingKafkaInstance(t, client, name).InstanceID

	list, err := tasks.List(client, instanceID, tasks.ListOpts{Limit: 10})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)
	if len(list.Tasks) == 0 {
		t.Skip("no background tasks on the instance")
	}

	task, err := tasks.Get(client, instanceID, list.Tasks[0].ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, task)
	th.AssertEquals(t, 1, len(task.Tasks))
	th.AssertEquals(t, list.Tasks[0].ID, task.Tasks[0].ID)
	th.AssertEquals(t, list.Tasks[0].Name, task.Tasks[0].Name)
}
