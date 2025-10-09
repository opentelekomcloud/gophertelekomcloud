package v3

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTaurusInstanceLogs(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	createResp := createTaurusInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db")
		_, err = instance.Delete(client, createResp.Instance.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create taurus db read replica")

	createReplicaJobs, err := instance.CreateReplica(client, createResp.Instance.Id,
		instance.CreateReplicaOpts{
			Priorities: []int{1, 2},
		},
	)
	th.AssertNoErr(t, err)

	replicaJobs := strings.Split(*createReplicaJobs, ",")

	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, replicaJobs[0]))
	th.AssertNoErr(t, job.WaitForJobSuccess(client, 600, replicaJobs[1]))

	getResp, err := instance.Get(client, createResp.Instance.Id)
	th.AssertNoErr(t, err)

	nodeId := getResp.Nodes[0].Id

	endDate := time.Now().Format("2006-01-02T15:04:05-0700")
	startDate := time.Now().Add(-1 * time.Hour).Format("2006-01-02T15:04:05-0700")

	t.Logf("Attempting to list taurus db error logs")
	errLogs, err := logs.GetErrorLogs(client, logs.GetErrorLogsOpts{
		InstanceId: createResp.Instance.Id,
		NodeId:     nodeId,
		StartDate:  startDate,
		EndDate:    endDate,
	},
	)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, errLogs)

	t.Logf("Attempting to list taurus db slow logs")
	slowLogs, err := logs.GetSlowLogs(client, logs.GetSlowLogsOpts{
		InstanceId: createResp.Instance.Id,
		NodeId:     nodeId,
		StartDate:  startDate,
		EndDate:    endDate,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, slowLogs)
}
