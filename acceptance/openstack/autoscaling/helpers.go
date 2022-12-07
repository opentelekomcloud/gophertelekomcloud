package autoscaling

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/configurations"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateAutoScalingGroup(t *testing.T, client *golangsdk.ServiceClient, networkID, vpcID, asName string) string {
	defaultSGID := openstack.DefaultSecurityGroup(t)

	createOpts := groups.CreateOpts{
		Name: asName,
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
	t.Logf("Attempting to delete AutoScaling Group")
	err := groups.Delete(client, groups.DeleteOpts{
		ScalingGroupId: groupID,
	})
	th.AssertNoErr(t, err)
	t.Logf("Deleted AutoScaling Group: %s", groupID)
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
