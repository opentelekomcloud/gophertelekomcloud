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

	f := false
	newAlarm, err := alarms.CreateAlarm(client, alarms.CreateAlarmRequest{
		AlarmName: "alarm-acc-test",
		Metric: alarms.MetricForAlarm{
			Namespace:  "SYS.VPC",
			MetricName: "upstream_bandwidth",
			Dimensions: []alarms.MetricsDimension{
				{
					Name:  "bandwidth_id",
					Value: "026c495c-fake-test-8b11-a113ba530d11",
				},
			},
		},
		Condition: alarms.Condition{
			ComparisonOperator: ">=",
			Count:              3,
			Filter:             "average",
			Period:             300,
			Value:              4000000,
		},
		AlarmEnabled:       &f,
		AlarmActionEnabled: &f,
	}).Extract()
	th.AssertNoErr(t, err)

	err = alarms.UpdateAlarmAction(client, newAlarm.AlarmId, alarms.ModifyAlarmActionReq{
		AlarmEnabled: true,
	})
	th.AssertNoErr(t, err)

	showAlarm, err := alarms.ShowAlarm(client, newAlarm.AlarmId).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, showAlarm.MetricAlarms[0].AlarmEnabled, true)

	err = alarms.DeleteAlarm(client, newAlarm.AlarmId)
	th.AssertNoErr(t, err)
}
