package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/alarms"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/resources"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarmResources(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create alarm rule")
	createOpts := alarms.CreateOpts{
		Name:      "test-alarm-resources-acc",
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

	t.Log("Attempting to list resources")
	listResp, err := resources.List(client, alarmId, resources.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count, 1)

	t.Log("Attempting to add resources")
	newResources := [][]resources.Dimension{
		{
			{
				Name:  "instance_id",
				Value: "00000000-0000-0000-0000-000000000002",
			},
		},
		{
			{
				Name:  "instance_id",
				Value: "00000000-0000-0000-0000-000000000003",
			},
		},
	}

	err = resources.Add(client, alarmId, resources.AddOpts{
		Resources: newResources,
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify resources were added")
	listResp, err = resources.List(client, alarmId, resources.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count, 3)

	t.Log("Attempting to delete resources")
	err = resources.Delete(client, alarmId, resources.DeleteOpts{
		Resources: newResources,
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify resources were deleted")
	listResp, err = resources.List(client, alarmId, resources.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count, 1)
}
