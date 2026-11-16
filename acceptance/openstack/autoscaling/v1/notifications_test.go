package v1

import (
	"sort"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/notifications"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func notificationTestGroup(t *testing.T) (*golangsdk.ServiceClient, string, string) {
	t.Helper()
	configurationID := clients.EnvOS.GetEnv("AS_CONFIGURATION_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	topicURN := clients.EnvOS.GetEnv("SMN_TOPIC_URN")
	if configurationID == "" || networkID == "" || vpcID == "" || topicURN == "" {
		t.Skip("OS_AS_CONFIGURATION_ID, OS_NETWORK_ID, OS_VPC_ID and OS_SMN_TOPIC_URN (with subscriptions) are required")
	}
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)
	groupID, err := groups.Create(client, groups.CreateOpts{
		Name: tools.RandomString("as-notification-", 8), ConfigurationID: configurationID,
		Networks: []groups.ID{{ID: networkID}}, VpcID: vpcID,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, groups.Delete(client, groups.DeleteOpts{ScalingGroupId: groupID}))
	})
	return client, groupID, topicURN
}

func TestNotificationCreate(t *testing.T) {
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

func TestNotificationDelete(t *testing.T) {
	client, groupID, topicURN := notificationTestGroup(t)
	_, err := notifications.Create(client, notifications.CreateOpts{
		TopicURN: topicURN, TopicScene: []string{"SCALING_UP"},
	}, groupID)
	th.AssertNoErr(t, err)
	deleted := false
	t.Cleanup(func() {
		if !deleted {
			th.AssertNoErr(t, notifications.Delete(client, groupID, topicURN))
		}
	})
	th.AssertNoErr(t, notifications.Delete(client, groupID, topicURN))
	deleted = true
	listed, err := notifications.List(client, groupID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(listed))
}
