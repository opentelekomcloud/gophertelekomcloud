package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/configurations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestConfigurationsList(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	listOpts := configurations.ListOpts{}

	allPages, err := configurations.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	configs, err := configurations.ExtractConfigurations(allPages)
	th.AssertNoErr(t, err)

	for _, config := range configs {
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

	secGroupID := openstack.CreateSecurityGroup(t)
	defer openstack.DeleteSecurityGroup(t, secGroupID)

	defaultSGID := openstack.DefaultSecurityGroup(t)

	createOpts := configurations.CreateOpts{
		Name: asCreateName,
		InstanceConfig: configurations.InstanceConfigOpts{
			FlavorRef: "s3.xlarge.4",
			ImageRef:  imageID,
			Disk: []configurations.DiskOpts{
				{
					Size:       40,
					VolumeType: "SATA",
					DiskType:   "SYS",
				},
			},
			SSHKey: keyPairName,
			SecurityGroups: []configurations.SecurityGroupOpts{
				{
					ID: defaultSGID,
				},
				{
					ID: secGroupID,
				},
			},
		},
	}
	t.Logf("Attempting to create AutoScaling Configuration")
	configID, err := configurations.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Configuration: %s", configID)
	defer func() {
		t.Logf("Attempting to delete AutoScaling Configuration")
		err := configurations.Delete(client, configID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted AutoScaling Configuration: %s", configID)
	}()

	config, err := configurations.Get(client, configID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, config)
	th.AssertEquals(t, 2, len(config.InstanceConfig.SecurityGroups))
}
