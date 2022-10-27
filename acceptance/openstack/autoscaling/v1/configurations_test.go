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

	secGroupID := openstack.CreateSecurityGroup(t)
	t.Cleanup(func() {
		openstack.DeleteSecurityGroup(t, secGroupID)
	})

	t.Logf("Attempting to create AutoScaling Configuration")
	configID, err := configurations.Create(client, configurations.CreateOpts{
		Name: asCreateName,
		InstanceConfig: configurations.InstanceConfigOpts{
			FlavorRef: "s3.xlarge.4",
			ImageRef:  imageID,
			Disk: []configurations.Disk{
				{
					Size:       40,
					VolumeType: "SATA",
					DiskType:   "SYS",
					Metadata: configurations.SystemMetadata{
						SystemEncrypted: "0",
					},
				},
			},
			SSHKey: keyPairName,
			SecurityGroups: []configurations.SecurityGroup{
				{
					ID: openstack.DefaultSecurityGroup(t),
				},
				{
					ID: secGroupID,
				},
			},
			Metadata: configurations.AdminPassMetadata{
				AdminPass: "Test1234",
			},
		},
	})
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Configuration: %s", configID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete AutoScaling Configuration")
		err := configurations.Delete(client, configID)
		th.AssertNoErr(t, err)
		t.Logf("Deleted AutoScaling Configuration: %s", configID)
	})

	config, err := configurations.Get(client, configID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, config)
	th.AssertEquals(t, 2, len(config.InstanceConfig.SecurityGroups))
}
