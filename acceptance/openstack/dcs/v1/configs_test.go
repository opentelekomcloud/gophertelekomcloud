package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
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

	enabled := true

	whitelistOpts := whitelists.WhitelistOpts{
		Enable: &enabled,
		Groups: []whitelists.WhitelistGroupOpts{
			{
				GroupName: "test-group-1",
				IPList: []string{
					"10.10.10.1", "10.10.10.2",
				},
			},
		},
	}

	t.Logf("Attempting to update whitelist configuration")
	_ = whitelists.Put(client, dcsInstance.InstanceID, whitelistOpts)
	th.AssertNoErr(t, err)

	t.Logf("Retrieving whitelist configuration")
	whitelistResp, err := whitelists.Get(client, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, whitelistResp.InstanceID, dcsInstance.InstanceID)
	th.AssertDeepEquals(t, whitelistResp.Groups[0].GroupName, "test-group-1")
}
