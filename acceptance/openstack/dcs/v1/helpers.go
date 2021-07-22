package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/availablezones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/products"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createDCSInstance(t *testing.T, client *golangsdk.ServiceClient) *instances.Instance {
	t.Logf("Attempting to create DCSv1 instance")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if vpcID == "" || networkID == "" {
		t.Skip("OS_VPC_ID or OS_NETWORK_ID is missing but test requires using existing network")
	}
	if az == "" {
		az = "eu-de-01"
	}

	availabilityZone, err := availablezones.Get(client).Extract()
	th.AssertNoErr(t, err)

	product, err := products.Get(client).Extract()
	th.AssertNoErr(t, err)

	dcsName := tools.RandomString("dcs-instance-", 3)
	createOpts := instances.CreateOps{
		Name:           dcsName,
		Description:    "some test DCSv1 instance",
		Engine:         "Redis",
		EngineVersion:  "3.0",
		Capacity:       64,
		Password:       "Qwerty123!",
		VPCID:          vpcID,
		SubnetID:       networkID,
		AvailableZones: []string{availabilityZone.AvailableZones[0].ID},
		ProductID:      product.Products[0].ProductID,
	}

	instanceCreate, err := instances.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
}
