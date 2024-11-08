package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
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
