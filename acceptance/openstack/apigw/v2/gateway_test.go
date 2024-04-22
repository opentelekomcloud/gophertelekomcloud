package v2

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGatewayLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")

	if vpcID == "" || subnetID == "" {
		t.Skip("Both `VPC_ID` and `NETWORK_ID` need to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	createOpts := gateway.CreateOpts{
		VpcID:        vpcID,
		SubnetID:     subnetID,
		InstanceName: tools.RandomString("test-gateway-", 5),
		SpecID:       "PROFESSIONAL",
		SecGroupID:   openstack.DefaultSecurityGroup(t),
		AvailableZoneIDs: []string{
			"eu-de-01",
			"eu-de-02",
		},
		LoadBalancerProvider:         "elb",
		Description:                  "All Work And No Play Makes Jack A Dull Boy",
		IngressBandwidthChargingMode: "bandwidth",
		IngressBandwidthSize:         pointerto.Int(5),
		// Tags: []tags.ResourceTag{
		// 	{
		// 		Key:   "TestKey",
		// 		Value: "TestValue",
		// 	},
		// 	{
		// 		Key:   "empty",
		// 		Value: "",
		// 	},
		// },
	}
	t.Logf("Attempting to CREATE APIGW Gateway")
	createResp, err := gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, WaitForJob(client, createResp.InstanceID, 1800))
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW Gateway: %s", createResp.InstanceID)
		th.AssertNoErr(t, gateway.Delete(client, createResp.InstanceID))
	})

	updateOpts := gateway.UpdateOpts{
		Description:   "it's not getting better",
		MaintainBegin: "22:00:00",
		MaintainEnd:   "02:00:00",
		InstanceName:  createOpts.InstanceName + "updated",
		ID:            createResp.InstanceID,
	}

	t.Logf("Attempting to UPDATE APIGW Gateway: %s", createResp.InstanceID)
	_, err = gateway.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to GET APIGW Gateway: %s", createResp.InstanceID)
	getResp, err := gateway.Get(client, createResp.InstanceID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)

	// API not published yet
	// t.Logf("Attempting to UPDATE APIGW Gateway tags: %s", createResp.InstanceID)
	// err = gateway.UpdateTags(client, &gateway.TagsUpdateOpts{
	// 	InstanceId: createResp.InstanceID,
	// 	Action:     "create",
	// 	Tags: []tags.ResourceTag{
	// 		{
	// 			Key:   "NewKey",
	// 			Value: "NewValue",
	// 		},
	// 	},
	// })
	// th.AssertNoErr(t, err)

	// t.Logf("Attempting to GET APIGW Gateway tags list: %s", createResp.InstanceID)
	// getTags, err := gateway.GetTags(client, createResp.InstanceID)
	// th.AssertNoErr(t, err)
	// th.CheckEquals(t, 3, len(getTags))

	// t.Logf("Attempting to DISABLE APIGW Gateway ELB ingress access: %s", createResp.InstanceID)
	// err = gateway.DisableElbIngressAccess(client, createResp.InstanceID)
	// th.AssertNoErr(t, err)
	//
	// optsIngressAccess := gateway.ElbIngressAccessOpts{
	// 	InstanceId:                  createResp.InstanceID,
	// 	IngressBandwithSize:         10,
	// 	IngressBandwithChargingMode: "traffic",
	// }
	// t.Logf("Attempting to ENABLE APIGW Gateway ELB ingress access: %s", createResp.InstanceID)
	// updateIngress, err := gateway.EnableElbIngressAccess(client, optsIngressAccess)
	// th.AssertNoErr(t, WaitForJob(client, updateIngress.InstanceID, 1800))

	t.Logf("Attempting to ENABLE APIGW Gateway EIP: %s", createResp.InstanceID)
	th.AssertNoErr(t, gateway.EnableEIP(client, gateway.EipOpts{
		ID:                    createResp.InstanceID,
		BandwidthChargingMode: "traffic",
		BandwidthSize:         "5",
	}))

	t.Logf("Attempting to UPDATE APIGW Gateway EIP: %s", createResp.InstanceID)
	th.AssertNoErr(t, gateway.UpdateEIP(client, gateway.EipOpts{
		ID:                    createResp.InstanceID,
		BandwidthChargingMode: "traffic",
		BandwidthSize:         "10",
	}))

	t.Logf("Attempting to DISABLE APIGW Gateway EIP: %s", createResp.InstanceID)
	th.AssertNoErr(t, gateway.DisableEIP(client, createResp.InstanceID))
}

func TestGatewayList(t *testing.T) {
	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	_, err = gateway.List(client, gateway.ListOpts{})
	th.AssertNoErr(t, err)
}

func WaitForJob(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := gateway.QueryProgress(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "success" {
			return true, nil
		}

		if current.Status == "failed" {
			return false, fmt.Errorf("job failed: %s", current.ErrorMsg)
		}

		return false, nil
	})
}
