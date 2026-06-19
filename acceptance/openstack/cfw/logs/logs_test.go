package logs

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/logs"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLogConfigLifecycle(t *testing.T) {
	ltsLogGroupID := clients.EnvOS.GetEnv("LTS_LOG_GROUP_ID")
	if ltsLogGroupID == "" {
		t.Skip("Test requires OS_LTS_LOG_GROUP_ID")
	}
	clientv1, err := clients.NewCFWV1Client()
	th.AssertNoErr(t, err)
	clientv2, err := clients.NewCFWV2Client()
	th.AssertNoErr(t, err)
	clientv3, err := clients.NewCFWV3Client()
	th.AssertNoErr(t, err)

	instanceName := tools.RandomString("test-acc-firewall-", 3)
	createOpts := managementv2.CreateOpts{
		Name: instanceName,
		Flavor: managementv2.CreateFlavor{
			Version: "standard",
		},
		ChargeInfo: managementv2.ChargeInfo{
			ChargeMode: "postPaid",
		},
	}
	createResp, err := managementv2.Create(clientv2, createOpts)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, common.WaitForJobCompleted(clientv3, 600, 5, createResp.JobID))
	instanceId := createResp.JobID
	t.Cleanup(func() {
		_, err = managementv2.Delete(clientv2, instanceId)
		th.AssertNoErr(t, err)
	})

	zero := 0
	one := 1
	_, err = logs.CreateLogConfig(clientv1, logs.LogConfigOpts{
		FWInstanceID:  instanceId,
		LtsEnable:     &one,
		LtsLogGroupID: ltsLogGroupID,
	})
	th.AssertNoErr(t, err)

	logConfig, err := logs.GetLogConfig(clientv1, logs.QueryParameters{
		FwInstanceID: instanceId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ltsLogGroupID, logConfig.LtsLogGroupID)
	th.AssertEquals(t, 1, logConfig.LtsEnable)

	_, err = logs.UpdateLogConfig(clientv1, logs.LogConfigOpts{
		FWInstanceID:  instanceId,
		LtsEnable:     &zero,
		LtsLogGroupID: ltsLogGroupID,
	})
	th.AssertNoErr(t, err)

	updatedLogConfig, err := logs.GetLogConfig(clientv1, logs.QueryParameters{
		FwInstanceID: instanceId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ltsLogGroupID, updatedLogConfig.LtsLogGroupID)
	th.AssertEquals(t, 0, updatedLogConfig.LtsEnable)
}
