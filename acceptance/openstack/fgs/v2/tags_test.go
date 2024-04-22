package v2

import (
	"log"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	rstag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphTags(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")
	t.Cleanup(func() {
		err = function.Delete(client, funcUrn)
		th.AssertNoErr(t, err)
	})

	createTagsOpts := tags.TagsActionOpts{
		Id:     funcUrn,
		Action: "create",
		Tags: []rstag.ResourceTag{
			{
				Key:   "test",
				Value: "test-2",
			},
		},
	}

	log.Printf("Attempting to create Funcgraph tag")
	err = tags.CreateResourceTag(client, createTagsOpts)
	th.AssertNoErr(t, err)

	// API no published
	// getTags, err := tags.GetResourceTags(client, funcUrn)
	// th.AssertNoErr(t, err)
	// tools.PrintResource(t, getTags)

	log.Printf("Attempting to delete Funcgraph tag")
	createTagsOpts.Action = "delete"
	err = tags.DeleteResourceTag(client, createTagsOpts)
	th.AssertNoErr(t, err)
}
