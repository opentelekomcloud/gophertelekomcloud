package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/acl"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/api"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/env"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAclPolicyLifecycle(t *testing.T) {
	apigw := os.Getenv("APIGW_RUN")

	if apigw == "" {
		t.Skip("`APIGW_RUN`needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	gw := CreateGateway(client, t)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW Gateway: %s", gw.InstanceID)
		th.AssertNoErr(t, gateway.Delete(client, gw.InstanceID))
	})

	name := tools.RandomString("test_acl_policy_", 5)
	policy := CreateAclPolicy(client, t, gw.InstanceID, name)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW Acl Policy: %s", policy.ID)
		th.AssertNoErr(t, acl.Delete(client, gw.InstanceID, policy.ID))
	})

	updateOpts := acl.CreateOpts{
		GatewayID:  gw.InstanceID,
		Name:       name + "_updated",
		Type:       "PERMIT",
		Value:      "192.168.1.50,192.168.10.10",
		EntityType: "IP",
	}
	t.Logf("Attempting to UPDATE APIGW Acl Policy: %s", policy.ID)
	policyUpdate, err := acl.Update(client, policy.ID, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to RETRIEVE APIGW Acl Policy: %s", policy.ID)
	policyGet, err := acl.Get(client, gw.InstanceID, policyUpdate.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policyGet)

	t.Logf("Attempting to RETRIEVE APIGW Acl Policies")
	policyList, err := acl.List(client, acl.ListOpts{
		GatewayID: gw.InstanceID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, policyList)
}

func TestAclPolicyBinding(t *testing.T) {
	apigw := os.Getenv("APIGW_RUN")

	if apigw == "" {
		t.Skip("`APIGW_RUN`needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	gw := CreateGateway(client, t)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW Gateway: %s", gw.InstanceID)
		th.AssertNoErr(t, gateway.Delete(client, gw.InstanceID))
	})

	gatewayID := gw.InstanceID

	name := tools.RandomString("test_acl_policy_", 5)
	policy := CreateAclPolicy(client, t, gatewayID, name)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW Acl Policy: %s", policy.ID)
		th.AssertNoErr(t, acl.Delete(client, gatewayID, policy.ID))
	})

	t.Logf("Attempting to CREATE APIGW Group")
	apigwGroup := CreateGroup(client, t, gatewayID)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW Group: %s", apigwGroup.ID)
		th.AssertNoErr(t, group.Delete(client, gatewayID, apigwGroup.ID))
	})

	groupID := apigwGroup.ID

	t.Logf("Attempting to CREATE APIGW API")
	_, createAPIResp := CreateAPI(client, t, gatewayID, groupID)

	t.Cleanup(func() {
		t.Logf("Attempting to DELETE APIGW API: %s", createAPIResp.ID)
		th.AssertNoErr(t, api.Delete(client, gatewayID, createAPIResp.ID))
	})

	t.Logf("Attempting to CREATE APIGW ENVIRONMENT")
	envResp := CreateEnv(client, t, gatewayID)

	t.Cleanup(func() {
		manageOpts := api.ManageOpts{
			GatewayID:   gatewayID,
			Action:      "offline",
			EnvID:       envResp.ID,
			ApiID:       createAPIResp.ID,
			Description: "test-api-publish",
		}

		_, err = api.ManageApi(client, manageOpts)
		th.AssertNoErr(t, err)
		t.Logf("Attempting to DELETE APIGW ENVIRONMENT: %s", envResp.ID)
		th.AssertNoErr(t, env.Delete(client, gatewayID, envResp.ID))
	})

	manageOpts := api.ManageOpts{
		GatewayID:   gatewayID,
		Action:      "online",
		EnvID:       envResp.ID,
		ApiID:       createAPIResp.ID,
		Description: "test-api-publish",
	}

	t.Logf("Attempting to PUBLISH APIGW API")
	publishAPI, err := api.ManageApi(client, manageOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to RETRIEVE APIGW unbounded APIs")
	listUnbound, err := acl.ListAPIUnoundPolicy(client, acl.ListBoundOpts{
		GatewayID: gatewayID,
		ID:        policy.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listUnbound)

	bindOpts := acl.BindOpts{
		GatewayID: gatewayID,
		PolicyID:  policy.ID,
		PublishIds: []string{
			publishAPI.PublishID,
		},
	}
	t.Logf("Attempting to BIND APIGW Acl Policy to API")
	bindResp, err := acl.BindPolicy(client, bindOpts)
	t.Cleanup(func() {
		t.Logf("Attempting to UNBIND APIGW Acl Policy from API")
		th.AssertNoErr(t, acl.UnbindPolicy(client, gatewayID, bindResp[0].ID))
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, bindResp)

	t.Logf("Attempting to RETRIEVE APIGW bounded to policy APIs")
	listBound, err := acl.ListAPIBoundPolicy(client, acl.ListBoundOpts{
		GatewayID: gatewayID,
		ID:        policy.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listBound)

	t.Logf("Attempting to RETRIEVE APIGW bounded APIs policies")
	listBoundPolicies, err := acl.ListBoundPolicies(client, acl.ListBindingOpts{
		GatewayID: gatewayID,
		PolicyId:  policy.ID,
		ApiId:     createAPIResp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listBoundPolicies)
}

func CreateAclPolicy(client *golangsdk.ServiceClient, t *testing.T, gatewayID, name string) *acl.AclResp {
	createOpts := acl.CreateOpts{
		GatewayID:  gatewayID,
		Name:       name,
		Type:       "PERMIT",
		Value:      "192.168.1.5,192.168.10.1",
		EntityType: "IP",
	}
	t.Logf("Attempting to CREATE APIGW Acl Policy")
	createResp, err := acl.Create(client, createOpts)
	th.AssertNoErr(t, err)
	return createResp
}

func CreateGateway(client *golangsdk.ServiceClient, t *testing.T) *gateway.GatewayResp {
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
		SpecID:       "BASIC",
		SecGroupID:   openstack.DefaultSecurityGroup(t),
		AvailableZoneIDs: []string{
			"eu-de-01",
			"eu-de-02",
		},
		LoadBalancerProvider:         "elb",
		Description:                  "All Work And No Play Makes Jack A Dull Boy",
		IngressBandwidthChargingMode: "bandwidth",
		IngressBandwidthSize:         pointerto.Int(5),
	}
	t.Logf("Attempting to CREATE APIGW Gateway")
	createResp, err := gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, WaitForJob(client, createResp.InstanceID, 1800))
	return createResp
}
