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

	allConfigurations, err := configurations.List(client).Extract()
	th.AssertNoErr(t, err)

	for _, configuration := range allConfigurations {
		tools.PrintResource(t, configuration)
	}
}

func TestConfigurationsLifecycle(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	config := createRDSConfiguration(t, client)
	defer deleteRDSConfiguration(t, client, config.ID)

	tools.PrintResource(t, config)

	updateRDSConfiguration(t, client, config.ID)

	newConfig, err := configurations.Get(client, config.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newConfig)
}

func TestConfigurationsApply(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	config := createRDSConfiguration(t, client)
	defer deleteRDSConfiguration(t, client, config.ID)

	// Create RDSv3 instance
	rds := createRDS(t, client, cc.RegionName)
	defer deleteRDS(t, client, rds.Id)

	t.Logf("Attempting to apply template config %s to instance %s", config.ID, rds.Id)
	configApplyOpts := configurations.ApplyOpts{
		InstanceIDs: []string{rds.Id},
	}
	applyResult, err := configurations.Apply(client, config.ID, configApplyOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Template config applied")

	tools.PrintResource(t, applyResult)
}
