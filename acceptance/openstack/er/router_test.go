package er

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	commontags "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnterpriseRouterLifeCycle(t *testing.T) {
	if os.Getenv("RUN_ER_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	routerName := tools.RandomString("acctest_er_router-", 4)

	createOpts := instance.CreateOpts{
		Name:                        routerName,
		Description:                 "terraform created enterprise router",
		Asn:                         64512,
		EnableDefaultAssociation:    pointerto.Bool(true),
		EnableDefaultPropagation:    pointerto.Bool(true),
		AutoAcceptSharedAttachments: pointerto.Bool(true),
		AvailabilityZoneIDs: []string{
			"eu-de-01",
			"eu-de-02",
		},
	}

	t.Logf("Attempting to create enterprise router")

	createResp, err := instance.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = WaitForInstanceAvailable(client, 100, createResp.Instance.ID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, createOpts.Name, createResp.Instance.Name)
	th.AssertEquals(t, createOpts.Description, createResp.Instance.Description)
	th.AssertEquals(t, *createOpts.EnableDefaultPropagation, createResp.Instance.EnableDefaultPropagation)
	th.AssertEquals(t, *createOpts.EnableDefaultAssociation, createResp.Instance.EnableDefaultAssociation)
	th.AssertEquals(t, *createOpts.AutoAcceptSharedAttachments, createResp.Instance.AutoAcceptSharedAttachments)

	t.Cleanup(func() {
		t.Logf("Attempting to delete enterprise router")
		err = instance.Delete(client, createResp.Instance.ID)
		th.AssertNoErr(t, err)
		err = WaitForInstanceDeleted(client, 500, createResp.Instance.ID)
	})

	tOpts := tags.TagOpts{
		Tag: commontags.ResourceTag{
			Key:   "TestKey",
			Value: "TestValue",
		},
	}
	err = tags.Create(client, "instance", createResp.Instance.ID, tOpts)
	th.AssertNoErr(t, err)

	getResp, err := instance.Get(client, createResp.Instance.ID)
	th.AssertNoErr(t, err)

	tList, err := tags.List(client, "instance", createResp.Instance.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, tList[0].Value, getResp.Instance.Tags[0].Value)

	err = tags.Delete(client, "instance", createResp.Instance.ID, tList[0].Key)
	th.AssertNoErr(t, err)

	tListEmpty, err := tags.List(client, "instance", createResp.Instance.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(tListEmpty), 0)

	updateOpts := instance.UpdateOpts{
		InstanceID:                  createResp.Instance.ID,
		Name:                        routerName + "-updated",
		EnableDefaultAssociation:    pointerto.Bool(false),
		EnableDefaultPropagation:    pointerto.Bool(false),
		AutoAcceptSharedAttachments: pointerto.Bool(false),
		Description:                 createOpts.Description + " updated",
	}

	t.Logf("Attempting to update enterprise router")

	updateResp, err := instance.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, updateOpts.Name, updateResp.Instance.Name)
	th.AssertEquals(t, updateOpts.Description, updateResp.Instance.Description)
	th.AssertEquals(t, *updateOpts.EnableDefaultPropagation, updateResp.Instance.EnableDefaultPropagation)
	th.AssertEquals(t, *updateOpts.EnableDefaultAssociation, updateResp.Instance.EnableDefaultAssociation)
	th.AssertEquals(t, *updateOpts.AutoAcceptSharedAttachments, updateResp.Instance.AutoAcceptSharedAttachments)
}

func TestEnterpriseRouterList(t *testing.T) {
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	listOpts := instance.ListOpts{}
	_, err = instance.List(client, listOpts)
	th.AssertNoErr(t, err)
}
