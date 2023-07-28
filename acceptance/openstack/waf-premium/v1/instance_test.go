package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestWafPremiumInstanceWorkflow(t *testing.T) {
	if os.Getenv("RUN_WAFD_INSTANCE_WORKFLOW") == "" {
		t.Skip("too slow to run in zuul")
	}
	region := os.Getenv("OS_REGION_NAME")
	az := os.Getenv("OS_AVAILABILITY_ZONE")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	if vpcID == "" && subnetID == "" && region == "" && az == "" {
		t.Skip("OS_REGION_NAME, OS_AVAILABILITY_ZONE, OS_VPC_ID and OS_NETWORK_ID env vars is required for this test")
	}

	architecture := "x86"
	if region == "eu-ch2" {
		architecture = "x86_64"
	}

	client, err := getWafdClient(t, region)

	opts := instances.CreateOpts{
		Count:            1,
		Region:           region,
		AvailabilityZone: az,
		Architecture:     architecture,
		InstanceName:     tools.RandomString("waf-instance-", 3),
		Specification:    "waf.instance.enterprise",
		Flavor:           "s3.2xlarge.2",
		VpcId:            vpcID,
		SubnetId:         subnetID,
		SecurityGroupsId: []string{openstack.DefaultSecurityGroup(t)},
	}

	t.Logf("Attempting to create WAF premium instance")
	inst, err := instances.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created WAF instance: %s", inst.Instances[0].Id)
	instanceId := inst.Instances[0].Id

	th.AssertNoErr(t, waitForInstanceToBeCreated(client, 600, instanceId))

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium instance: %s", instanceId)
		th.AssertNoErr(t, instances.Delete(client, instanceId))
		th.AssertNoErr(t, waitForInstanceToBeDeleted(client, 600, instanceId))
		t.Logf("Deleted WAF Premium instance: %s", instanceId)
	})

	instance, err := instances.Get(client, instanceId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, instance.ID, instanceId)
	th.AssertEquals(t, instance.ResourceSpecification, "waf.instance.enterprise")

	instancesList, err := instances.List(client, instances.ListOpts{})
	th.AssertNoErr(t, err)

	if len(instancesList) < 1 {
		t.Fatal("empty WAF instances list")
	}
	updatedName := tools.RandomString("waf-instance-updated-", 3)
	instanceUpdated, err := instances.Update(client, instanceId, instances.UpdateOpts{
		Name: updatedName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, instanceUpdated.Name, updatedName)
}
