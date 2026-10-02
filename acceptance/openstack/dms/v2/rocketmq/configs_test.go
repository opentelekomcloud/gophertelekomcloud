package rocketmq

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/configs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const fileReservedTime = "fileReservedTime"

func TestRocketMQConfigs(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	listed, err := configs.List(client, instanceID, configs.ListOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listed)
	th.AssertEquals(t, true, listed.Total > 0)

	original := findConfig(t, listed.Configs, fileReservedTime)
	th.AssertEquals(t, "dynamic", original.ConfigType)
	th.AssertEquals(t, "integer", original.ValueType)

	target := "72"
	if original.Value == target {
		target = "73"
	}

	t.Logf("Attempting to set %s of RocketMQ instance %s to %s", fileReservedTime, instanceID, target)
	err = configs.Update(client, instanceID, configs.UpdateOpts{
		Configs: []configs.UpdateConfig{{Name: fileReservedTime, Value: target}},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to restore %s of RocketMQ instance %s to %s", fileReservedTime, instanceID, original.Value)
		err := configs.Update(client, instanceID, configs.UpdateOpts{
			Configs: []configs.UpdateConfig{{Name: fileReservedTime, Value: original.Value}},
		})
		th.AssertNoErr(t, err)
	})

	listed, err = configs.List(client, instanceID, configs.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, target, findConfig(t, listed.Configs, fileReservedTime).Value)
}

func findConfig(t *testing.T, configList []configs.Config, name string) configs.Config {
	for _, c := range configList {
		if c.Name == name {
			return c
		}
	}
	t.Fatalf("no RocketMQ configuration %s", name)
	return configs.Config{}
}
