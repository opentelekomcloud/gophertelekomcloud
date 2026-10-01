package rocketmq

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/migration"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// migrationTaskError is returned both for a missing metadata migration task
// and for a rabbitToRocket task queried without a type.
const migrationTaskError = "DMS.00405011"

func TestRocketMQMetadataMigration(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	t.Run("RocketMQ", func(t *testing.T) {
		testRocketMQMetadataMigrationFromRocketMQ(t, client, instanceID)
	})
	t.Run("RabbitMQ", func(t *testing.T) {
		testRocketMQMetadataMigrationFromRabbitMQ(t, client, instanceID)
	})
}

func testRocketMQMetadataMigrationFromRocketMQ(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	topicName := tools.RandomString("topic-acc-", 8)
	groupName := tools.RandomString("group-acc-", 8)

	t.Cleanup(func() {
		err := topics.Delete(client, instanceID, topicName)
		if err != nil && !hasErrorCode(err, topicNotFound) {
			t.Errorf("failed to delete RocketMQ topic %s: %s", topicName, err)
		}
		err = groups.Delete(client, instanceID, groupName)
		if err != nil && !hasErrorCode(err, groupNotFound) {
			t.Errorf("failed to delete RocketMQ consumer group %s: %s", groupName, err)
		}
	})

	taskID := createRocketMQMigrationTask(t, client, instanceID, migration.CreateOpts{
		Overwrite: true,
		Name:      tools.RandomString("task-acc-", 8),
		Type:      "rocketmq",
		TopicConfigTable: map[string]migration.TopicConfig{
			topicName: {
				TopicName:       topicName,
				Perm:            6,
				ReadQueueNums:   3,
				WriteQueueNums:  3,
				TopicFilterType: "SINGLE_TAG",
			},
		},
		SubscriptionGroupTable: map[string]migration.SubscriptionGroup{
			groupName: {
				GroupName:                      groupName,
				ConsumeBroadcastEnable:         pointerto.Bool(false),
				ConsumeEnable:                  pointerto.Bool(true),
				ConsumeFromMinEnable:           pointerto.Bool(true),
				NotifyConsumerIDsChangedEnable: pointerto.Bool(true),
				RetryMaxTimes:                  2,
				RetryQueueNums:                 1,
				WhichBrokerWhenConsumeSlow:     1,
			},
		},
	})

	task := waitForRocketMQMigrationTask(t, client, instanceID, taskID, migration.GetOpts{})
	tools.PrintResource(t, task)
	th.AssertEquals(t, "rocketmq", task.Type)
	th.AssertEquals(t, true, strings.Contains(task.JSONContent, topicName))
	th.AssertEquals(t, true, strings.Contains(task.JSONContent, groupName))
	th.AssertEquals(t, true, strings.Contains(task.Result, topicName))

	topic, err := topics.Get(client, instanceID, topicName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, topic)
	th.AssertEquals(t, topicName, topic.Name)

	group, err := groups.Get(client, instanceID, groupName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, group)
	th.AssertEquals(t, groupName, group.Name)
	th.AssertEquals(t, 2, group.RetryMaxTime)

	deleteRocketMQMigrationTask(t, client, instanceID, taskID)
}

func testRocketMQMetadataMigrationFromRabbitMQ(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	vhost := tools.RandomString("vhost-acc-", 4)

	taskID := createRocketMQMigrationTask(t, client, instanceID, migration.CreateOpts{
		Overwrite: true,
		Name:      tools.RandomString("task-acc-", 8),
		Type:      "rabbitToRocket",
		Vhosts:    []migration.Vhost{{Name: vhost}},
		Queues:    []migration.Queue{{Vhost: vhost, Name: "queue-acc", Durable: true}},
		Exchanges: []migration.Exchange{{Vhost: vhost, Name: "exchange-acc", Type: "direct", Durable: true}},
		Bindings: []migration.Binding{
			{
				Vhost:           vhost,
				Source:          "exchange-acc",
				Destination:     "queue-acc",
				DestinationType: "queue",
				RoutingKey:      "routing-key-acc",
			},
		},
	})

	// Details of rabbitToRocket tasks can only be queried by vhost,
	// exchange or queue.
	task := waitForRocketMQMigrationTask(t, client, instanceID, taskID, migration.GetOpts{Type: "vhost"})
	tools.PrintResource(t, task)
	th.AssertEquals(t, "rabbitToRocket", task.Type)
	th.AssertEquals(t, true, strings.Contains(task.JSONContent, vhost))

	_, err := migration.Get(client, instanceID, taskID, migration.GetOpts{})
	th.AssertEquals(t, true, hasErrorCode(err, migrationTaskError))

	exchanges, err := migration.Get(client, instanceID, taskID, migration.GetOpts{Type: "exchange", Name: vhost})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, exchanges)

	deleteRocketMQMigrationTask(t, client, instanceID, taskID)
}

func createRocketMQMigrationTask(t *testing.T, client *golangsdk.ServiceClient, instanceID string, opts migration.CreateOpts) string {
	t.Logf("Attempting to create RocketMQ metadata migration task: %s", opts.Name)
	created, err := migration.Create(client, instanceID, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, created.TaskID != "")
	t.Logf("RocketMQ metadata migration task created: %s", created.TaskID)

	// Deleting a task that no longer exists succeeds with an empty list.
	t.Cleanup(func() {
		_, err := migration.BatchDelete(client, instanceID, migration.BatchDeleteOpts{TaskIDs: []string{created.TaskID}})
		if err != nil {
			t.Errorf("failed to delete RocketMQ metadata migration task %s: %s", created.TaskID, err)
		}
	})

	return created.TaskID
}

func waitForRocketMQMigrationTask(t *testing.T, client *golangsdk.ServiceClient, instanceID, taskID string, opts migration.GetOpts) *migration.TaskDetails {
	var task *migration.TaskDetails
	th.AssertNoErr(t, golangsdk.WaitFor(300, func() (bool, error) {
		var err error
		task, err = migration.Get(client, instanceID, taskID, opts)
		if err != nil {
			return false, err
		}
		return task.Status == "finished" || task.Status == "failed", nil
	}))
	th.AssertEquals(t, taskID, task.ID)
	th.AssertEquals(t, "finished", task.Status)

	listed, err := migration.List(client, instanceID, migration.ListOpts{Offset: 1, Limit: 50})
	th.AssertNoErr(t, err)
	var found bool
	for _, tk := range listed.Tasks {
		if tk.ID == taskID {
			found = true
			th.AssertEquals(t, task.Name, tk.Name)
			th.AssertEquals(t, task.Type, tk.Type)
			th.AssertEquals(t, "finished", tk.Status)
			break
		}
	}
	th.AssertEquals(t, true, found)

	return task
}

func deleteRocketMQMigrationTask(t *testing.T, client *golangsdk.ServiceClient, instanceID, taskID string) {
	t.Logf("Attempting to delete RocketMQ metadata migration task: %s", taskID)
	deleted, err := migration.BatchDelete(client, instanceID, migration.BatchDeleteOpts{TaskIDs: []string{taskID}})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []string{taskID}, deleted.SuccessTaskList)

	_, err = migration.Get(client, instanceID, taskID, migration.GetOpts{})
	th.AssertEquals(t, true, hasErrorCode(err, migrationTaskError))
}
