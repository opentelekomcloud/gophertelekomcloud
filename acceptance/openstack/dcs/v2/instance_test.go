package v2

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/whitelists"
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

	t.Logf("Attempting to update whitelist configuration")
	err = whitelists.Put(client, dcsInstance.InstanceID, whitelists.WhitelistOpts{
		Enable: pointerto.Bool(true),
		Groups: []whitelists.WhitelistGroupOpts{
			{
				GroupName: "test-group-1",
				IPList: []string{
					"10.10.10.1", "10.10.10.2",
				},
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to update DCSv2 instance")
	err = instance.Update(client, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated DCSv2 instance")

	// No way to check if password was reset, refer to HC
	// t.Logf("Attempting to update DCSv2 password")
	// pwd, err := instance.UpdatePassword(client, instance.UpdatePasswordOpts{
	// 	InstanceId:  dcsInstance.InstanceID,
	// 	OldPassword: "Qwerty123!",
	// 	NewPassword: "Qwerty123!New-TEst!@",
	// })
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, pwd.Result, "success")
	//
	// err = waitForInstanceAvailable(client, 100, dcsInstance.InstanceID)
	// th.AssertNoErr(t, err)
	// t.Logf("Updated DCSv2 instance password")

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

	t.Logf("Retrieving whitelist configuration")
	err = WaitForAWhitelistToBeRetrieved(client, dcsInstance.InstanceID, 180)
	if err == nil {
		whitelistResp, err := whitelists.Get(client, dcsInstance.InstanceID)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, whitelistResp.InstanceID, dcsInstance.InstanceID)
		th.AssertEquals(t, whitelistResp.Groups[0].GroupName, "test-group-1")
	}

	t.Logf("Retrieving instance tags")
	instanceTags, err := tags.Get(client, "instances", dcsInstance.InstanceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(instanceTags), 2)
	th.AssertEquals(t, instanceTags[0].Key, "muh")
	th.AssertEquals(t, instanceTags[0].Value, "kuh")

	t.Logf("Updating instance tags")
	err = updateDcsTags(client, dcsInstance.InstanceID, instanceTags,
		[]tags.ResourceTag{
			{
				Key:   "muhUpdated",
				Value: "kuhUpdated",
			},
		})
	th.AssertNoErr(t, err)
	t.Logf("Retrieving updated instance tags")
	instanceTagsUpdated, err := tags.Get(client, "instances", dcsInstance.InstanceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(instanceTagsUpdated), 1)
	th.AssertEquals(t, instanceTagsUpdated[0].Key, "muhUpdated")
	th.AssertEquals(t, instanceTagsUpdated[0].Value, "kuhUpdated")

	_, err = instance.Restart(client, instance.ChangeInstanceStatusOpts{
		Instances: []string{
			dcsInstance.InstanceID,
		},
		Action: "restart",
	})
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 100, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)
}
