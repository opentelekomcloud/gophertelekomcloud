package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
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

func TestAlarmsCreateEventSysWithEmptyResources(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	alarmName := tools.RandomString("test-alarm-event-sys-", 4)

	t.Log("Attempting to create EVENT.SYS alarm rule with empty resources")
	createOpts := alarms.CreateOpts{
		Name:      alarmName,
		Namespace: "SYS.ECS",
		Type:      "EVENT.SYS",
		Resources: [][]alarms.Dimension{},
		Policies: []alarms.Policy{
			{
				MetricName:         "stopServer",
				Period:             0,
				Filter:             "average",
				ComparisonOperator: ">=",
				Value:              1,
				Unit:               "count",
				Count:              1,
				SuppressDuration:   0,
				Level:              2,
			},
		},
		NotificationEnabled: pointerto.Bool(false),
		Enabled:             pointerto.Bool(true),
	}

	alarmId, err := alarms.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Log("Attempting to delete EVENT.SYS alarm rule")
		_, err := alarms.Delete(client, alarms.DeleteOpts{
			AlarmIds: []string{alarmId},
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to verify EVENT.SYS alarm rule")
	listResp, err := alarms.List(client, alarms.ListOpts{
		AlarmId: alarmId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count, 1)
	th.AssertEquals(t, listResp.Alarms[0].Name, alarmName)
	th.AssertEquals(t, listResp.Alarms[0].Namespace, "SYS.ECS")
	th.AssertEquals(t, listResp.Alarms[0].Type, "EVENT.SYS")
	th.AssertEquals(t, listResp.Alarms[0].Enabled, true)
	th.AssertEquals(t, listResp.Alarms[0].NotificationEnabled, false)
	th.AssertEquals(t, len(listResp.Alarms[0].Resources), 1)
	th.AssertEquals(t, len(listResp.Alarms[0].Policies), 1)
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].MetricName, "stopServer")
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].Period, 0)
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].Filter, "average")
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].ComparisonOperator, ">=")
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].Value, float64(1))
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].Unit, "count")
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].Count, 1)
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].SuppressDuration, 0)
	th.AssertEquals(t, listResp.Alarms[0].Policies[0].Level, 2)
}
