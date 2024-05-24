package v1_1

import (
	"fmt"
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	v1_1 "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/dataarts/v1.1"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/script"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const jobName = "testJob"

func TestDataArtsJobsLifecycle(t *testing.T) {
	client, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	workspace := ""

	t.Log("create a job")

	createOpts := &job.Job{
		Name: jobName,
		Schedule: job.Schedule{
			Type: "EXECUTE_ONCE",
		},
		ProcessType: "BATCH",
	}

	err = job.Create(client, *createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule job cleanup")
	t.Cleanup(func() {
		jobCleanup(t, client)
	})

	t.Log("get job")

	storedJob, err := job.Get(client, jobName, workspace)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedJob)

	storedJob.ProcessType = "REAL_TIME"

	err = job.Update(client, jobName, *storedJob)
	th.AssertNoErr(t, err)

	t.Log("should wait 5 seconds")
	time.Sleep(5 * time.Second)

	t.Log("get job")

	storedJob, err = job.Get(client, jobName, workspace)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedJob)
	th.CheckEquals(t, "REAL_TIME", storedJob.ProcessType)

	// TODO: this api is now working
	// t.Log("get jobs list")
	// storedJobs, err := job.List(client, job.ListOpts{}, workspace)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, 1, storedJobs.Total)
	// th.AssertEquals(t, 1, len(storedJobs.Jobs))
	// th.AssertEquals(t, jobName, storedJobs.Jobs[0].Name)

	// // TODO: this api is now working
	// t.Log("get job file destination")
	// storedFile, err := job.GetJobFile(client, nil)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, 1, len(storedFile.Jobs))
	// th.AssertEquals(t, 1, len(storedFile.Scripts))
}
func TestDataArtsJobsImportExport(t *testing.T) {
	client, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	t.Log("create a job")

	createOpts := &job.Job{
		Name: jobName,
		Schedule: job.Schedule{
			Type: "EXECUTE_ONCE",
		},
		ProcessType: "BATCH",
	}

	err = job.Create(client, *createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule job cleanup")
	t.Cleanup(func() {
		jobCleanup(t, client)
	})

	t.Log("export job")
	storedFile, err := job.ExportJob(client, jobName, "")
	th.AssertNoErr(t, err)
	defer storedFile.Close()

	f, err := os.Create("jobTest.zip")
	defer f.Close()

	th.AssertNoErr(t, err)
	_, err = f.ReadFrom(storedFile)
	th.AssertNoErr(t, err)

	t.Log("import job")

	clientOBS, err := clients.NewOBSClient()
	prepareTestBucket(t, clientOBS)
	t.Cleanup(func() {
		cleanupBucket(t, clientOBS)
	})

	uploadFile(t, clientOBS, f.Name(), f)

	_, err = job.ImportJob(client, job.ImportReq{
		Path: fmt.Sprintf("obs://%s/%s", bucketName, "jobTest.zip"),
		JobsParam: []*job.JobParamImported{
			{
				Name: "jobTestImported",
			},
		},
	})

	th.AssertNoErr(t, err)
}

func jobCleanup(t *testing.T, client *golangsdk.ServiceClient) {
	t.Helper()
	t.Logf("attempting to delete job: %s", jobName)
	err := job.Delete(client, jobName, nil)
	th.AssertNoErr(t, err)
	t.Logf("job is deleted: %s", jobName)
}

func TestDataArtsJobExecution(t *testing.T) {
	// if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
	// 	t.Skip("too slow to run in zuul")
	// }

	workspace := os.Getenv("AWS_WORKSPACE")

	client, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	client11, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	cluster := v1_1.GetTestCluster(t, client11)
	tools.PrintResource(t, cluster)

	ec, clientEC := getECInstance(t)
	defer openstack.DeleteCloudServer(t, clientEC, ec.ID)

	script := getScript(t)

	// client, err := clients.NewDataArtsV1Client()
	// th.AssertNoErr(t, err)
	//
	// workspace := ""
	//
	// t.Log("create a job")
	//
	// createOpts := &job.Job{
	// 	Name: jobName,
	// 	Schedule: job.Schedule{
	// 		Type: "EXECUTE_ONCE",
	// 	},
	// 	ProcessType: "BATCH",
	// }
	//
	// err = job.Create(client, *createOpts)
	// th.AssertNoErr(t, err)
	//
	// t.Log("schedule job cleanup")
	// t.Cleanup(func() {
	// 	jobCleanup(t, client)
	// })
}

func getECInstance(t *testing.T) (*cloudservers.CloudServer, *golangsdk.ServiceClient) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	// Get ECSv1 createOpts
	createOpts := openstack.GetCloudServerCreateOpts(t)

	// Check ECSv1 createOpts
	openstack.DryRunCloudServerConfig(t, client, createOpts)
	t.Logf("CreateOpts are ok for creating a cloudServer")

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, client, createOpts)
	// defer openstack.DeleteCloudServer(t, client, ecs.ID)
	return ecs, client
}

func getScript(t *testing.T, client *golangsdk.ServiceClient) {

	t.Log("create a script")

	createOpts := script.Script{
		Name:           scriptName,
		Type:           "Shell",
		Content:        "echo 123456",
		ConnectionName: "anyConnection",
	}

	err := script.Create(client, createOpts)
	th.AssertNoErr(t, err)
}
