package v1_1

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/job"
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
		t.Logf("attempting to delete job: %s", jobName)
		err := job.Delete(client, jobName, nil)
		th.AssertNoErr(t, err)
		t.Logf("job is deleted: %s", jobName)
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
}
