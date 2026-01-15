package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/alarms"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/policies"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarmPolicies(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create alarm rule")
	createOpts := alarms.CreateOpts{
		Name:      "test-alarm-policies-acc",
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

	t.Log("Attempting to list policies")
	listResp, err := policies.List(client, alarmId, policies.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count, 1)
	th.AssertEquals(t, listResp.Policies[0].MetricName, "cpu_util")
	th.AssertEquals(t, listResp.Policies[0].Value, float64(80))

	t.Log("Attempting to update policies")
	_, err = policies.Update(client, alarmId, policies.UpdateOpts{
		Policies: []policies.Policy{
			{
				MetricName:         "cpu_util",
				Period:             300,
				Filter:             "average",
				ComparisonOperator: ">=",
				Value:              90,
				Count:              5,
				SuppressDuration:   600,
				Level:              1,
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify policies were updated")
	listResp, err = policies.List(client, alarmId, policies.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Policies[0].ComparisonOperator, ">=")
	th.AssertEquals(t, listResp.Policies[0].Value, float64(90))
	th.AssertEquals(t, listResp.Policies[0].Count, 5)
	th.AssertEquals(t, listResp.Policies[0].Level, 1)
}
