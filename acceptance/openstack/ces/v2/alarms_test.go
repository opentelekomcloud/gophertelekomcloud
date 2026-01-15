package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/alarms"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarmsCRUD(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create alarm rule")
	createOpts := alarms.CreateOpts{
		Name:      "test-alarm-v2-acc",
		Namespace: "SYS.ECS",
		Type:      "MULTI_INSTANCE",
		Resources: [][]alarms.Dimension{
			{
				{
					Name:  "instance_id",
					Value: "00000000-0000-0000-0000-000000000001",
				},
			},
		},
		Policies: []alarms.Policy{
			{
				MetricName:         "cpu_util",
				Period:             300,
				Filter:             "average",
				ComparisonOperator: ">",
				Value:              80,
				Count:              3,
				SuppressDuration:   300,
				Level:              2,
			},
		},
		Enabled:             pointerto.Bool(false),
		NotificationEnabled: pointerto.Bool(false),
	}

	alarmId, err := alarms.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Log("Attempting to delete alarm rule")
		_, err := alarms.Delete(client, alarms.DeleteOpts{
			AlarmIds: []string{alarmId},
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to get alarm rule")
	listResp, err := alarms.List(client, alarms.ListOpts{
		AlarmId: alarmId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count, 1)
	th.AssertEquals(t, listResp.Alarms[0].Name, "test-alarm-v2-acc")
	th.AssertEquals(t, listResp.Alarms[0].Enabled, false)

	t.Log("Attempting to enable alarm rule")
	_, err = alarms.Action(client, alarms.ActionOpts{
		AlarmIds:     []string{alarmId},
		AlarmEnabled: true,
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify alarm rule is enabled")
	listResp, err = alarms.List(client, alarms.ListOpts{
		AlarmId: alarmId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Alarms[0].Enabled, true)

	t.Log("Attempting to disable alarm rule")
	_, err = alarms.Action(client, alarms.ActionOpts{
		AlarmIds:     []string{alarmId},
		AlarmEnabled: false,
	})
	th.AssertNoErr(t, err)
}
