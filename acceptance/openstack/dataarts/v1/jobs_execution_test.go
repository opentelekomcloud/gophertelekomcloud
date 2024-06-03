package v1_1

import (
	"fmt"
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/connection"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/script"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDataArtsJobExecution(t *testing.T) {
	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	workspace := os.Getenv("AWS_WORKSPACE")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	kms := clients.EnvOS.GetEnv("KMS_ID")

	clientV1, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	ec, c := prepareEnv(t)

	t.Log("create a connection")

	getConnection(t, clientV1, ec, c.Name, kms, vpcID)
	t.Cleanup(func() {
		t.Logf("attempting to delete connection: %s", connectionName)
		err := connection.Delete(clientV1, connectionName, workspace)
		th.AssertNoErr(t, err)
		t.Logf("connection is deleted: %s", connectionName)
	})

	getScript(t, clientV1, connectionName, workspace)
	t.Cleanup(func() {
		t.Logf("attempting to delete script: %s", scriptName)
		opts := script.DeleteReq{Workspace: workspace}
		err := script.Delete(clientV1, scriptName, opts)
		th.AssertNoErr(t, err)
		t.Logf("script is deleted: %s", scriptName)
	})

	t.Log("create a job")

	createOpts := &job.Job{
		Name: jobName,
		Nodes: []job.Node{
			{
				Name: "testNode",
				Type: "Shell",
				Location: job.Location{
					X: -332,
					Y: -150,
				},
				PollingInterval:  20,
				MaxExecutionTime: 360,
				RetryTimes:       0,
				RetryInterval:    120,
				FailPolicy:       "FAIL_CHILD",
				Properties: []*job.Property{
					{
						Name:  "scriptName",
						Value: "testScript",
					},
					{
						Name:  "connectionName",
						Value: "testConnection",
					},
					{
						Name:  "statementOrScript",
						Value: "SCRIPT",
					},
					{
						Name:  "emptyRunningJob",
						Value: "0",
					},
				},
			},
		},
		Schedule: job.Schedule{
			Type: "EXECUTE_ONCE",
		},
		Params:       nil,
		ProcessType:  "BATCH",
		LogPath:      "",
		TargetStatus: "",
		Approvers:    nil,
	}

	err = job.Create(clientV1, *createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule job cleanup")
	t.Cleanup(func() {
		jobCleanup(t, clientV1)
	})

	t.Log("execute a job")

	_, err = job.ExecuteImmediately(clientV1, jobName, nil)
	th.AssertNoErr(t, err)

	t.Log("get job's running status")
	resp, err := job.GetRunningStatus(clientV1, jobName, "")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "STOPPED", resp.Status)

	opts := job.GetJobInstanceListReq{
		JobName: jobName,
	}

	t.Log("trying to get job results in 60 seconds")

	err = golangsdk.WaitFor(120, func() (bool, error) {
		list, err := job.GetJobInstanceList(clientV1, opts)
		if err != nil {
			return false, err
		}

		for _, j := range list.Instances {
			if j.Status == "success" {
				return true, nil
			}
		}

		return false, nil
	})
	th.AssertNoErr(t, err)
}

func getScript(t *testing.T, client *golangsdk.ServiceClient, conn, workspace string) *script.Script {
	t.Helper()

	s, err := script.Get(client, scriptName, workspace)
	if err == nil {
		return s
	}

	errNew, ok := err.(golangsdk.ErrDefault400)
	th.AssertEquals(t, true, ok)

	if errNew.Actual != 400 {
		th.AssertNoErr(t, err)
	}

	t.Log("create a script")

	createOpts := script.Script{
		Name:           scriptName,
		Type:           "Shell",
		Content:        "echo 123456",
		ConnectionName: conn,
	}

	err = script.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Logf("waiting for the script to be created, 3 seconds")
	time.Sleep(3 * time.Second)

	s, err = script.Get(client, scriptName, workspace)
	th.AssertNoErr(t, err)
	return s
}

func getConnection(t *testing.T, client *golangsdk.ServiceClient, ec *cloudservers.CloudServer, clusterName, kms, vpcID string) {
	t.Helper()

	t.Log("try to get connection")
	c, err := connection.Get(client, connectionName, "")
	if c != nil && err == nil {
		t.Log("connection is found")
		tools.PrintResource(t, c)
		return
	}

	t.Log("create a connection")

	createOpts := connection.Connection{
		Name: connectionName,
		Type: "HOST",
		Config: connection.HOSTConfig{
			IP:          ec.Addresses[vpcID][0].Addr,
			Port:        "22",
			Username:    "linux",
			AgentName:   clusterName,
			KMSKey:      kms,
			KeyLocation: fmt.Sprintf("obs://%s/%s.pem", bucketName, keyPairName),
		},
	}

	err = connection.Create(client, createOpts)
	th.AssertNoErr(t, err)
}
