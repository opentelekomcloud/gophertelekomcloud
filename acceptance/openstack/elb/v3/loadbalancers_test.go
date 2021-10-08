package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLoadBalancerList(t *testing.T) {
	client, err := clients.NewElbV3Client(t)
	th.AssertNoErr(t, err)

	listOpts := loadbalancers.ListOpts{}
	loadbalancerPages, err := loadbalancers.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	loadbalancerList, err := loadbalancers.ExtractLoadbalancers(loadbalancerPages)
	th.AssertNoErr(t, err)

	for _, lb := range loadbalancerList {
		tools.PrintResource(t, lb)
	}
}

func TestLoadBalancerLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client(t)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create ELBv3 LoadBalancer")
	lbName := tools.RandomString("create-lb-", 3)
	adminStateUp := true

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	subnetID := clients.EnvOS.GetEnv("SUBNET_ID")
	if networkID == "" || vpcID == "" || subnetID == "" {
		t.Skip("OS_NETWORK_ID/OS_VPC_ID/OS_SUBNET_ID env vars are missing but LBv3 test requires")
	}

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-nl-01"
	}

	createOpts := loadbalancers.CreateOpts{
		Name:                 lbName,
		Description:          "some interesting loadbalancer",
		VipSubnetCidrID:      subnetID,
		VpcID:                vpcID,
		AvailabilityZoneList: []string{az},
		Tags: []tags.ResourceTag{
			{
				Key:   "gophertelekomcloud",
				Value: "loadbalancer",
			},
		},
		AdminStateUp: &adminStateUp,
		PublicIp: &loadbalancers.PublicIp{
			NetworkType: "5_bgp",
			Bandwidth: loadbalancers.Bandwidth{
				Name:       "elb_eip_traffic",
				Size:       10,
				ChargeMode: "traffic",
				ShareType:  "PER",
			},
		},
		ElbSubnetIDs: []string{networkID},
	}

	loadbalancer, err := loadbalancers.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		t.Logf("Attempting to delete ELBv3 LoadBalancer: %s", loadbalancer.ID)
		err := loadbalancers.Delete(client, loadbalancer.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 LoadBalancer: %s", loadbalancer.ID)
	}()
	th.AssertEquals(t, createOpts.Name, loadbalancer.Name)
	th.AssertEquals(t, createOpts.Description, loadbalancer.Description)
	t.Logf("Created ELBv3 LoadBalancer: %s", loadbalancer.ID)

	t.Logf("Attempting to update ELBv3 LoadBalancer: %s", loadbalancer.ID)
	lbName = tools.RandomString("update-lb-", 3)
	emptyDescription := ""
	updateOpts := loadbalancers.UpdateOpts{
		Name:        lbName,
		Description: &emptyDescription,
	}

	_, err = loadbalancers.Update(client, loadbalancer.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 LoadBalancer: %s", loadbalancer.ID)

	newLoadbalancer, err := loadbalancers.Get(client, loadbalancer.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newLoadbalancer.Name)
	th.AssertEquals(t, emptyDescription, newLoadbalancer.Description)
}
