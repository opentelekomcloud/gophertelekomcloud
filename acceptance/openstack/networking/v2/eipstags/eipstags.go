package eipstags

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/eipstags"
)

func CreateEip(t *testing.T, clientV1 *golangsdk.ServiceClient) eips.PublicIp {
	t.Logf("Attempting to create VPC EIPv1")
	createOps := eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			Name:       "test",
			Size:       100,
			ShareType:  "PER",
			ChargeMode: "traffic",
		},
	}
	eip, err := eips.Apply(clientV1, createOps).Extract()
	if err != nil {
		t.Fatalf("Unable to create VPC EIPv1: %s", err)
	}

	t.Logf("Created VPC EIPv1: %s", eip.ID)
	return eip
}

func DeleteEip(t *testing.T, clientV1 *golangsdk.ServiceClient, eipID string) {
	t.Logf("Attempting to delete VPC EIPv1: %s", eipID)
	if err := eips.Delete(clientV1, eipID).ExtractErr(); err != nil {
		t.Fatalf("Unable to delete VPC EIPv1: %s", err)
	}
	t.Logf("Deleted VPC EIPv1: %s", eipID)
}

func CreateTag(t *testing.T, clientV2 *golangsdk.ServiceClient, eipID string, tagKey string) {
	t.Logf("Attempting to create tag for VPC EIPv1")
	createOps := eipstags.CreateOpts{
		Tag: eipstags.Tag{
			Key:   tagKey,
			Value: "kuh",
		},
	}
	err := eipstags.Create(clientV2, createOps, eipID).ExtractErr()
	if err != nil {
		t.Fatal("Unable to create tag for VPC EIPv1")
	}
	t.Logf("Created tag for VPC EIPv1: %s", eipID)
}

func DeleteTag(t *testing.T, clientV2 *golangsdk.ServiceClient, eipID string, tagKey string) {
	t.Logf("Attempting to delete tag %s for VPC EIPv1: %s", tagKey, eipID)
	if err := eipstags.Delete(clientV2, eipID, tagKey).ExtractErr(); err != nil {
		t.Fatal("Unable to delete tag for VPC EIPv1")
	}
	t.Logf("Deleted tag for VPC EIPv1")
}

func CreateTags(t *testing.T, clientV2 *golangsdk.ServiceClient) {

}

func DeleteTags(t *testing.T, clientV2 *golangsdk.ServiceClient) {

}
