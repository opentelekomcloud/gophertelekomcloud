package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/configurations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDdsConfigurationLifeCycle(t *testing.T) {
	instanceId := os.Getenv("DDS_INSTANCE_ID")

	if instanceId == "" {
		t.Skip("`DDS_INSTANCE_ID` need to be defined")
	}

	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get DDSv3 instance config")
	config, err := configurations.GetInstanceConfig(client, instanceId, configurations.ConfigOpts{EntityId: instanceId})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "mongodb", config.DatastoreName)

	t.Logf("Attempting to list DDSv3 template configurations")
	configList, err := configurations.List(client, configurations.ListConfigOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 29, configList.TotalCount)

	t.Logf("Attempting to get DDSv3 template configuration")
	getConfig, err := configurations.Get(client, configList.Configurations[0].ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "Default-DDS-3.2-Shard", getConfig.Name)

	t.Logf("Attempting to apply DDSv3 template configuration")
	_, err = configurations.Apply(client, configList.Configurations[11].ID,
		configurations.ApplyOpts{EntityIDs: []string{instanceId}})
	th.AssertNoErr(t, err)
}
