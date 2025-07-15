package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/configs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsV2ConfigLifeCycle(t *testing.T) {
	client, err := clients.NewDcsV2Client()
	th.AssertNoErr(t, err)

	dcsInstance := createDCSInstance(t, client)

	updateOpts := configs.ModifyConfigOpt{
		InstanceId: dcsInstance.InstanceID,
		RedisConfig: []configs.RedisConfigs{
			{
				ParamID:    "1",
				ParamName:  "timeout",
				ParamValue: "100",
			},
		},
	}
	t.Logf("Attempting to update DCSv2 configuration")
	err = configs.Update(client, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated DCSv2 configuration")

	configList, err := configs.Get(client, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updateOpts.RedisConfig[0].ParamID, configList.RedisConfigs[0].ParamID)
	th.AssertDeepEquals(t, updateOpts.RedisConfig[0].ParamValue, configList.RedisConfigs[0].ParamValue)
	th.AssertDeepEquals(t, updateOpts.RedisConfig[0].ParamName, configList.RedisConfigs[0].ParamName)
}
