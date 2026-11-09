package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestKafkaTags(t *testing.T) {
	name := kafkaInstanceName(t)

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := existingKafkaInstance(t, client, name).InstanceID

	tag := tags.TagEntity{Key: tools.RandomString("sdk-acc-", 6), Value: pointerto.String("value1")}
	deleteOpts := tags.ActionOpts{Action: "delete", Tags: []tags.TagEntity{tag}}

	err = tags.Action(client, instanceID, tags.ActionOpts{Action: "create", Tags: []tags.TagEntity{tag}})
	th.AssertNoErr(t, err)
	deleted := false
	t.Cleanup(func() {
		if !deleted {
			th.AssertNoErr(t, tags.Action(client, instanceID, deleteOpts))
		}
	})

	instanceTags, err := tags.Get(client, instanceID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, instanceTags)
	th.AssertEquals(t, "value1", findInstanceTag(t, instanceTags.Tags, tag.Key).Value)

	projectTags, err := tags.List(client)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []string{"value1"}, findProjectTag(t, projectTags.Tags, tag.Key).Values)

	th.AssertNoErr(t, tags.Action(client, instanceID, deleteOpts))
	deleted = true

	err = golangsdk.WaitFor(60, func() (bool, error) {
		instanceTags, err := tags.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		for _, v := range instanceTags.Tags {
			if v.Key == tag.Key {
				return false, nil
			}
		}
		return true, nil
	})
	th.AssertNoErr(t, err)
}

func findInstanceTag(t *testing.T, list []tags.Tag, key string) tags.Tag {
	t.Helper()
	for _, v := range list {
		if v.Key == key {
			return v
		}
	}
	t.Fatalf("instance tag %q not found", key)
	return tags.Tag{}
}

func findProjectTag(t *testing.T, list []tags.ProjectTag, key string) tags.ProjectTag {
	t.Helper()
	for _, v := range list {
		if v.Key == key {
			return v
		}
	}
	t.Fatalf("project tag %q not found", key)
	return tags.ProjectTag{}
}
