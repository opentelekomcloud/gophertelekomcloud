// Package openstack contains common functions that can be used
// across all OpenStack components for acceptance testing.
package openstack

import (
	"fmt"
	"net"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/extensions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// PrintExtension prints an extension and all of its attributes.
func PrintExtension(t *testing.T, extension *extensions.Extension) {
	t.Logf("Name: %s", extension.Name)
	t.Logf("Namespace: %s", extension.Namespace)
	t.Logf("Alias: %s", extension.Alias)
	t.Logf("Description: %s", extension.Description)
	t.Logf("Updated: %s", extension.Updated)
	t.Logf("Links: %v", extension.Links)
}

func DefaultSecurityGroup(t *testing.T) string {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	securityGroupPages, err := secgroups.List(client).AllPages()
	th.AssertNoErr(t, err)
	securityGroups, err := secgroups.ExtractSecurityGroups(securityGroupPages)
	th.AssertNoErr(t, err)
	var sgId string
	for _, val := range securityGroups {
		if val.Name == "default" {
			sgId = val.ID
			break
		}
	}
	if sgId == "" {
		t.Fatalf("Unable to find default secgroup")
	}
	return sgId
}

func CreateSecurityGroup(t *testing.T) string {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	createSGOpts := secgroups.CreateOpts{
		Name:        tools.RandomString("acc-sg-", 3),
		Description: "security group for acceptance testing",
	}
	secGroup, err := secgroups.Create(client, createSGOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Security group %s was created", secGroup.ID)
	return secGroup.ID
}

func DeleteSecurityGroup(t *testing.T, secGroupID string) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	err = secgroups.DeleteWithRetry(client, secGroupID, 600)
	th.AssertNoErr(t, err)

	t.Logf("Security group %s was deleted", secGroupID)
}

func CreateVolume(t *testing.T) *volumes.Volume {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)
	vol, err := volumes.Create(client, volumes.CreateOpts{
		Name:       tools.RandomString("test-vol-", 6),
		Size:       10,
		VolumeType: "SSD",
	}).Extract()
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(300, func() (bool, error) {
		volume, err := volumes.Get(client, vol.ID).Extract()
		if err != nil {
			return false, err
		}
		if volume.Status == "available" {
			return true, nil
		}
		if volume.Status == "error" {
			return false, fmt.Errorf("error creating a volume")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return vol
}

func DeleteVolume(t *testing.T, id string) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, volumes.Delete(client, id, volumes.DeleteOpts{Cascade: true}).Err)
}

const (
	imageName = "Standard_Debian_10_latest"
	flavorID  = "s3.large.2"
)

func GetCloudServerCreateOpts(t *testing.T) cloudservers.CreateOpts {
	prefix := "ecs-"
	ecsName := tools.RandomString(prefix, 3)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	var encryption string
	if kmsID != "" {
		encryption = "1"
	} else {
		encryption = "0"
	}

	imageV2Client, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	image, err := images.ListImages(imageV2Client, images.ListImagesOpts{
		Name: imageName,
	})
	th.AssertNoErr(t, err)

	if vpcID == "" || subnetID == "" || az == "" {
		t.Skip("One of OS_VPC_ID, OS_NETWORK_ID or OS_AVAILABILITY_ZONE env vars is missing but ECSv1 test requires")
	}

	createOpts := cloudservers.CreateOpts{
		ImageRef:  image[0].Id,
		FlavorRef: flavorID,
		Name:      ecsName,
		VpcId:     vpcID,
		Nics: []cloudservers.Nic{
			{
				SubnetId: subnetID,
			},
		},
		RootVolume: cloudservers.RootVolume{
			VolumeType: "SSD",
			Metadata: map[string]interface{}{
				"__system__encrypted": encryption,
				"__system__cmkid":     kmsID,
			},
		},
		DataVolumes: []cloudservers.DataVolume{
			{
				VolumeType: "SSD",
				Size:       40,
				Metadata: map[string]interface{}{
					"__system__encrypted": encryption,
					"__system__cmkid":     kmsID,
				},
			},
		},
		AvailabilityZone: az,
	}

	return createOpts
}

func DryRunCloudServerConfig(t *testing.T, client *golangsdk.ServiceClient, createOpts cloudservers.CreateOpts) {
	t.Logf("Attempting to check ECSv1 createOpts")
	err := cloudservers.DryRun(client, createOpts).Err
	th.AssertNoErr(t, err)
}

func CreateCloudServer(t *testing.T, client *golangsdk.ServiceClient, createOpts cloudservers.CreateOpts) *cloudservers.CloudServer {
	t.Logf("Attempting to create ECSv1")

	jobResponse, err := cloudservers.Create(client, createOpts).ExtractJobResponse()
	th.AssertNoErr(t, err)

	err = cloudservers.WaitForJobSuccess(client, 1200, jobResponse.JobID)
	th.AssertNoErr(t, err)

	serverID, err := cloudservers.GetJobEntity(client, jobResponse.JobID, "server_id")
	th.AssertNoErr(t, err)

	ecs, err := cloudservers.Get(client, serverID.(string)).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created ECSv1 instance: %s", ecs.ID)

	return ecs
}

func DeleteCloudServer(t *testing.T, client *golangsdk.ServiceClient, ecsID string) {
	t.Logf("Attempting to delete ECSv1: %s", ecsID)

	deleteOpts := cloudservers.DeleteOpts{
		Servers: []cloudservers.Server{
			{
				Id: ecsID,
			},
		},
		DeletePublicIP: true,
		DeleteVolume:   true,
	}
	jobResponse, err := cloudservers.Delete(client, deleteOpts).ExtractJobResponse()
	th.AssertNoErr(t, err)

	err = cloudservers.WaitForJobSuccess(client, 1200, jobResponse.JobID)
	th.AssertNoErr(t, err)

	t.Logf("ECSv1 instance deleted: %s", ecsID)
}

// ValidIP returns valid value for IP in subnet by VPC subnet ID / OpenStack network ID
func ValidIP(t *testing.T, networkID string) string {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	sn, err := subnets.Get(client, networkID).Extract()
	th.AssertNoErr(t, err)

	_, nw, err := net.ParseCIDR(sn.CIDR)
	th.AssertNoErr(t, err)
	singleIP := nw.IP
	singleIP[len(singleIP)-1] += 3
	th.AssertEquals(t, true, nw.Contains(singleIP))
	return singleIP.String()
}
