package v2

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/smart_connect"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/credentials"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSmartConnectDMS(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	ak, sk := obsSinkCredentials(t)

	instanceID := smartConnectInstance(t, client)

	if !connectorEnabled(t, client, instanceID) {
		t.Logf("Attempting to enable Smart Connect for instance %s", instanceID)
		_, err = smart_connect.Enable(client, smart_connect.EnableOpts{
			InstanceId: instanceID,
		})
		th.AssertNoErr(t, err)

		t.Cleanup(func() {
			_, err := smart_connect.Disable(client, instanceID)
			th.AssertNoErr(t, err)
			th.AssertNoErr(t, waitForConnector(client, 900, instanceID, false))

			t.Logf("Smart Connect disabled for instance %s", instanceID)
		})

		th.AssertNoErr(t, waitForConnector(client, 900, instanceID, true))

		t.Logf("Smart Connect enabled for instance %s", instanceID)
	}

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

// obsSinkCredentials returns ACCESS_KEY/SECRET_KEY from the environment, if set.
// Otherwise it creates a permanent AK/SK for the current user and deletes it on cleanup:
// the OBS_SINK task body has no security token field, so temporary AK/SK cannot be used.
func obsSinkCredentials(t *testing.T) (string, string) {
	t.Helper()

	if ak, sk := os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY"); ak != "" && sk != "" {
		return ak, sk
	}

	iamClient, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	cred, err := credentials.Create(iamClient, credentials.CreateOpts{
		UserID:      iamClient.UserID,
		Description: "DMS Smart Connect acceptance test key",
	}).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created access key for user %s", iamClient.UserID)

	t.Cleanup(func() {
		th.AssertNoErr(t, credentials.Delete(iamClient, cred.AccessKey).ExtractErr())
		t.Logf("Test access key deleted")
	})

	return cred.AccessKey, cred.SecretKey
}

// smartConnectInstance uses the existing instance from OS_DMS_KAFKA_INSTANCE_NAME, if set; otherwise it creates a temporary one.
func smartConnectInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Helper()

	if name := clients.EnvOS.GetEnv("DMS_KAFKA_INSTANCE_NAME"); name != "" {
		return existingKafkaInstance(t, client, name).InstanceID
	}

	instanceID := createDmsInstance(t, client)
	t.Cleanup(func() {
		deleteDmsInstance(t, client, instanceID)
	})
	return instanceID
}

// connectorState reports whether Smart Connect is enabled on the instance.
// connector_enable is not returned by instance GET/List, so the connector tasks endpoint is used:
// it answers 404 (DMS.00404024 "The connector does not exist.") while Smart Connect is disabled.
func connectorState(client *golangsdk.ServiceClient, instanceID string) (bool, error) {
	_, err := smart_connect.ListTasks(client, smart_connect.QueryTasksOpts{InstanceId: instanceID})
	if err == nil {
		return true, nil
	}
	var notFound golangsdk.ErrDefault404
	if errors.As(err, &notFound) {
		return false, nil
	}
	return false, err
}

func connectorEnabled(t *testing.T, client *golangsdk.ServiceClient, instanceID string) bool {
	t.Helper()

	enabled, err := connectorState(client, instanceID)
	th.AssertNoErr(t, err)
	return enabled
}

// waitForConnector waits until the instance is RUNNING and the Smart Connect state matches enabled.
// The instance status alone is not enough: it may be RUNNING before the Smart Connect job finishes.
// Unexpected errors are tolerated while waiting and reported if the wait times out.
func waitForConnector(client *golangsdk.ServiceClient, secs int, instanceID string, enabled bool) error {
	var lastErr error
	err := golangsdk.WaitFor(secs, func() (bool, error) {
		instance, err := lifecycle.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		if instance.Status != dmsTargetStatus {
			return false, nil
		}
		state, err := connectorState(client, instanceID)
		if err != nil {
			lastErr = err
			return false, nil
		}
		return state == enabled, nil
	})
	if err != nil && lastErr != nil {
		return fmt.Errorf("%w (last connector error: %v)", err, lastErr)
	}
	return err
}
