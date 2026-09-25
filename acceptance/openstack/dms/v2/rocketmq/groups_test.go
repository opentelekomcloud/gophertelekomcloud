package rocketmq

import (
	"strings"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/groups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	groupNotOnline = "DMS.00050004"
	groupNotFound  = "DMS.00050012"
)

func TestRocketMQConsumerGroups(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	t.Run("LifeCycle", func(t *testing.T) {
		testRocketMQConsumerGroupLifeCycle(t, client, instanceID)
	})
	t.Run("BatchDelete", func(t *testing.T) {
		testRocketMQConsumerGroupBatchDelete(t, client, instanceID)
	})
}

func testRocketMQConsumerGroupLifeCycle(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	groupName := createRocketMQGroup(t, client, instanceID)

	group, err := groups.Get(client, instanceID, groupName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, groupName, group.Name)
	th.AssertEquals(t, "consumer group description", group.Description)
	th.AssertEquals(t, 16, group.RetryMaxTime)
	th.AssertEquals(t, true, group.Enabled)
	th.AssertEquals(t, false, group.Broadcast)
	th.AssertEquals(t, false, group.ConsumeOrderly)

	listed, err := groups.List(client, instanceID, groups.ListOpts{Group: groupName})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listed.Groups))
	th.AssertEquals(t, groupName, listed.Groups[0].Name)

	t.Logf("Attempting to update RocketMQ consumer group: %s", groupName)
	err = groups.Update(client, instanceID, groupName, groups.UpdateOpts{
		Enabled:      pointerto.Bool(true),
		Broadcast:    pointerto.Bool(true),
		RetryMaxTime: 10,
	})
	th.AssertNoErr(t, err)

	group, err = groups.Get(client, instanceID, groupName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 10, group.RetryMaxTime)
	th.AssertEquals(t, true, group.Broadcast)

	t.Logf("Attempting to batch update RocketMQ consumer group: %s", groupName)
	job, err := groups.BatchUpdate(client, instanceID, groups.BatchUpdateOpts{
		Groups: []groups.BatchUpdateGroup{
			{
				Name:           groupName,
				Enabled:        pointerto.Bool(true),
				Broadcast:      pointerto.Bool(false),
				ConsumeOrderly: pointerto.Bool(false),
				RetryMaxTime:   5,
				Description:    pointerto.String("updated description"),
			},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, job.JobID != "")

	th.AssertNoErr(t, golangsdk.WaitFor(300, func() (bool, error) {
		g, err := groups.Get(client, instanceID, groupName)
		if err != nil {
			return false, err
		}
		return g.RetryMaxTime == 5, nil
	}))
	group, err = groups.Get(client, instanceID, groupName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated description", group.Description)
	th.AssertEquals(t, false, group.Broadcast)

	topics, err := groups.ListTopics(client, instanceID, groupName, groups.ListTopicsOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, topics.Total)

	topicName := createRocketMQTopic(t, client, instanceID)

	t.Logf("Attempting to reset consumer offset of group %s for topic %s", groupName, topicName)
	reset, err := groups.ResetOffset(client, instanceID, groupName, groups.ResetOffsetOpts{
		Topic:     topicName,
		Timestamp: time.Now().UnixMilli(),
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, reset)

	details, err := groups.ListTopics(client, instanceID, groupName, groups.ListTopicsOpts{Topic: topicName})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, details)

	// The consumer list is only available while a consumer is connected to
	// the group, which the test has no client for.
	_, err = groups.ListClients(client, instanceID, groupName, groups.ListClientsOpts{
		Limit:    10,
		IsDetail: true,
	})
	th.AssertEquals(t, true, hasErrorCode(err, groupNotOnline))

	t.Logf("Attempting to delete RocketMQ consumer group: %s", groupName)
	th.AssertNoErr(t, groups.Delete(client, instanceID, groupName))
	_, err = groups.Get(client, instanceID, groupName)
	th.AssertEquals(t, true, hasErrorCode(err, groupNotFound))
}

func testRocketMQConsumerGroupBatchDelete(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	names := []string{
		createRocketMQGroup(t, client, instanceID),
		createRocketMQGroup(t, client, instanceID),
	}

	t.Logf("Attempting to batch delete RocketMQ consumer groups: %v", names)
	job, err := groups.BatchDelete(client, instanceID, groups.BatchDeleteOpts{Groups: names})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, job.JobID != "")

	for _, name := range names {
		th.AssertNoErr(t, golangsdk.WaitFor(300, func() (bool, error) {
			_, err := groups.Get(client, instanceID, name)
			if hasErrorCode(err, groupNotFound) {
				return true, nil
			}
			return false, err
		}))
	}
}

// prepareRocketMQInstance reuses the instance from OS_DMS_ROCKETMQ_INSTANCE_ID
// when it is set and creates a new one otherwise, which takes ~30 minutes.
func prepareRocketMQInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	if instanceID := clients.EnvOS.GetEnv("DMS_ROCKETMQ_INSTANCE_ID"); instanceID != "" {
		return instanceID
	}
	if clients.EnvOS.GetEnv("RUN_DMS_ROCKETMQ") == "" {
		t.Skip("Neither OS_DMS_ROCKETMQ_INSTANCE_ID nor OS_RUN_DMS_ROCKETMQ env var is set, RocketMQ instance creation takes ~30 minutes")
	}

	instanceID := createRocketMQInstance(t, client)
	t.Cleanup(func() { deleteRocketMQInstance(t, client, instanceID) })
	return instanceID
}

func createRocketMQGroup(t *testing.T, client *golangsdk.ServiceClient, instanceID string) string {
	t.Logf("Attempting to create RocketMQ consumer group")

	group, err := groups.Create(client, instanceID, groups.CreateOpts{
		Name:           tools.RandomString("group-acc-", 8),
		Description:    "consumer group description",
		Enabled:        pointerto.Bool(true),
		Broadcast:      pointerto.Bool(false),
		ConsumeOrderly: pointerto.Bool(false),
		RetryMaxTime:   16,
	})
	th.AssertNoErr(t, err)
	t.Logf("RocketMQ consumer group created: %s", group.Name)

	t.Cleanup(func() {
		err := groups.Delete(client, instanceID, group.Name)
		if err != nil && !hasErrorCode(err, groupNotFound) {
			t.Errorf("failed to delete RocketMQ consumer group %s: %s", group.Name, err)
		}
	})

	return group.Name
}

// createRocketMQTopic creates a topic directly through the API, as the SDK has
// no RocketMQ topic management yet.
// Send POST to /v2/{project_id}/instances/{instance_id}/topics
func createRocketMQTopic(t *testing.T, client *golangsdk.ServiceClient, instanceID string) string {
	name := tools.RandomString("topic-acc-", 8)
	t.Logf("Attempting to create RocketMQ topic: %s", name)

	_, err := client.Post(client.ServiceURL("instances", instanceID, "topics"), map[string]interface{}{
		"name":         name,
		"queue_num":    3,
		"message_type": "NORMAL",
	}, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		_, err := client.Delete(client.ServiceURL("instances", instanceID, "topics", name), &golangsdk.RequestOpts{
			OkCodes: []int{204},
		})
		if err != nil {
			t.Errorf("failed to delete RocketMQ topic %s: %s", name, err)
		}
	})

	return name
}

// hasErrorCode reports whether err is a 400 response with the given DMS error
// code. RocketMQ answers with 400 rather than 404 for a missing consumer group.
func hasErrorCode(err error, code string) bool {
	e, ok := err.(golangsdk.ErrDefault400)
	return ok && strings.Contains(string(e.Body), code)
}
