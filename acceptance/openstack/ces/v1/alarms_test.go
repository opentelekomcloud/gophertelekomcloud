package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/alarms"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarms(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	alarmsRes, err := alarms.ListAlarms(client, alarms.ListAlarmsOpts{
		Limit: 10,
		Order: "desc",
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, alarmsRes.MetaData.Count <= 10, true)

	newAlarm, err := alarms.CreateAlarm(client, alarms.CreateAlarmOpts{
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
		AlarmEnabled:       pointerto.Bool(false),
		AlarmActionEnabled: pointerto.Bool(false),
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = alarms.DeleteAlarm(client, newAlarm)
		th.AssertNoErr(t, err)
	})

	err = alarms.UpdateAlarmAction(client, newAlarm, alarms.ModifyAlarmActionRequest{
		AlarmEnabled: true,
	})
	th.AssertNoErr(t, err)

	showAlarm, err := alarms.ShowAlarm(client, newAlarm)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, showAlarm[0].AlarmEnabled, true)
}
