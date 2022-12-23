package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/configurations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestConfigurationsList(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	allConfigurations, err := configurations.List(client)
	th.AssertNoErr(t, err)

	for _, configuration := range allConfigurations {
		tools.PrintResource(t, configuration)
	}
}

func TestConfigurationsLifecycle(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	config := createRDSConfiguration(t, client)
	t.Cleanup(func() { deleteRDSConfiguration(t, client, config.ID) })

	tools.PrintResource(t, config)

	updateRDSConfiguration(t, client, config.ID)

	newConfig, err := configurations.Get(client, config.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newConfig)
}

func TestConfigurations(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	config := createRDSConfiguration(t, client)
	t.Cleanup(func() { deleteRDSConfiguration(t, client, config.ID) })

	// Create RDSv3 instance
	rds := createRDS(t, client, cc.RegionName)
	t.Cleanup(func() { deleteRDS(t, client, rds.Id) })

	t.Logf("Attempting to apply template config %s to instance %s", config.ID, rds.Id)
	configApplyOpts := configurations.ApplyOpts{
		ConfigId:    config.ID,
		InstanceIDs: []string{rds.Id},
	}
	applyResult, err := configurations.Apply(client, configApplyOpts)
	th.AssertNoErr(t, err)
	t.Logf("Template config applied")

	tools.PrintResource(t, applyResult)

	instanceConfig, err := configurations.GetForInstance(client, rds.Id)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, config.DatastoreName, instanceConfig.DatastoreName)
	th.CheckEquals(t, config.DatastoreVersionName, instanceConfig.DatastoreVersionName)
	if len(instanceConfig.Parameters) == 0 {
		t.Errorf("instance config has empty parameter list")
	}

	opts := configurations.UpdateInstanceConfigurationOpts{
		InstanceId: rds.Id,
		Values: map[string]interface{}{
			"max_connections": "37",
			"autocommit":      "OFF",
		}}
	result, err := configurations.UpdateInstanceConfiguration(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, result.RestartRequired)
}
