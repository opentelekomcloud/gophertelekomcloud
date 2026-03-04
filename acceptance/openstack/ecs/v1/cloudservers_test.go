package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/disk"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCloudServerLifecycle(t *testing.T) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	// Get ECSv1 createOpts
	createOpts := openstack.GetCloudServerCreateOpts(t)

	// Check ECSv1 createOpts
	openstack.DryRunCloudServerConfig(t, client, createOpts)
	t.Logf("CreateOpts are ok for creating a cloudServer")

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, client, createOpts)
	defer openstack.DeleteCloudServer(t, client, ecs.ID)

	// Update ECSv1 instance
	newName := tools.RandomString("ecs-updated-", 3)
	newDescription := "updated ecs description"
	updated, err := cloudservers.Update(client, ecs.ID, cloudservers.UpdateOpts{
		Name:        newName,
		Description: &newDescription,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, newName, updated.Name)
	th.AssertEquals(t, newDescription, updated.Description)
	t.Logf("Successfully updated ECS name to %s", updated.Name)

	tagsList := []tags.ResourceTag{
		{
			Key:   "TestKey",
			Value: "TestValue",
		},
		{
			Key:   "empty",
			Value: "",
		},
	}
	err = tags.Create(client, "cloudservers", ecs.ID, tagsList).ExtractErr()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, ecs)
}

func TestCloudServersRandomAzLifecycle(t *testing.T) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	prefix := "ecs-"
	ecsName := tools.RandomString(prefix, 3)
	imageName := "Standard_Debian_11_latest"
	flavorID := "s3.large.2"

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")

	imageV2Client, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	image, err := images.ListImages(imageV2Client, images.ListImagesOpts{
		Name: imageName,
	})
	th.AssertNoErr(t, err)
	if len(image) == 0 {
		t.Skip("Change image query filter, no results returned")
	}
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID, OS_NETWORK_ID env vars is missing but ECSv1 test requires")
	}

	// Get ECSv1 createOpts
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
		},
		DataVolumes: []cloudservers.DataVolume{
			{
				VolumeType: "SSD",
				Size:       20,
			},
		},
	}

	// Check ECSv1 createOpts
	openstack.DryRunCloudServerConfig(t, client, createOpts)
	t.Logf("CreateOpts are ok for creating a cloudServer")

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, client, createOpts)
	defer openstack.DeleteCloudServer(t, client, ecs.ID)

	tagsList := []tags.ResourceTag{
		{
			Key:   "TestKey",
			Value: "TestValue",
		},
		{
			Key:   "empty",
			Value: "",
		},
	}
	err = tags.Create(client, "cloudservers", ecs.ID, tagsList).ExtractErr()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, ecs)
}

func TestCloudServersIPV6(t *testing.T) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	prefix := "ecs-"
	ecsName := tools.RandomString(prefix, 3)
	imageName := "Standard_Debian_11_latest"
	flavorID := "s3.large.2"

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("IPV6_ENABLED_NETWORK_ID")

	imageV2Client, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	image, err := images.ListImages(imageV2Client, images.ListImagesOpts{
		Name: imageName,
	})
	th.AssertNoErr(t, err)
	if len(image) == 0 {
		t.Skip("Change image query filter, no results returned")
	}
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID, OS_IPV6_ENABLED_NETWORK_ID env vars is missing but ECSv1 test requires")
	}

	// Get ECSv1 createOpts
	createOpts := cloudservers.CreateOpts{
		ImageRef:  image[0].Id,
		FlavorRef: flavorID,
		Name:      ecsName,
		VpcId:     vpcID,
		Nics: []cloudservers.Nic{
			{
				SubnetId:   subnetID,
				Ipv6Enable: true,
			},
		},
		RootVolume: cloudservers.RootVolume{
			VolumeType: "SSD",
		},
		DataVolumes: []cloudservers.DataVolume{
			{
				VolumeType: "SSD",
				Size:       20,
			},
		},
	}

	// Check ECSv1 createOpts
	openstack.DryRunCloudServerConfig(t, client, createOpts)
	t.Logf("CreateOpts are ok for creating a cloudServer")

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, client, createOpts)
	defer openstack.DeleteCloudServer(t, client, ecs.ID)

	ipv6enabled := false
	for _, addresses := range ecs.Addresses {
		for _, addr := range addresses {
			if addr.Version == "6" {
				ipv6enabled = true
			}
		}
	}
	th.AssertEquals(t, true, ipv6enabled)
}

func TestCloudServerVolumeLifecycle(t *testing.T) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	clientEvs, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		t.Skip("OS_AVAILABILITY_ZONE env vars is missing but ECSv1 test requires")
	}
	createVolumeOpts := volumes.CreateOpts{
		Size:             40,
		Name:             tools.RandomString("tf-evs-disk-", 4),
		VolumeType:       "SSD",
		AvailabilityZone: az,
	}

	vol, err := volumes.Create(clientEvs, createVolumeOpts).Extract()
	th.AssertNoErr(t, err)

	err = waitForEvsAvailable(clientEvs, 100, vol.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = volumes.Delete(clientEvs, vol.ID, volumes.DeleteOpts{}).ExtractErr()
		th.AssertNoErr(t, err)
	})

	// Get ECSv1 createOpts
	createOpts := openstack.GetCloudServerCreateOpts(t)

	// Check ECSv1 createOpts
	openstack.DryRunCloudServerConfig(t, client, createOpts)
	t.Logf("CreateOpts are ok for creating a cloudServer")

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, client, createOpts)

	t.Cleanup(func() {
		openstack.DeleteCloudServer(t, client, ecs.ID)
	})

	t.Logf("Attaching volume to cloudserver: %s", vol.ID)
	attach, err := disk.Attach(client, disk.CreateOpts{
		ServerID: ecs.ID,
		VolumeAttachment: &disk.VolumeAttachment{
			VolumeID: vol.ID,
		},
	})
	th.AssertNoErr(t, err)

	err = cloudservers.WaitForJobSuccess(client, 120, attach.JobID)
	th.AssertNoErr(t, err)

	t.Logf("Get all attached volumes to cloudserver: %s", ecs.ID)
	attachments, err := disk.GetAttachments(client, ecs.ID)

	tools.PrintResource(t, attachments)

	t.Logf("Force Detaching volume from cloudserver: %s", vol.ID)
	detach, err := disk.Detach(client, ecs.ID, vol.ID, 1)
	th.AssertNoErr(t, err)

	err = cloudservers.WaitForJobSuccess(client, 120, detach.JobID)
	th.AssertNoErr(t, err)
}

func waitForEvsAvailable(client *golangsdk.ServiceClient, secs int, volId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		vol, err := volumes.Get(client, volId).Extract()
		if err != nil {
			return false, err
		}

		if vol.Status == "available" {
			return true, nil
		}
		return false, nil
	})
}
