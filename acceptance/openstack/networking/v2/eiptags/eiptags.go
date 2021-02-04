package eiptags

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/eiptags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateEip(t *testing.T, clientV1 *golangsdk.ServiceClient) eips.PublicIp {
	t.Logf("Attempting to create VPC EIPv1")
	createOpts := eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			Name:       "test-bandwidth",
			Size:       100,
			ShareType:  "PER",
			ChargeMode: "traffic",
		},
	}
	eip, err := eips.Apply(clientV1, createOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created VPC EIPv1: %s", eip.ID)
	return eip
}

func DeleteEip(t *testing.T, clientV1 *golangsdk.ServiceClient, eipID string) {
	t.Logf("Attempting to delete VPC EIPv1: %s", eipID)
	err := eips.Delete(clientV1, eipID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Deleted VPC EIPv1: %s", eipID)
}

func CreateTag(t *testing.T, clientV2 *golangsdk.ServiceClient, eipID string, tagKey string) {
	t.Logf("Attempting to create tag for VPC EIPv1")
	createOpts := eiptags.CreateOpts{
		Tag: eiptags.Tag{
			Key:   tagKey,
			Value: "value1",
		},
	}
	err := eiptags.Create(clientV2, createOpts, eipID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Created tag for VPC EIPv1: %s", eipID)
}

func DeleteTag(t *testing.T, clientV2 *golangsdk.ServiceClient, eipID string, tagKey string) {
	t.Logf("Attempting to delete tag %s for VPC EIPv1: %s", tagKey, eipID)
	err := eiptags.Delete(clientV2, eipID, tagKey).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Deleted tag for VPC EIPv1")
}

func CreateTags(t *testing.T, clientV2 *golangsdk.ServiceClient, eipID string, tagKeys []string) {
	t.Logf("Attempting to create tags for VPC EIPv1")
	createOpts := eiptags.BatchActionOpts{
		Tags: []eiptags.Tag{
			{
				Key:   tagKeys[0],
				Value: "value2",
			},
			{
				Key:   tagKeys[1],
				Value: "value3",
			},
		},
		Action: "create",
	}
	err := eiptags.Action(clientV2, createOpts, eipID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Created tags for VPC EIPv1: %s", eipID)
}

func DeleteTags(t *testing.T, clientV2 *golangsdk.ServiceClient, eipID string, tagKeys []string) {
	t.Logf("Attempting to delete tags for VPC EIPv1: %s", eipID)
	deleteOpts := eiptags.BatchActionOpts{
		Tags: []eiptags.Tag{
			{
				Key:   tagKeys[0],
				Value: "value2",
			},
			{
				Key:   tagKeys[1],
				Value: "value3",
			},
		},
		Action: "delete",
	}
	err := eiptags.Action(clientV2, deleteOpts, eipID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Deleted tags for VPC EIPv1")
}
