package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/alarms"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarms(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	lim := 10
	alarmsRes, err := alarms.ListAlarms(client, alarms.ListAlarmsRequest{
		Limit: &lim,
		Order: "desc",
	}).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, alarmsRes.MetaData.Count <= lim, true)

	expAlarm := alarmsRes.MetricAlarms[0]
	newAlarm, err := alarms.CreateAlarm(client, alarms.CreateAlarmRequest{
		AlarmName: expAlarm.AlarmName + "-copy",
		Metric:    expAlarm.Metric,
		Condition: expAlarm.Condition,
	}).Extract()
	th.AssertNoErr(t, err)

	err = alarms.UpdateAlarmAction(client, newAlarm.AlarmId, alarms.ModifyAlarmActionReq{
		AlarmEnabled: !expAlarm.AlarmEnabled,
	})
	th.AssertNoErr(t, err)

	showAlarm, err := alarms.ShowAlarm(client, newAlarm.AlarmId).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, showAlarm.MetricAlarms[0].AlarmName, expAlarm.AlarmName)
	th.AssertEquals(t, showAlarm.MetricAlarms[0].AlarmEnabled, !expAlarm.AlarmEnabled)

	err = alarms.DeleteAlarm(client, newAlarm.AlarmId)
	th.AssertNoErr(t, err)
}
