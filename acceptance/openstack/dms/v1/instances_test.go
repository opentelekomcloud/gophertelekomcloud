package v1

import (
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/availablezones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/products"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDmsList(t *testing.T) {
	client, err := clients.NewDmsV1Client()
	th.AssertNoErr(t, err)

	listOpts := instances.ListDmsInstanceOpts{}
	dmsAllPages, err := instances.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	dmsInstances, err := instances.ExtractDmsInstances(dmsAllPages)
	th.AssertNoErr(t, err)
	for _, val := range dmsInstances.Instances {
		tools.PrintResource(t, val)
	}
}

func TestDmsLifeCycle(t *testing.T) {
	client, err := clients.NewDmsV1Client()
	th.AssertNoErr(t, err)

	instanceID := createDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsInstance, err := instances.Get(client, instanceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "some interesting description", dmsInstance.Description)

	updateDmsInstance(t, client, instanceID)
	dmsInstance, err = instances.Get(client, instanceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", dmsInstance.Description)
}

func createDmsInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create DMSv1 instance")
	dmsName := tools.RandomString("dms-acc-", 8)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but DMS test requires using existing network")
	}

	defaultSgID := openstack.DefaultSecurityGroup(t)
	details := getDmsInstanceSpecification(t, client)
	az := getDmsInstanceAz(t, client)
	partitionNum, _ := strconv.Atoi(details.PartitionNum)
	storage, _ := strconv.Atoi(details.Storage)

	createOpts := instances.CreateOpts{
		Name:            dmsName,
		Description:     "some interesting description",
		Engine:          "kafka",
		EngineVersion:   "2.3.0",
		StorageSpace:    storage,
		Password:        "5ecuredPa55w0rd!",
		AccessUser:      "rgyrbu",
		VpcID:           vpcID,
		SecurityGroupID: defaultSgID,
		SubnetID:        subnetID,
		AvailableZones:  []string{az},
		ProductID:       details.ProductID,
		PartitionNum:    partitionNum,
		SslEnable:       true,
		Specification:   details.VMSpecification,
		StorageSpecCode: details.IOs[0].StorageSpecCode,
	}
	dmsInstance, err := instances.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, dmsInstance.InstanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv1 instance successfully created: %s", dmsInstance.InstanceID)

	return dmsInstance.InstanceID
}

func deleteDmsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete DMSv1 instance: %s", instanceID)

	err := instances.Delete(client, instanceID).ExtractErr()
	th.AssertNoErr(t, err)

	err = waitForInstanceDelete(client, 600, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv1 instance deleted successfully: %s", instanceID)
}

func updateDmsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to update DMSv1 instance: %s", instanceID)

	emptyDescription := ""
	updateOpts := instances.UpdateOpts{
		Description: &emptyDescription,
	}

	err := instances.Update(client, instanceID, updateOpts).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("DMSv1 instance updated successfully: %s", instanceID)
}

func getDmsInstanceSpecification(t *testing.T, client *golangsdk.ServiceClient) products.Detail {
	v, err := products.Get(client, "kafka").Extract()
	th.AssertNoErr(t, err)
	productList := v.Hourly

	var filteredPd []products.Detail
	for _, pd := range productList {
		if pd.Version != "2.3.0" {
			continue
		}
		for _, value := range pd.Values {
			if value.Name != "cluster" {
				continue
			}

			filteredPd = append(filteredPd, value.Details...)
		}
	}

	return filteredPd[0]
}

func getDmsInstanceAz(t *testing.T, client *golangsdk.ServiceClient) string {
	az, err := availablezones.Get(client).Extract()
	th.AssertNoErr(t, err)

	return az.AvailableZones[0].ID
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		dmsInstance, err := instances.Get(client, instanceID).Extract()
		if err != nil {
			return false, err
		}
		if dmsInstance.Status == "RUNNING" {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDelete(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := instances.Get(client, instanceID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}
