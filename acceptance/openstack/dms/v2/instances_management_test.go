package v2

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/management"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/specification"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	dmsUserPasswordNew = "2ecuredPa22w2rd!"

	dmsConsumerGroup            = "testConsumerGroup"
	dmsConsumerGroupDescription = "test consumer group description"
)

func TestDmsManagement(t *testing.T) {
	t.Skip("DMS Creation takes too long to complete")
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsInstance, err := lifecycle.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "some interesting description", dmsInstance.Description)

	err = resetPass(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 password updated")

	// should wait a little bit
	time.Sleep(5 * time.Second)

	err = setupAutoTopicCreation(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 auto topic creation enabled")

	err = createConsumerGroup(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 consumer group created")

	defer func() {
		errDel := deleteConsumerGroup(t, client, instanceID)
		th.AssertNoErr(t, errDel)
		t.Logf("DMSv2.1 consumer group deleted")
	}()

	err = getConsumerGroup(t, client, instanceID, dmsConsumerGroup)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 consumer group get function done")

	err = getConsumerGroupDetails(t, client, instanceID, dmsConsumerGroup)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 consumer group detailed get function done")

	err = listConsumerGroups(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 list of consumer group is done")

	err = getCoordinator(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 get coordinator is done")

	err = getDiskUsageStatusOfTopics(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 get disk usage is done")

	err = getMetadata(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 get metadata is done")

	err = getSpecification(t, client, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 get specification is done")
}

func resetPass(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	opts := management.PasswordOpts{NewPassword: dmsUserPasswordNew}
	return management.ResetPassword(client, instanceId, opts)
}

func createConsumerGroup(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	opts := management.CreateConsumerGroupOpts{
		GroupName:   dmsConsumerGroup,
		Description: dmsConsumerGroupDescription,
	}

	return management.CreateConsumerGroup(client, instanceId, opts)
}

func getConsumerGroup(t *testing.T, client *golangsdk.ServiceClient, instanceId, groupId string) error {
	t.Helper()
	resp, err := management.GetConsumerGroup(client, instanceId, groupId)
	t.Log("got consumergroup response: ")
	tools.PrintResource(t, resp)

	return err
}

func getConsumerGroupDetails(t *testing.T, client *golangsdk.ServiceClient, instanceId, groupId string) error {
	t.Helper()
	resp, err := management.GetConsumerGroupDetails(client, instanceId, groupId)
	t.Log("got detailed consumergroup response: ")
	tools.PrintResource(t, resp)
	return err
}

func listConsumerGroups(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := management.ListConsumerGroups(client, instanceId, management.ListConsumerGroupsOpts{})
	t.Log("got list of consumer groups: ")
	tools.PrintResource(t, resp)
	return err
}

func getCoordinator(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := management.GetCoordinator(client, instanceId)
	t.Log("got coordinator response: ")
	tools.PrintResource(t, resp)
	return err
}

func getDiskUsageStatusOfTopics(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := management.GetDiskUsageStatusOfTopics(client, instanceId, management.GetDiskUsageOpts{})
	t.Log("got disk usage response: ")
	tools.PrintResource(t, resp)
	return err
}

func getMetadata(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := management.GetMetadata(client, instanceId)
	t.Log("got metadata: ")
	tools.PrintResource(t, resp)
	return err
}

func getSpecification(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()

	resp, err := specification.GetSpec(client, instanceId, specification.GetSpecOpts{
		Engine: dmsEngine,
	})
	t.Log("got specification: ")
	tools.PrintResource(t, resp)
	return err
}

func deleteConsumerGroup(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	return management.DeleteConsumerGroup(client, instanceId, dmsConsumerGroup)
}

func setupAutoTopicCreation(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	opts := management.ConfAutoTopicCreationOpts{EnableAutoTopic: true}

	return management.ConfAutoTopicCreation(client, instanceId, opts)
}
