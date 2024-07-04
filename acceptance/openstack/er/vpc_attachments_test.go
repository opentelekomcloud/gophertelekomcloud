package er

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/vpc"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

var (
	vpcId       = os.Getenv("VPC_ID")
	networkId   = os.Getenv("NETWORK_ID")
	vpcName     = tools.RandomString("acctest_vpc_attachments-", 4)
	description = "test vpc attachment"
)

func TestVPCAttachmentsLifeCycle(t *testing.T) {
	if os.Getenv("RUN_ER_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	createOpts := instance.CreateOpts{
		Name:                        tools.RandomString("acctest_er_router-", 4),
		Asn:                         64512,
		AutoAcceptSharedAttachments: pointerto.Bool(true),
		AvailabilityZoneIDs: []string{
			"eu-de-01",
			"eu-de-02",
		},
	}

	t.Logf("Attempting to create enterprise router")

	createResp, err := instance.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 100, createResp.Instance.ID)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, id string) {
		t.Logf("Attempting to delete enterprise router")
		err = instance.Delete(client, id)
		th.AssertNoErr(t, err)
		err = waitForInstanceDeleted(client, 500, createResp.Instance.ID)
	}(client, createResp.Instance.ID)

	createVpcOpts := vpc.CreateOpts{
		Name:                vpcName,
		RouterID:            createResp.Instance.ID,
		VpcId:               vpcId,
		SubnetId:            networkId,
		Description:         description,
		AutoCreateVpcRoutes: true,
		Tags: []tag.ResourceTag{
			{
				Key:   "muh",
				Value: "muh",
			},
			{
				Key:   "test",
				Value: "test",
			},
		},
	}

	t.Logf("Attempting to create vpc attachemnt")
	createVpcResp, err := vpc.Create(client, createVpcOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createVpcOpts.Name, createVpcResp.Name)
	th.AssertEquals(t, createVpcOpts.Description, createVpcResp.Description)
	th.AssertEquals(t, createVpcOpts.AutoCreateVpcRoutes, createVpcResp.AutoCreateVpcRoutes)

	err = waitForVpcAttachmentsAvailable(client, 100, createResp.Instance.ID, createVpcResp.ID)
	th.AssertNoErr(t, err)

	defer func(client *golangsdk.ServiceClient, erId, vpcId string) {
		t.Logf("Attempting to delete vpc attachemnt")
		err = vpc.Delete(client, erId, vpcId)
		th.AssertNoErr(t, err)
		err = waitForVpcAttachmentsDeleted(client, 500, erId, vpcId)
	}(client, createResp.Instance.ID, createVpcResp.ID)

	updateOpts := vpc.UpdateOpts{
		RouterID:        createResp.Instance.ID,
		VpcAttachmentID: createVpcResp.ID,
		Description:     pointerto.String(description + "_new"),
		Name:            vpcName + "_new",
	}

	t.Logf("Attempting to update vpc attachemnt")

	updateResp, err := vpc.Update(client, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, updateResp.Name)
	th.AssertEquals(t, updateOpts.Description, updateResp.Description)

	t.Logf("Attempting to update vpc attachemnt")
}

func TestVPCAttachmentsList(t *testing.T) {
	client, err := clients.NewERClient()
	th.AssertNoErr(t, err)

	_, err = vpc.List(client, vpc.ListOpts{})
	th.AssertNoErr(t, err)
}

func waitForVpcAttachmentsAvailable(client *golangsdk.ServiceClient, secs int, erId, vpcId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		vpcInstance, err := vpc.Get(client, erId, vpcId)
		if err != nil {
			return false, err
		}
		if vpcInstance.VpcAttachment.State == "available" {
			return true, nil
		}
		return false, nil
	})
}

func waitForVpcAttachmentsDeleted(client *golangsdk.ServiceClient, secs int, erId, vpcId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := vpc.Get(client, erId, vpcId)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
