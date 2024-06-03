package v1_1

import (
	"fmt"
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/job"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const jobName = "testJob"

func TestDataArtsJobsLifecycle(t *testing.T) {
	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

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
	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

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
	th.AssertNoErr(t, err)

	defer f.Close()

	th.AssertNoErr(t, err)
	_, err = f.ReadFrom(storedFile)
	th.AssertNoErr(t, err)

	t.Log("import job")

	clientOBS, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

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
