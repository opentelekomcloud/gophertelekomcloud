package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/subscriptions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTopicSubscriptionWorkflow(t *testing.T) {
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	topic := createTopic(t, client)
	defer deleteTopic(t, client, topic)

	t.Logf("Attempting to create SMN subscription")
	createOpts := subscriptions.CreateOpts{
		Endpoint: "example@email.com",
		Protocol: "email",
	}

	subscription, err := subscriptions.Create(client, createOpts, topic).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created SMN subscription: %s", subscription.SubscriptionUrn)
	defer func() {
		t.Logf("Attempting to delete SMN subscription: %s", subscription.SubscriptionUrn)
		err := subscriptions.Delete(client, subscription.SubscriptionUrn).Err
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN subscription: %s", subscription.SubscriptionUrn)
	}()

	subscriptionList, err := subscriptions.List(client).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(subscriptionList))
}
