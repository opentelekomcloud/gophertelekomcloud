package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func getCloudServerCreateOpts(t *testing.T) cloudservers.CreateOpts {
	prefix := "ecs-"
	ecsName := tools.RandomString(prefix, 3)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-01"
	}

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	imageID := clients.EnvOS.GetEnv("IMAGE_ID")
	flavorID := clients.EnvOS.GetEnv("FLAVOR_ID")
	keyPairName := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	if vpcID == "" || subnetID == "" || imageID == "" || flavorID == "" || keyPairName == "" {
		t.Skip(`One of OS_VPC_ID or OS_NETWORK_ID or OS_IMAGE_ID or
OS_FLAVOR_ID or OS_KEYPAIR_NAME env vars is missing but ECSv1 test requires`)
	}

	createOpts := cloudservers.CreateOpts{
		ImageRef:  imageID,
		FlavorRef: flavorID,
		Name:      ecsName,
		KeyName:   keyPairName,
		VpcId:     vpcID,
		Nics: []cloudservers.Nic{
			{
				SubnetId: subnetID,
			},
		},
		RootVolume: cloudservers.RootVolume{
			VolumeType: "SATA",
		},
		AvailabilityZone: az,
	}

	return createOpts
}

func dryRunCloudServerConfig(t *testing.T, client *golangsdk.ServiceClient, createOpts cloudservers.CreateOpts) {
	t.Logf("Attempting to check ECSv1 createOpts")
	createOpts.DryRun = true
	err := cloudservers.DryRun(client, createOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func createCloudServer(t *testing.T, client *golangsdk.ServiceClient, createOpts cloudservers.CreateOpts) *cloudservers.CloudServer {
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

func deleteCloudServer(t *testing.T, client *golangsdk.ServiceClient, ecsID string) {
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
