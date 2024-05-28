package v2

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/instance"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsInstanceV2LifeCycle(t *testing.T) {
	client, err := clients.NewDcsV2Client()
	th.AssertNoErr(t, err)

	dcsInstance := createDCSInstance(t, client)
	th.AssertEquals(t, dcsInstance.Capacity, float64(0))
	th.AssertEquals(t, dcsInstance.CapacityMinor, ".125")

	newName := tools.RandomString("dcs-instance-", 3)
	description := fmt.Sprintf("description for %s", newName)
	updateOpts := instance.ModifyInstanceOpt{
		InstanceId:  dcsInstance.InstanceID,
		Name:        newName,
		Description: pointerto.String(description),
		BackupPolicy: &instance.InstanceBackupPolicyOpts{
			BackupType: "auto",
			SaveDays:   1,
			PeriodicalBackupPlan: &instance.BackupPlan{
				PeriodType: "weekly",
				BeginAt:    "00:00-01:00",
				BackupAt:   []int{1, 2, 3, 4, 6, 7},
			},
		},
	}
	t.Logf("Attempting to update DCSv2 instance")
	err = instance.Update(client, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated DCSv2 instance")

	configList, err := instance.List(client, instance.ListDcsInstanceOpts{
		InstanceId: dcsInstance.InstanceID,
	})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updateOpts.Name, configList.Instances[0].Name)
	th.AssertDeepEquals(t, *updateOpts.Description, configList.Instances[0].Description)

	t.Logf("Attempting to resize DCSv2 instance")

	resizeOpts := instance.ResizeInstanceOpts{
		InstanceId:  dcsInstance.InstanceID,
		SpecCode:    "redis.ha.xu1.tiny.r2.256",
		NewCapacity: 0.25,
	}

	err = instance.Resize(client, resizeOpts)
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 100, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)

	t.Logf("Resized DCSv2 instance")

	ins, err := instance.Get(client, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)

	capacity, err := strconv.ParseFloat(ins.CapacityMinor, 64)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, ins.SpecCode, resizeOpts.SpecCode)
	th.AssertEquals(t, capacity, resizeOpts.NewCapacity)
}
