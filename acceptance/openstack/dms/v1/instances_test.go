package v1

import (
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/others"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/products"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/topics"
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

	dmsTopic := createTopic(t, client, instanceID)

	getTopic, err := topics.Get(client, instanceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, dmsTopic, getTopic.Topics[0].Name)

	delTopic := deleteTopic(t, client, instanceID, dmsTopic)
	th.AssertEquals(t, delTopic[0].Name, dmsTopic)

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
		AccessUser:      "root",
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

	err := instances.Delete(client, instanceID).Err
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

	err := instances.Update(client, instanceID, updateOpts).Err
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
	az, err := others.ListAvailableZones(client)
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

func createTopic(t *testing.T, client *golangsdk.ServiceClient, instanceId string) string {
	t.Logf("Attempting to create DMSv1 Topic")
	topicName := tools.RandomString("dms-topic-", 8)

	createOpts := topics.CreateOpts{
		Name:             topicName,
		Partition:        10,
		Replication:      2,
		SyncReplication:  true,
		RetentionTime:    100,
		SyncMessageFlush: true,
	}
	dmsTopic, err := topics.Create(client, createOpts, instanceId).Extract()
	th.AssertNoErr(t, err)
	t.Logf("DMSv1 Topic successfully created: %s", dmsTopic.Name)

	return dmsTopic.Name
}

func deleteTopic(t *testing.T, client *golangsdk.ServiceClient, instanceId string, name string) []topics.TopicDelete {
	t.Logf("Attempting to delete DMSv1 Topic")

	deleteOpts := topics.DeleteOpts{
		Topics: []string{
			name,
		},
	}
	dmsTopic, err := topics.Delete(client, deleteOpts, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, dmsTopic[0].Success, true)

	getTopic, err := topics.Get(client, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getTopic.Size, 0)

	t.Logf("DMSv1 Topic successfully deleted")

	return dmsTopic
}
