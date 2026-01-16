package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/records"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarmRecordsList(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to list alarm records")
	listResp, err := records.List(client, records.ListOpts{
		Limit: 10,
	})
	th.AssertNoErr(t, err)

	t.Logf("Found %d alarm records", listResp.Count)
	for _, record := range listResp.AlarmHistories {
		t.Logf("Record ID: %s, Name: %s, Status: %s, Level: %d",
			record.RecordId, record.Name, record.Status, record.Level)
	}
}

func TestAlarmRecordsListWithFilters(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to list alarm records with status filter")
	listResp, err := records.List(client, records.ListOpts{
		Status: "alarm",
		Limit:  10,
	})
	th.AssertNoErr(t, err)

	t.Logf("Found %d alarm records with status 'alarm'", listResp.Count)
	for _, record := range listResp.AlarmHistories {
		th.AssertEquals(t, record.Status, "alarm")
	}
}
