package v1

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/configs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/whitelists"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsConfigLifeCycle(t *testing.T) {
	client, err := clients.NewDcsV1Client()
	th.AssertNoErr(t, err)

	dcsInstance := createDCSInstance(t, client)
	th.AssertEquals(t, dcsInstance.Capacity, 0)
	th.AssertEquals(t, dcsInstance.CapacityMinor, ".125")

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

	updateOpts := configs.UpdateOpts{
		RedisConfigs: []configs.RedisConfig{
			{
				ParamID:    "1",
				ParamName:  "timeout",
				ParamValue: "100",
			},
		},
	}
	t.Logf("Attempting to update DCSv1 configuration")
	err = configs.Update(client, dcsInstance.InstanceID, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated DCSv1 configuration")

	configList, err := configs.List(client, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updateOpts.RedisConfigs[0].ParamID, configList.RedisConfigs[0].ParamID)
	th.AssertDeepEquals(t, updateOpts.RedisConfigs[0].ParamValue, configList.RedisConfigs[0].ParamValue)
	th.AssertDeepEquals(t, updateOpts.RedisConfigs[0].ParamName, configList.RedisConfigs[0].ParamName)

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
}

// WaitForAWhitelistToBeRetrieved - wait until whitelist is retrieved
func WaitForAWhitelistToBeRetrieved(client *golangsdk.ServiceClient, id string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		wl, err := whitelists.Get(client, id)
		if err != nil {
			return false, fmt.Errorf("error retriving whitelist: %w", err)
		}
		if wl.InstanceID != "" {
			return true, nil
		}
		return false, nil
	})
}
