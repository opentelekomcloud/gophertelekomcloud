package rocketmq

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/diagnosis"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const diagnosisReportNotFound = "DMS.00500975"

func TestRocketMQDiagnosis(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)
	groupName := createRocketMQGroup(t, client, instanceID)

	t.Logf("Attempting to create RocketMQ diagnosis task for consumer group: %s", groupName)
	created, err := diagnosis.Create(client, instanceID, diagnosis.CreateOpts{
		GroupName: groupName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, created.ReportID != "")
	reportID := created.ReportID

	t.Cleanup(func() {
		_, err := diagnosis.BatchDelete(client, instanceID, diagnosis.BatchDeleteOpts{
			ReportIDList: []string{reportID},
		})
		if err != nil && !isNotFound(err, diagnosisReportNotFound) {
			t.Errorf("failed to delete RocketMQ diagnosis report %s: %s", reportID, err)
		}
	})

	var report *diagnosis.Report
	th.AssertNoErr(t, golangsdk.WaitFor(300, func() (bool, error) {
		report, err = diagnosis.Get(client, reportID)
		if err != nil {
			return false, err
		}
		return report.Status != "diagnosing", nil
	}))
	tools.PrintResource(t, report)
	th.AssertEquals(t, reportID, report.ReportID)
	th.AssertEquals(t, groupName, report.GroupName)
	th.AssertEquals(t, "finished", report.Status)
	th.AssertEquals(t, 0, report.ConsumerNums)
	th.AssertEquals(t, false, report.Online)
	th.AssertEquals(t, true, report.CreatedAt > 0)

	listed, err := diagnosis.List(client, instanceID, diagnosis.ListOpts{Limit: 10})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, listed.TotalNum)
	th.AssertEquals(t, 1, len(listed.Reports))
	th.AssertEquals(t, reportID, listed.Reports[0].ReportID)
	th.AssertEquals(t, groupName, listed.Reports[0].GroupName)
	th.AssertEquals(t, "finished", listed.Reports[0].Status)

	// Node reports with stack IDs are only produced while consumers are
	// connected to the group, which the test has no client for.
	th.AssertEquals(t, 0, len(report.NodeReports))
	_, err = diagnosis.GetStack(client, tools.RandomString("stack-acc-", 8))
	th.AssertEquals(t, true, isNotFound(err, diagnosisReportNotFound))

	t.Logf("Attempting to delete RocketMQ diagnosis report: %s", reportID)
	deleted, err := diagnosis.BatchDelete(client, instanceID, diagnosis.BatchDeleteOpts{
		ReportIDList: []string{reportID},
	})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []string{reportID}, deleted.ReportIDList)

	_, err = diagnosis.Get(client, reportID)
	th.AssertEquals(t, true, isNotFound(err, diagnosisReportNotFound))

	listed, err = diagnosis.List(client, instanceID, diagnosis.ListOpts{Limit: 10})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, listed.TotalNum)
}

// isNotFound reports whether err is a 404 response with the given DMS error code.
func isNotFound(err error, code string) bool {
	e, ok := err.(golangsdk.ErrDefault404)
	return ok && strings.Contains(string(e.Body), code)
}
