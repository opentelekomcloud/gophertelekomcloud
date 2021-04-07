package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/protectedinstances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSDRSInstanceList(t *testing.T) {
	client, err := clients.NewSDRSV1()
	th.AssertNoErr(t, err)

	listOpts := protectedinstances.ListOpts{}
	allPages, err := protectedinstances.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	sdrsInstances, err := protectedinstances.ExtractInstances(allPages)
	th.AssertNoErr(t, err)

	for _, instance := range sdrsInstances {
		tools.PrintResource(t, instance)
	}
}

func TestSDRSInstanceLifecycle(t *testing.T) {
	client, err := clients.NewSDRSV1()
	th.AssertNoErr(t, err)

	group := createSDRSGroup(t, client)
	defer deleteSDRSGroup(t, client, group.Id)

	ecs := openstack.CreateCloudServer(t, client, openstack.GetCloudServerCreateOpts(t))
	defer openstack.DeleteCloudServer(t, client, ecs.ID)

	t.Logf("Attempting to create SDRS protected instance")
	createName := tools.RandomString("sdrs-instance-", 3)
	createDescription := "some description"
	createOpts := protectedinstances.CreateOpts{
		GroupID:     group.Id,
		ServerID:    ecs.ID,
		Name:        createName,
		Description: createDescription,
	}

	jobCreate, err := protectedinstances.Create(client, createOpts).ExtractJobResponse()
	th.AssertNoErr(t, err)
	err = protectedinstances.WaitForJobSuccess(client, 600, jobCreate.JobID)
	th.AssertNoErr(t, err)

	jobEntity, err := protectedinstances.GetJobEntity(client, jobCreate.JobID, "protected_instance_id")
	th.AssertNoErr(t, err)

	instance, err := protectedinstances.Get(client, jobEntity.(string)).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		t.Logf("Attempting to delete SDRS protected instance: %s", instance.ID)

		jobDelete, err := protectedinstances.Delete(client, instance.ID, nil).ExtractJobResponse()
		th.AssertNoErr(t, err)

		err = protectedinstances.WaitForJobSuccess(client, 600, jobDelete.JobID)
		th.AssertNoErr(t, err)

		t.Logf("Deleted SDRS protected instance: %s", instance.ID)
	}()
	th.AssertEquals(t, createName, instance.Name)
	th.AssertEquals(t, createDescription, instance.Description)

	t.Logf("Created SDRS protected instance: %s", instance.ID)
}
