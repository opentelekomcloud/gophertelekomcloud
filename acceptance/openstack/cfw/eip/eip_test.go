package eip

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	common "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cfw"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	cfweip "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/eip"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	managementv2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v2/management"
	eips "github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCFWEIPLifecycle(t *testing.T) {
	clientv1, err := clients.NewCFWV1Client()
	th.AssertNoErr(t, err)
	clientv2, err := clients.NewCFWV2Client()
	th.AssertNoErr(t, err)
	clientv3, err := clients.NewCFWV3Client()
	th.AssertNoErr(t, err)
	clientNetV1, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	eip1, err := createEip(clientNetV1)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = eips.Delete(clientNetV1, eip1.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	instanceName := tools.RandomString("test-acc-firewall-", 3)
	createOpts := managementv2.CreateOpts{
		Name: instanceName,
		Flavor: managementv2.CreateFlavor{
			Version: "standard",
		},
		ChargeInfo: managementv2.ChargeInfo{
			ChargeMode: "postPaid",
		},
	}
	createResp, err := managementv2.Create(clientv2, createOpts)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, common.WaitForJobCompleted(clientv3, 600, 5, createResp.JobID))
	instanceId := createResp.JobID
	t.Cleanup(func() {
		_, err = managementv2.Delete(clientv2, instanceId)
		th.AssertNoErr(t, err)
	})

	firewall, err := managementv1.Get(clientv1, instanceId, 0)
	th.AssertNoErr(t, err)

	changeEipProtectionOpts := cfweip.ChangeEIPProtectionOpts{
		ObjectID: firewall.ProtectObjects[0].ObjectID,
		Status:   0,
		IPInfos: []cfweip.IPInfo{
			{
				ID:       eip1.ID,
				PublicIP: eip1.PublicAddress,
			},
		},
	}
	_, err = cfweip.ChangeEIPProtection(clientv1, instanceId, changeEipProtectionOpts)
	th.AssertNoErr(t, err)

	queryOpts := cfweip.ListOpts{
		ObjectID: firewall.ProtectObjects[0].ObjectID,
	}
	eipResourceList, err := cfweip.List(clientv1, queryOpts)
	th.AssertNoErr(t, err)

	for _, eipResource := range eipResourceList {
		if eipResource.ID == eip1.ID {
			th.AssertEquals(t, 0, eipResource.Status)
		}
	}
}

func createEip(clientNet *golangsdk.ServiceClient) (*eips.PublicIp, error) {
	eipName := tools.RandomString("test-acc-eip-", 3)
	eip, err := eips.Apply(clientNet, eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Name: eipName,
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			Name:       "bandwidth-" + eipName,
			Size:       1,
			ShareType:  "PER",
			ChargeMode: "traffic"},
	}).Extract()
	if err != nil {
		return nil, err
	}
	return eip, nil
}
