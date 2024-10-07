package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	lifecyclehooks "github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/lifecycle_hooks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLifecycleHooksLifecycle(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if networkID == "" || vpcID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but are required for AS Lifecycle Hooks test")
	}

	groupID := autoscaling.CreateAutoScalingGroup(t, client, networkID, vpcID, tools.RandomString("as-group-create-", 3))
	t.Cleanup(func() {
		autoscaling.DeleteAutoScalingGroup(t, client, groupID)
	})

	topicName := tools.RandomString("as-lifecycle-hooks-topic-", 3)
	lifecycleHookName := tools.RandomString("as-lifecycle-hook-create-", 3)
	topicURN, err := autoscaling.GetNotificationTopicURN(topicName)
	if err != nil {
		t.Logf("Error while creating the notification topic: %s", topicName)
	}
	defer autoscaling.DeleteTopic(t, topicURN)
	createOpts := lifecyclehooks.CreateOpts{
		LifecycleHookName:    lifecycleHookName,
		LifecycleHookType:    "INSTANCE_LAUNCHING",
		NotificationTopicUrn: topicURN,
	}

	t.Logf("Attempting to create Lifecycle Hook")
	lifecycleHook, err := lifecyclehooks.Create(client, createOpts, groupID)
	th.AssertNoErr(t, err)
	t.Logf("Ceate Lifecycle Hook: %s", lifecycleHook.LifecycleHookName)

	requestedLifecycleHook, err := lifecyclehooks.Get(client, groupID, lifecycleHookName)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, requestedLifecycleHook, lifecycleHook)

	updateOpts := lifecyclehooks.UpdateOpts{
		LifecycleHookType: "INSTANCE_TERMINATING",
		DefaultTimeout:    4800,
	}
	t.Logf("Attempting to update Lifecycle Hook")
	lifecycleHook, err = lifecyclehooks.Update(client, updateOpts, groupID, lifecycleHookName)
	th.AssertEquals(t, updateOpts.DefaultTimeout, lifecycleHook.DefaultTimeout)
	th.AssertEquals(t, updateOpts.LifecycleHookType, lifecycleHook.LifecycleHookType)
	th.AssertNoErr(t, err)
	t.Logf("Updated Lifecycle Hook: %s", lifecycleHookName)

	t.Logf("Attempting to delete Lifecycle Hook")
	err = lifecyclehooks.Delete(client, groupID, lifecycleHookName)
	th.AssertNoErr(t, err)
	t.Logf("Deleted Lifecycle Hook: %s", lifecycleHookName)
}
