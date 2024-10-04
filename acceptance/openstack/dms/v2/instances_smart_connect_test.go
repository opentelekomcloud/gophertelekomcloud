package v2

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/smart_connect"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSmartConnectDMS(t *testing.T) {
	ak, sk := os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY")
	if ak == "" || sk == "" {
		t.Skip("ACCESS_KEY and SECRET_KEY are required for this test")
	}

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	// Create DMS instance
	instanceID := createDmsInstance(t, client)
	t.Cleanup(func() {
		deleteDmsInstance(t, client, instanceID)
	})

	t.Logf("Attempting to enable Smart Connect for instance %s", instanceID)
	enableOpts := smart_connect.EnableOpts{
		InstanceId: instanceID,
	}
	_, err = smart_connect.Enable(client, enableOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		// Disable Smart Connect
		_, err = smart_connect.Disable(client, instanceID)
		th.AssertNoErr(t, err)
		err = waitForInstanceAvailable(client, 600, instanceID)
		th.AssertNoErr(t, err)

		t.Logf("Smart Connect disabled for instance %s", instanceID)
	})

	err = waitForInstanceAvailable(client, 600, instanceID)
	th.AssertNoErr(t, err)

	t.Logf("Smart Connect enabled for instance %s", instanceID)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))

	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)
	_, err = obsClient.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = obsClient.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	// Create a task
	createTaskOpts := smart_connect.CreateTaskOpts{
		TaskName:    "test-task",
		SourceType:  "NONE",
		StartLater:  pointerto.Bool(true),
		TopicsRegex: "topic-obs",
		SinkType:    "OBS_SINK",
		SinkTask: &smart_connect.SmartConnectTaskSinkConfig{
			ConsumerStrategy:    "earliest",
			DestinationFileType: "TEXT",
			DeliverTimeInterval: 300,
			AccessKey:           ak,
			SecretKey:           sk,
			ObsBucketName:       bucketName,
			ObsPath:             "obsTransfer-test",
			PartitionFormat:     "yyyy/MM/dd/HH/mm",
			RecordDelimiter:     ",",
			StoreKeys:           pointerto.Bool(false),
		},
	}
	createResp, err := smart_connect.CreateTask(client, instanceID, createTaskOpts)
	th.AssertNoErr(t, err)
	taskID := createResp.ID
	t.Logf("Task created with ID: %s", taskID)

	t.Cleanup(func() {
		// Delete task
		err = smart_connect.DeleteTask(client, instanceID, taskID)
		th.AssertNoErr(t, err)
		t.Logf("Task %s deleted", taskID)
	})

	// Get task details
	taskDetails, err := smart_connect.GetTask(client, instanceID, taskID)
	th.AssertNoErr(t, err)
	t.Logf("Got task details for task %s", taskID)
	th.AssertEquals(t, taskDetails.TaskName, "test-task")
	th.AssertEquals(t, taskDetails.TopicsRegex, "topic-obs")
	th.AssertEquals(t, taskDetails.SinkType, "OBS_SINK")

	// List tasks
	listOpts := smart_connect.QueryTasksOpts{
		InstanceId: instanceID,
	}
	listResp, err := smart_connect.ListTasks(client, listOpts)
	th.AssertNoErr(t, err)
	t.Logf("Listed %d tasks", len(listResp.Tasks))
	th.AssertEquals(t, listResp.Tasks[0].TaskName, "test-task")
	th.AssertEquals(t, listResp.Tasks[0].TopicsRegex, "topic-obs")
	th.AssertEquals(t, listResp.Tasks[0].SinkType, "OBS_SINK")

	// Start or restart task
	err = smart_connect.StartOrRestartTask(client, instanceID, taskID)
	th.AssertNoErr(t, err)
	t.Logf("Task %s started or restarted", taskID)

	// Pause task
	err = smart_connect.PauseTask(client, instanceID, taskID)
	th.AssertNoErr(t, err)
	t.Logf("Task %s paused", taskID)

	// Restart task
	err = smart_connect.RestartTask(client, instanceID, taskID)
	th.AssertNoErr(t, err)
	t.Logf("Task %s restarted", taskID)
}
