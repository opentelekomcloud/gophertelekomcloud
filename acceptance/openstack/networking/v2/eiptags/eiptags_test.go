package eiptags

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/eiptags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEipsTagsList(t *testing.T) {
	clientV1, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	eip := CreateEip(t, clientV1)
	defer DeleteEip(t, clientV1, eip.ID)

	clientV2, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	tags, err := eiptags.List(clientV2, eip.ID).Extract()
	th.AssertNoErr(t, err)

	for _, tag := range tags {
		tools.PrintResource(t, tag)
	}
}

func TestEipsTagsLifecycle(t *testing.T) {
	clientV1, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	eip := CreateEip(t, clientV1)
	defer DeleteEip(t, clientV1, eip.ID)

	clientV2, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	tagKey := "muh"
	CreateTag(t, clientV2, eip.ID, tagKey)
	defer DeleteTag(t, clientV2, eip.ID, tagKey)

	tagKeys := []string{
		"luh", "kuh",
	}
	CreateTags(t, clientV2, eip.ID, tagKeys)
	defer DeleteTags(t, clientV2, eip.ID, tagKeys)

	tagList, err := eiptags.List(clientV2, eip.ID).Extract()
	th.AssertNoErr(t, err)

	for _, tag := range tagList {
		tools.PrintResource(t, tag)
	}
}
