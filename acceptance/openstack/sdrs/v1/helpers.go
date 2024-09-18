package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/sdrs/v1/protectiongroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createSDRSGroup(t *testing.T, client *golangsdk.ServiceClient, domainID string) *protectiongroups.Group {
	t.Logf("Attempting to create SDRS protection group")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID env var is missing but SDRS group test requires")
	}

	serverGroup := protectiongroups.ServerGroupInfo{
		Name:        tools.RandomString("sdrs-group-", 3),
		Description: "some interesting description",
		SourceAZ:    "eu-de-02",
		TargetAZ:    "eu-de-01",
		DomainID:    domainID,
		SourceVpcID: vpcID,
	}

	createOpts := protectiongroups.CreateOpts{
		ServerGroup: serverGroup,
	}

	job, err := protectiongroups.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Logf("Waiting for SDRS group job %s", job)
	err = protectiongroups.WaitForJobSuccess(client, 600, job.JobId)
	th.AssertNoErr(t, err)

	jobEntity, err := protectiongroups.GetJobEntity(client, job.JobId, "server_group_id")
	th.AssertNoErr(t, err)

	group, err := protectiongroups.Get(client, jobEntity.(string)).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created SDRS protection group: %s", group.Id)

	return group
}

func deleteSDRSGroup(t *testing.T, client *golangsdk.ServiceClient, groupID string) {
	t.Logf("Attempting to delete SDRS protection group: %s", groupID)

	job, err := protectiongroups.Delete(client, groupID).ExtractJobResponse()
	th.AssertNoErr(t, err)

	err = protectiongroups.WaitForJobSuccess(client, 600, job.JobID)
	th.AssertNoErr(t, err)

	t.Logf("Deleted SDRS protection group: %s", groupID)
}
