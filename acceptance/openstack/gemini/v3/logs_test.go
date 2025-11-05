package v3

import (
	"os"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGeminiInstanceLogs(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	createResp := createGeminiInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete gemini db")
		_, err = instance.Delete(client, createResp.Id)
		th.AssertNoErr(t, err)
	})

	endDate := time.Now().Format("2006-01-02T15:04:05-0700")
	startDate := time.Now().Add(-1 * time.Hour).Format("2006-01-02T15:04:05-0700")

	logsOpts := logs.GetSlowLogsOpts{
		InstanceId: createResp.Id,
		StartDate:  startDate,
		EndDate:    endDate,
	}

	slowLogs, err := logs.GetSlowLogs(client, logsOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, slowLogs)
}
