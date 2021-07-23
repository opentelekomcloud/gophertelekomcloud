package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/configs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsConfigLifeCycle(t *testing.T) {
	client, err := clients.NewDcsV1Client()
	th.AssertNoErr(t, err)

	dcsInstance := createDCSInstance(t, client)
	defer deleteDCSInstance(t, client, dcsInstance.InstanceID)

	updateOpts := configs.UpdateOpts{
		RedisConfigs: []configs.RedisConfig{
			{
				ParamID:    "1",
				ParamName:  "timeout",
				ParamValue: "100",
			},
		},
	}
	configList, err := configs.List(client, dcsInstance.InstanceID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to update DCSv1 configuration")
	err = configs.Update(client, dcsInstance.InstanceID, updateOpts).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Updated DCSv1 configuration")

	configList, err = configs.List(client, dcsInstance.InstanceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updateOpts.RedisConfigs[0].ParamID, configList.RedisConfigs[0].ParamID)
	th.AssertDeepEquals(t, updateOpts.RedisConfigs[0].ParamValue, configList.RedisConfigs[0].ParamValue)
	th.AssertDeepEquals(t, updateOpts.RedisConfigs[0].ParamName, configList.RedisConfigs[0].ParamName)
}
