package tms

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/tms/v1/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTMSV1Lifecycle(t *testing.T) {
	if os.Getenv("RUN_TMS_TAGS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewTmsV1Client()
	th.AssertNoErr(t, err)

	predefinedTags := []tags.Tag{
		{
			Key:   "test-1",
			Value: "test-1",
		},
	}

	createOpts := tags.BatchOpts{
		Tags:   predefinedTags,
		Action: tags.ActionCreate,
	}

	_, err = tags.BatchAction(client, "", createOpts).Extract()
	th.AssertNoErr(t, err)

	listTags, err := tags.Get(client).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listTags.Tags[0].Key, "test-1")
	th.AssertEquals(t, listTags.Tags[0].Value, "test-1")

	deleteOpts := tags.BatchOpts{
		Action: tags.ActionDelete,
		Tags:   predefinedTags,
	}
	_, err = tags.BatchAction(client, "", deleteOpts).Extract()
	th.AssertNoErr(t, err)

	_, err = tags.Get(client).Extract()
	th.AssertNoErr(t, err)
}
