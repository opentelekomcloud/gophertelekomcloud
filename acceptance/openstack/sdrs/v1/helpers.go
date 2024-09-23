package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/protectiongroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createSDRSGroup(t *testing.T, client *golangsdk.ServiceClient, domainID string) *protectiongroups.ServerGroupResponseInfo {
	t.Logf("Attempting to create SDRS protection group")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if vpcID == "" {
		t.Skip("OS_VPC_ID env var is missing but SDRS group test requires")
	}

	createOpts := protectiongroups.CreateOpts{
		Name:        tools.RandomString("sdrs-group-", 3),
		Description: "some interesting description",
		SourceAZ:    az,
		TargetAZ:    "eu-de-01",
		DomainID:    domainID,
		SourceVpcID: vpcID,
	}

	job, err := protectiongroups.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for SDRS group job %s", job.JobID)
	err = protectiongroups.WaitForJobSuccess(client, 600, job.JobID)
	th.AssertNoErr(t, err)

	jobEntity, err := protectiongroups.GetJobEntity(client, job.JobID, "server_group_id")
	th.AssertNoErr(t, err)

	group, err := protectiongroups.Get(client, jobEntity.(string))
	th.AssertNoErr(t, err)

	t.Logf("Created SDRS protection group: %s", group.Id)

	return group
}

func deleteSDRSGroup(t *testing.T, client *golangsdk.ServiceClient, groupID string) {
	t.Logf("Attempting to delete SDRS protection group: %s", groupID)

	job, err := protectiongroups.Delete(client, groupID)
	th.AssertNoErr(t, err)

	err = protectiongroups.WaitForJobSuccess(client, 600, job.JobID)
	th.AssertNoErr(t, err)

	t.Logf("Deleted SDRS protection group: %s", groupID)
}
