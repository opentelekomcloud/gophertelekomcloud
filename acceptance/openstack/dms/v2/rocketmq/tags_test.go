package rocketmq

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	rmqtags "github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRocketMQTags(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	keep := tags.ResourceTag{Key: tools.RandomString("acc-key-", 4), Value: "acc-value"}
	drop := tags.ResourceTag{Key: tools.RandomString("acc-key-", 4), Value: ""}

	t.Logf("Attempting to add tags to RocketMQ instance: %s", instanceID)
	th.AssertNoErr(t, rmqtags.Create(client, instanceID, []tags.ResourceTag{keep, drop}))

	t.Cleanup(func() {
		err := rmqtags.Delete(client, instanceID, []tags.ResourceTag{keep, drop})
		if err != nil {
			t.Errorf("failed to delete tags of RocketMQ instance %s: %s", instanceID, err)
		}
	})

	instanceTags, err := rmqtags.Get(client, instanceID, rmqtags.ListOpts{Limit: 20})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, instanceTags)
	th.AssertEquals(t, "acc-value", findTag(t, instanceTags.Tags, keep.Key).Value)
	th.AssertEquals(t, "", findTag(t, instanceTags.Tags, drop.Key).Value)

	projectTags, err := rmqtags.List(client, rmqtags.ListOpts{Limit: 100})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, projectTags)
	var found bool
	for _, tag := range projectTags.Tags {
		if tag.Key == keep.Key {
			found = true
			th.AssertDeepEquals(t, []string{"acc-value"}, tag.Values)
			break
		}
	}
	th.AssertEquals(t, true, found)

	t.Logf("Attempting to delete tag %s of RocketMQ instance: %s", drop.Key, instanceID)
	th.AssertNoErr(t, rmqtags.Delete(client, instanceID, []tags.ResourceTag{drop}))

	instanceTags, err = rmqtags.Get(client, instanceID, rmqtags.ListOpts{Limit: 20})
	th.AssertNoErr(t, err)
	for _, tag := range instanceTags.Tags {
		if tag.Key == drop.Key {
			t.Fatalf("tag %s is still present after deletion", drop.Key)
		}
	}
	findTag(t, instanceTags.Tags, keep.Key)
}

func findTag(t *testing.T, tagList []tags.ResourceTag, key string) tags.ResourceTag {
	for _, tag := range tagList {
		if tag.Key == key {
			return tag
		}
	}
	t.Fatalf("no tag with key %s", key)
	return tags.ResourceTag{}
}
