package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/antiddos/v1/antiddos"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	createOpt := antiddos.ConfigOpts{
		EnableL7:            true,
		TrafficPosId:        1,
		HttpRequestPosId:    2,
		CleaningAccessPosId: 3,
		AppTypeId:           1,
	}

	actual, err := antiddos.CreateDefaultConfig(client.ServiceClient(), createOpt)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	actual, err := antiddos.DeleteDefaultConfig(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &DeleteResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.ShowDDos(client.ServiceClient(), floatingIpId)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	updateOpt := antiddos.ConfigOpts{
		EnableL7:            true,
		TrafficPosId:        1,
		HttpRequestPosId:    2,
		CleaningAccessPosId: 3,
		AppTypeId:           1,
	}

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.UpdateDDos(client.ServiceClient(), floatingIpId, updateOpt)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestListStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListStatusSuccessfully(t)

	listOpt := antiddos.ListDDosStatusOpts{
		Limit:  2,
		Offset: 1,
		Status: "notConfig",
		Ip:     "49.",
	}

	actual, err := antiddos.ListDDosStatus(client.ServiceClient(), listOpt)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ListStatusResponse, actual.DDosStatus)
}

func TestListConfigs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListConfigsSuccessfully(t)

	actual, err := antiddos.ListNewConfigs(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListConfigsResponse, actual)
}

func TestWeeklyReport(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWeeklyReportSuccessfully(t)

	actual, err := antiddos.ListWeeklyReports(client.ServiceClient(), responsePeriodTime)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &WeeklyReportResponse, actual)
}

func TestListLogs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLogsSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.ListDailyLogs(client.ServiceClient(), floatingIpId, antiddos.ListDailyLogsOps{
		Limit:   2,
		Offset:  1,
		SortDir: "asc",
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ListLogsResponse, actual.Logs)
}

func TestGetStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetStatusSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.ShowDDosStatus(client.ServiceClient(), floatingIpId)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, "normal", actual)
}

func TestDailyReport(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDailyReportSuccessfully(t)

	floatingIpId := "82abaa86-8518-47db-8d63-ddf152824635"
	actual, err := antiddos.ListDailyReport(client.ServiceClient(), floatingIpId)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, DailyReportResponse, actual)
}

func TestGetTask(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetTaskSuccessfully(t)

	actual, err := antiddos.ShowNewTaskStatus(client.ServiceClient(), antiddos.ShowNewTaskStatusOpts{
		TaskId: "4a4fefe7-34a1-40e2-a87c-16932af3ac4a",
	})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetTaskResponse, actual)
}
