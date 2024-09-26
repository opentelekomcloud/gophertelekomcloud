package v2

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
	instMgmt "github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances/management"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances/specification"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	dmsUserPasswordNew = "2ecuredPa22w2rd!"

	dmsConsumerGroup            = "testConsumerGroup"
	dmsConsumerGroupDescription = "test consumer group description"
)

func TestDmsManagement(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsInstance, err := instances.Get(client, instanceID)
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
	opts := instMgmt.PasswordOpts{NewPassword: dmsUserPasswordNew}
	return instMgmt.ResetPassword(client, instanceId, opts)
}

func createConsumerGroup(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	opts := instMgmt.CreateConsumerGroupOpts{
		GroupName:   dmsConsumerGroup,
		Description: dmsConsumerGroupDescription,
	}

	return instMgmt.CreateConsumerGroup(client, instanceId, opts)
}

func getConsumerGroup(t *testing.T, client *golangsdk.ServiceClient, instanceId, groupId string) error {
	t.Helper()
	resp, err := instMgmt.GetConsumerGroup(client, instanceId, groupId)
	t.Log("got consumergroup response: ")
	tools.PrintResource(t, resp)

	return err
}

func getConsumerGroupDetails(t *testing.T, client *golangsdk.ServiceClient, instanceId, groupId string) error {
	t.Helper()
	resp, err := instMgmt.GetConsumerGroupDetails(client, instanceId, groupId)
	t.Log("got detailed consumergroup response: ")
	tools.PrintResource(t, resp)
	return err
}

func listConsumerGroups(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := instMgmt.ListConsumerGroups(client, instanceId, instMgmt.ListConsumerGroupsOpts{})
	t.Log("got list of consumer groups: ")
	tools.PrintResource(t, resp)
	return err
}

func getCoordinator(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := instMgmt.GetCoordinator(client, instanceId)
	t.Log("got coordinator response: ")
	tools.PrintResource(t, resp)
	return err
}

func getDiskUsageStatusOfTopics(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := instMgmt.GetDiskUsageStatusOfTopics(client, instanceId, instMgmt.GetDiskUsageOpts{})
	t.Log("got disk usage response: ")
	tools.PrintResource(t, resp)
	return err
}

func getMetadata(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	resp, err := instMgmt.GetMetadata(client, instanceId)
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
	return instMgmt.DeleteConsumerGroup(client, instanceId, dmsConsumerGroup)
}

func setupAutoTopicCreation(t *testing.T, client *golangsdk.ServiceClient, instanceId string) error {
	t.Helper()
	opts := instMgmt.ConfAutoTopicCreationOpts{EnableAutoTopic: true}

	return instMgmt.ConfAutoTopicCreation(client, instanceId, opts)
}