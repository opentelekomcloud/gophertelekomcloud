package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/domains"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/protectedinstances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/protectiongroups"
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

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	domainList, err := domains.Get(client).Extract()
	th.AssertNoErr(t, err)
	if len(domainList.Domains) == 0 {
		t.Skipf("you don't have any active-active domain, but SDRS test requires")
	}

	group := createSDRSGroup(t, client, domainList.Domains[0].Id)
	defer deleteSDRSGroup(t, client, group.Id)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	defer func() {
		openstack.DeleteCloudServer(t, computeClient, ecs.ID)
	}()

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
		deleteServer := false
		deleteOpts := protectedinstances.DeleteOpts{
			DeleteTargetServer: &deleteServer,
		}

		jobDelete, err := protectedinstances.Delete(client, instance.ID, deleteOpts).ExtractJobResponse()
		th.AssertNoErr(t, err)

		err = protectedinstances.WaitForJobSuccess(client, 600, jobDelete.JobID)
		th.AssertNoErr(t, err)

		t.Logf("Deleted SDRS protected instance: %s", instance.ID)
	}()
	th.AssertEquals(t, createName, instance.Name)
	th.AssertEquals(t, createDescription, instance.Description)

	t.Logf("Created SDRS protected instance: %s", instance.ID)

	jobEnable, err := protectiongroups.Enable(client, group.Id)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for SDRS group enabling job %s", jobEnable.JobID)
	err = protectiongroups.WaitForJobSuccess(client, 600, jobEnable.JobID)
	th.AssertNoErr(t, err)

	getEnablePg, err := protectiongroups.Get(client, group.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "protected", getEnablePg.Status)

	jobDisable, err := protectiongroups.Disable(client, group.Id)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for SDRS group disabling job %s", jobDisable.JobID)
	err = protectiongroups.WaitForJobSuccess(client, 600, jobDisable.JobID)
	th.AssertNoErr(t, err)

	getDisabledPg, err := protectiongroups.Get(client, group.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "available", getDisabledPg.Status)

}
