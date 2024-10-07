package autoscaling

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/configurations"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateAutoScalingGroup(t *testing.T, client *golangsdk.ServiceClient, networkID, vpcID, asName string) string {
	defaultSGID := openstack.DefaultSecurityGroup(t)

	asCreateName := tools.RandomString("as-create-", 3)
	keyPairName := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	imageID := clients.EnvOS.GetEnv("IMAGE_ID")
	if keyPairName == "" || imageID == "" {
		t.Skip("OS_KEYPAIR_NAME or OS_IMAGE_ID env vars is missing but AS Configuration test requires")
	}

	configID := CreateASConfig(t, client, asCreateName, imageID, keyPairName)

	createOpts := groups.CreateOpts{
		Name:            asName,
		ConfigurationID: configID,
		Networks: []groups.ID{
			{
				ID: networkID,
			},
		},
		SecurityGroup: []groups.ID{
			{
				ID: defaultSGID,
			},
		},
		VpcID:                vpcID,
		IsDeletePublicip:     pointerto.Bool(true),
		DesireInstanceNumber: 1,
		MinInstanceNumber:    1,
		MaxInstanceNumber:    5,
	}
	t.Logf("Attempting to create AutoScaling Group")
	groupID, err := groups.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Group: %s", groupID)

	return groupID
}

func DeleteAutoScalingGroup(t *testing.T, client *golangsdk.ServiceClient, groupID string) {
	group, err := groups.Get(client, groupID)
	th.AssertNoErr(t, err)
	configID := group.ConfigurationID
	t.Logf("Attempting to delete AutoScaling Group")
	err = groups.Delete(client, groups.DeleteOpts{
		ScalingGroupId: groupID,
	})
	th.AssertNoErr(t, err)
	t.Logf("Deleted AutoScaling Group: %s", groupID)

	DeleteASConfig(t, client, configID)
}

func CreateASConfig(t *testing.T, client *golangsdk.ServiceClient, asCreateName string, imageID string, keyPairName string) string {
	defaultSGID := openstack.DefaultSecurityGroup(t)

	t.Logf("Attempting to create AutoScaling Configuration")
	configID, err := configurations.Create(client, configurations.CreateOpts{
		Name: asCreateName,
		InstanceConfig: configurations.InstanceConfigOpts{
			FlavorRef: "s3.medium.1",
			ImageRef:  imageID,
			Disk: []configurations.Disk{
				{
					Size:       40,
					VolumeType: "SSD",
					DiskType:   "SYS",
				},
			},
			SSHKey: keyPairName,
			SecurityGroups: []configurations.SecurityGroup{
				{
					ID: defaultSGID,
				},
			},
			Metadata: configurations.AdminPassMetadata{
				AdminPass: "Test1234",
			},
		},
	})
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Configuration: %s", configID)
	return configID
}

func DeleteASConfig(t *testing.T, client *golangsdk.ServiceClient, configID string) {
	t.Logf("Attempting to delete AutoScaling Configuration")
	err := configurations.Delete(client, configID)
	th.AssertNoErr(t, err)
	t.Logf("Deleted AutoScaling Configuration: %s", configID)
}

func GetNotificationTopicURN(topicName string) (string, error) {
	client, _ := clients.NewSmnV2Client()

	opts := topics.CreateOps{
		Name: topicName,
	}
	topic, err := topics.Create(client, opts).Extract()

	return topic.TopicUrn, err
}

func DeleteTopic(t *testing.T, topicURN string) {
	client, _ := clients.NewSmnV2Client()
	t.Logf("Attempting to Delete Topic: %s", topicURN)
	err := topics.Delete(client, topicURN).ExtractErr()
	if err != nil {
		t.Logf("Error while deleting the topic: %s", topicURN)
	}
	t.Logf("Deleted Topic: %s", topicURN)
}
