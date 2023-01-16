package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/configurations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestConfigurationsList(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	listOpts := configurations.ListOpts{}

	configs, err := configurations.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, config := range configs.ScalingConfigurations {
		tools.PrintResource(t, config)
	}
}

func TestConfigurationsLifecycle(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	asCreateName := tools.RandomString("as-create-", 3)
	keyPairName := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	imageID := clients.EnvOS.GetEnv("IMAGE_ID")
	if keyPairName == "" || imageID == "" {
		t.Skip("OS_KEYPAIR_NAME or OS_IMAGE_ID env vars is missing but AS Configuration test requires")
	}

	configID := autoscaling.CreateASConfig(t, client, asCreateName, imageID, keyPairName)

	t.Cleanup(func() {
		autoscaling.DeleteASConfig(t, client, configID)
	})

	config, err := configurations.Get(client, configID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, config)
	th.AssertEquals(t, 1, len(config.InstanceConfig.SecurityGroups))
}
