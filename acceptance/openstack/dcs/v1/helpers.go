package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/availablezones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/lifecycle"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/products"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createDCSInstance(t *testing.T, client *golangsdk.ServiceClient) *lifecycle.Instance {
	t.Logf("Attempting to create DCSv1 instance")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || networkID == "" {
		t.Skip("OS_VPC_ID or OS_NETWORK_ID is missing but test requires using existing network")
	}

	availabilityZone, err := availablezones.Get(client).Extract()
	th.AssertNoErr(t, err)
	var az string
	for _, v := range availabilityZone.AvailableZones {
		if v.ResourceAvailability != "true" {
			continue
		}
		az = v.ID
	}
	if az == "" {
		t.Skip("Availability Zone ID wasn't found")
	}

	productList, err := products.Get(client).Extract()
	th.AssertNoErr(t, err)

	var productID string
	for _, v := range productList.Products {
		if v.SpecCode == "redis.ha.xu1.tiny.r2.128" {
			productID = v.ProductID
		}
	}
	if productID == "" {
		t.Skip("Product ID wasn't found")
	}

	defaultSG := openstack.DefaultSecurityGroup(t)
	dcsName := tools.RandomString("dcs-instance-", 3)
	createOpts := lifecycle.CreateOps{
		Name:            dcsName,
		Description:     "some test DCSv1 instance",
		Engine:          "Redis",
		EngineVersion:   "5.0",
		Capacity:        0.125,
		Password:        "Qwerty123!",
		VPCID:           vpcID,
		SubnetID:        networkID,
		AvailableZones:  []string{az},
		ProductID:       productID,
		SecurityGroupID: defaultSG,
	}
	dcsInstanceCreate, err := lifecycle.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 600, dcsInstanceCreate.InstanceID)
	th.AssertNoErr(t, err)

	t.Logf("DCSv1 instance successfully created: %s", dcsInstanceCreate.InstanceID)

	dcsInstance, err := lifecycle.Get(client, dcsInstanceCreate.InstanceID).Extract()
	th.AssertNoErr(t, err)

	return dcsInstance
}

func deleteDCSInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete DCSv1 instance: %s", instanceID)

	err := lifecycle.Delete(client, instanceID).ExtractErr()
	th.AssertNoErr(t, err)

	err = waitForInstanceDeleted(client, 600, instanceID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, err)
	t.Logf("Deleted DCSv1 instance: %s", instanceID)
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		dcsInstances, err := lifecycle.Get(client, instanceID).Extract()
		if err != nil {
			return false, err
		}
		if dcsInstances.Status == "RUNNING" {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDeleted(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := lifecycle.Get(client, instanceID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
