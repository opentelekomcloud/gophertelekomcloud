package v3

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	networking "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/networking/v1"
	rds "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/rds/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/drs/v3/public"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDrsTaskLifecycle(t *testing.T) {
	// if os.Getenv("RUN_DRS_LIFECYCLE") == "" {
	// 	t.Skip("too slow to run in zuul")
	// }
	subnetId := os.Getenv("OS_SUBNET_ID")

	client, err := clients.NewDrsV3Client()
	th.AssertNoErr(t, err)

	rdsClient, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	netClient, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	instanceName := tools.RandomString("drs-task-", 8)
	netType := "eip"
	dbType := "postgresql"

	rdsChan1 := make(chan *instances.Instance)
	rdsChan2 := make(chan *instances.Instance)

	go func() {
		defer close(rdsChan1)
		source := rds.CreateRDS(t, rdsClient, cc.RegionName)
		rdsChan1 <- source
	}()

	go func() {
		defer close(rdsChan2)
		target := rds.CreateRDS(t, rdsClient, cc.RegionName)
		rdsChan2 <- target
	}()

	rdsSource := <-rdsChan1
	rdsTarget := <-rdsChan2

	t.Cleanup(func() { rds.DeleteRDS(t, rdsClient, rdsSource.Id) })
	t.Cleanup(func() { rds.DeleteRDS(t, rdsClient, rdsTarget.Id) })

	elasticIP := networking.CreateEip(t, netClient, 100)
	t.Cleanup(func() {
		networking.DeleteEip(t, netClient, elasticIP.ID)
	})

	t.Log("AttachEip")

	err = instances.AttachEip(rdsClient, instances.AttachEipOpts{
		InstanceId: rdsSource.Id,
		PublicIp:   elasticIP.PublicAddress,
		PublicIpId: elasticIP.ID,
		IsBind:     pointerto.Bool(true),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = instances.AttachEip(rdsClient, instances.AttachEipOpts{
			InstanceId: rdsSource.Id,
			IsBind:     pointerto.Bool(false),
		})
		th.AssertNoErr(t, err)
	})

	createJobOpts := public.BatchCreateTaskOpts{
		Jobs: []public.CreateJobOpts{{
			Name:              instanceName,
			DbUseType:         "sync",
			EngineType:        dbType,
			JobDirection:      "up",
			CustomizeSubnetId: subnetId,
			NetType:           netType,
			BindEip:           pointerto.Bool(true),
			NodeType:          "high",
			TaskType:          "FULL_TRANS",
			SourceEndpoint: public.Endpoint{
				DbType: "postgresql",
				Region: cc.RegionName,
				InstId: rdsSource.Id,
			},
			TargetEndpoint: public.Endpoint{
				DbType: "postgresql",
				Region: cc.RegionName,
				InstId: rdsTarget.Id,
			},
		}},
	}

	task, err := public.BatchCreateTasks(client, createJobOpts)
	th.AssertNoErr(t, err)

	defer func() {
		deleteJob, err := public.BatchDeleteTasks(client, public.BatchDeleteTasksOpts{
			Jobs: []public.DeleteJobReq{{
				JobId:      task.Results[0].Id,
				DeleteType: "force_terminate",
			}},
		})
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, waitForTaskComplete(client, deleteJob.Results[0].Id, "RELEASE_RESOURCE_COMPLETE", 600))
		_, err = public.BatchDeleteTasks(client, public.BatchDeleteTasksOpts{
			Jobs: []public.DeleteJobReq{{
				JobId:      task.Results[0].Id,
				DeleteType: "delete",
			}},
		})
		th.AssertNoErr(t, err)
	}()

	th.AssertNoErr(t, waitForTaskComplete(client, task.Results[0].Id, "CONFIGURATION", 600))

	taskList, err := public.BatchListTaskDetails(client, public.BatchQueryTaskOpts{
		Jobs: []string{
			task.Results[0].Id,
		},
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, taskList)

	taskStatus, err := public.BatchListTaskStatus(client, public.BatchQueryTaskOpts{
		Jobs: []string{
			task.Results[0].Id,
		},
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, taskStatus)

	testConnection, err := public.BatchTestConnections(client, public.BatchTestConnectionOpts{
		Jobs: []public.TestEndPoint{
			{
				Id:           task.Results[0].Id,
				NetType:      netType,
				DbType:       dbType,
				DbPort:       8635,
				Ip:           elasticIP.PublicAddress,
				DbUser:       "root",
				DbPassword:   "acc-test-password1!",
				EndPointType: "so",
			},
		},
	},
	)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, testConnection.Results[0].Status, "failed")

	taskModify, err := public.BatchUpdateTask(client, public.BatchModifyJobOpts{
		Jobs: []public.ModifyJobReq{
			{
				JobId:       task.Results[0].Id,
				Name:        "updated-task",
				Description: "new_description",
			},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, taskModify.Results[0].Status, "success")

	setSpeed, err := public.BatchSetSpeed(client, public.BatchLimitSpeedOpts{
		SpeedLimits: []public.LimitSpeedReq{
			{
				JobId: task.Results[0].Id,
				SpeedLimit: []public.SpeedLimitInfo{
					{
						Speed: "15",
						Begin: "16:00",
						End:   "15:59",
					},
				},
			},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, setSpeed.Results[0].Status, "success")

	startTask, err := public.BatchStartTasks(client, public.BatchStartJobOpts{
		Jobs: []public.StartInfo{
			{
				JobId: task.Results[0].Id,
			},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, startTask.Results[0].Status, "success")
}

func waitForTaskComplete(c *golangsdk.ServiceClient, id, status string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		taskList, err := public.BatchListTaskDetails(c, public.BatchQueryTaskOpts{
			Jobs: []string{
				id,
			},
		})
		if err != nil {
			return false, err
		}

		if taskList.Results[0].Status == status {
			return true, nil
		}

		return false, nil
	})
}
