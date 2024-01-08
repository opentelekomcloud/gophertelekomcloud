package v2

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/gateway"
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
		LoadbalancerProvider: "elb",
		Description:          "All Work And No Play Makes Jack A Dull Boy",
	}
	createResp, err := gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, createResp)

	th.AssertNoErr(t, WaitForJob(client, createResp.InstanceID, 1800))

	updateOpts := gateway.UpdateOpts{
		Description:   "it's not getting better",
		MaintainBegin: "22:00:00",
		MaintainEnd:   "02:00:00",
		InstanceName:  createOpts.InstanceName + "updated",
		ID:            createResp.InstanceID,
	}

	_, err = gateway.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	getResp, err := gateway.Get(client, createResp.InstanceID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)

	th.AssertNoErr(t, gateway.EnableEIP(client, gateway.EipOpts{
		ID:                    createResp.InstanceID,
		BandwidthChargingMode: "traffic",
		BandwidthSize:         "5",
	}))

	th.AssertNoErr(t, gateway.UpdateEIP(client, gateway.EipOpts{
		ID:                    createResp.InstanceID,
		BandwidthChargingMode: "traffic",
		BandwidthSize:         "10",
	}))

	th.AssertNoErr(t, gateway.DisableEIP(client, createResp.InstanceID))

	th.AssertNoErr(t, gateway.Delete(client, createResp.InstanceID))
}

func TestGatewayList(t *testing.T) {
	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	listResp, err := gateway.List(client, gateway.ListOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listResp)
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
