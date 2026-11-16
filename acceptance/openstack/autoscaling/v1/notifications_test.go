package v1

import (
	"sort"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/notifications"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/subscriptions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func notificationTestGroup(t *testing.T) (*golangsdk.ServiceClient, string, string) {
	t.Helper()
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	endpoint := clients.EnvOS.GetEnv("SMN_SUBSCRIPTION_ENDPOINT")
	if networkID == "" || vpcID == "" || endpoint == "" {
		t.Skip("OS_NETWORK_ID, OS_VPC_ID and OS_SMN_SUBSCRIPTION_ENDPOINT (email address) are required")
	}
	keyPairName := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	imageID := clients.EnvOS.GetEnv("IMAGE_ID")
	if keyPairName == "" || imageID == "" {
		t.Skip("OS_KEYPAIR_NAME or OS_IMAGE_ID env vars is missing but AS Configuration test requires")
	}
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)
	smnClient, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	topic, err := topics.Create(smnClient, topics.CreateOpts{
		Name: tools.RandomString("as-notification-topic-", 8),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, topics.Delete(smnClient, topic.TopicUrn))
	})
	subscription, err := subscriptions.Create(smnClient, subscriptions.CreateOpts{
		Endpoint: endpoint,
		Protocol: "email",
	}, topic.TopicUrn)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, subscriptions.Delete(smnClient, subscription.SubscriptionUrn))
	})

	asCreateName := tools.RandomString("as-create-", 3)
	configID := autoscaling.CreateASConfig(t, client, asCreateName, imageID, keyPairName)

	t.Cleanup(func() {
		autoscaling.DeleteASConfig(t, client, configID)
	})

	groupID, err := groups.Create(client, groups.CreateOpts{
		Name: tools.RandomString("as-notification-", 8), ConfigurationID: configID,
		Networks: []groups.ID{{ID: networkID}}, VpcID: vpcID,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, groups.Delete(client, groups.DeleteOpts{ScalingGroupId: groupID}))
	})
	return client, groupID, topic.TopicUrn
}

func TestNotificationLifecycle(t *testing.T) {
	client, groupID, topicURN := notificationTestGroup(t)
	scenes := []string{"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL"}
	actual, err := notifications.Create(client, notifications.CreateOpts{
		TopicURN: topicURN, TopicScene: scenes,
	}, groupID)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, notifications.Delete(client, groupID, topicURN))
	})
	th.AssertEquals(t, topicURN, actual.TopicURN)
	th.AssertEquals(t, true, actual.TopicName != "")
	sort.Strings(scenes)
	sort.Strings(actual.TopicScene)
	th.AssertDeepEquals(t, scenes, actual.TopicScene)
	listed, err := notifications.List(client, groupID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listed))
	sort.Strings(listed[0].TopicScene)
	th.AssertDeepEquals(t, *actual, listed[0])
}
