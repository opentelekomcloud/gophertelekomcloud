package v1_1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1.1/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1.1/link"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const jobName = "testLink"

func TestDataArtsJobsLifecycle(t *testing.T) {
	t.Skip("skip this test before clarification")

	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	ak := os.Getenv("AWS_ACCESS_KEY")
	sk := os.Getenv("AWS_SECRET_KEY")

	client, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	c := getTestCluster(t, client)

	t.Log("create link")

	linkOpts := createLinkOpts(ak, sk)

	l, err := link.Create(client, c.Id, linkOpts, &link.CreateQuery{})
	th.AssertNoErr(t, err)

	t.Log("schedule link cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete link: %s", l.Name)
		err := link.Delete(client, c.Id, l.Name)
		th.AssertNoErr(t, err)
		t.Logf("link is deleted: %s", l.Name)
	})

	jobOpts := job.CreateSpecificOpts{
		Jobs: []job.Job{
			{
				FromConnectorName: "obs-connector",
				ToConfigValues: &job.ConfigValues{
					Configs: []*job.Config{{
						Inputs: []*job.Input{
							{
								Name:  "toJobConfig.schemaName",
								Value: "public",
							},
							{
								Name:  "toJobConfig.tablePreparation",
								Value: "DO_NOTHING",
							},
							{
								Name:  "toJobConfig.tableName",
								Value: "dailyactivity_merged",
							},
							{
								Name:  "toJobConfig.columnList",
								Value: "id&activitydate&totalsteps&totaldistance&trackerdistance&loggedactivitiesdistance&veryactivedistance&moderatelyactivedistance&lightactivedistance&sedentaryactivedistance&veryactiveminutes&fairlyactiveminutes&lightlyactiveminutes&sedentaryminutes&calories",
							},
							{
								Name:  "toJobConfig.useStageTable",
								Value: "false",
							},
							{
								Name:  "toJobConfig.shouldClearTable",
								Value: "false",
							},
							{
								Name:  "toJobConfig.beforeImportType",
								Value: "shouldClearTable",
							},
							{
								Name:  "toJobConfig.onConflict",
								Value: "EXCEPTION",
							},
						},
						Name: "toJobConfig",
					},
					},
					ExtendedConfigs: &job.ExtendedConfigs{
						Name:  "toJobConfig.extendedFields",
						Value: "eyLpbXBvcnRNb2RlIjoiQ09QWSIsImxvYWRlckNvbmN1cnLlbmN5IjoiMSL9",
					},
				},
				ToLinkName:         "dws-usecase",
				DriverConfigValues: nil,
				FromConfigValues:   nil,
				Name:               jobName,
			},
		},
	}

	j, err := job.CreateSpecific(client, c.Id, jobOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule job cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete job: %s", j.Name)
		err := job.Delete(client, c.Id, j.Name)
		th.AssertNoErr(t, err)
		t.Logf("job is deleted: %s", j.Name)
	})

	t.Log("get job")

	storedJob, err := job.Get(client, c.Id, j.Name, nil)
	tools.PrintResource(t, storedJob)
	th.AssertNoErr(t, err)
}
