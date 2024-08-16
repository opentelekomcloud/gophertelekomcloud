package v2

import (
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/availablezones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/products"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	dmsEngine        = "kafka"
	dmsEngineVersion = "2.7"
	dmsUser          = "root"
	dmsUserPassword  = "5ecuredPa55w0rd!"

	kafkaClusterSmall = "cluster.small"

	dmsTargetStatus = "RUNNING"
)

func TestDmsList(t *testing.T) {
	client, err := clients.NewDmsV21Client()
	th.AssertNoErr(t, err)

	listOpts := instances.ListOpts{}
	dmsInstances, err := instances.List(client, listOpts)
	th.AssertNoErr(t, err)
	for _, val := range dmsInstances.Instances {
		tools.PrintResource(t, val)
	}
}

func TestDmsLifeCycle(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsInstance, err := instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "some interesting description", dmsInstance.Description)

	dmsTopic := createTopic(t, client, instanceID)

	err = updateDmsTopic(t, client, instanceID, dmsTopic)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 Topic updated")

	listTopics, err := topics.List(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listTopics.Topics[0].Name, dmsTopic)

	getTopic, err := topics.Get(client, instanceID, dmsTopic)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, dmsTopic, getTopic.Name)

	delTopic := deleteTopic(t, client, instanceID, dmsTopic)
	th.AssertEquals(t, delTopic.Topics[0].Name, dmsTopic)

	updateDmsInstance(t, client, instanceID)
	dmsInstance, err = instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", dmsInstance.Description)
}

func createDmsInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create DMSv2.1 instance")
	dmsName := tools.RandomString("dms-acc-", 8)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but DMS test requires using existing network")
	}

	defaultSgID := openstack.DefaultSecurityGroup(t)
	details := getDmsInstanceSpecification(t, client, dmsEngine)
	if details == nil {
		t.Fatalf("product type %s not found", kafkaClusterSmall)
	}
	az := getDmsInstanceAz(t, client)
	minBroker, _ := strconv.Atoi(details.Properties.MinBroker)
	storageSpace, _ := strconv.Atoi(details.Properties.MinStoragePerNode)
	storageSpec := details.IOS[0].IOSpec

	sslEnable := true

	createOpts := instances.CreateOpts{
		Name:            dmsName,
		Description:     "some interesting description",
		Engine:          dmsEngine,
		EngineVersion:   dmsEngineVersion,
		StorageSpace:    storageSpace * 3,
		StorageSpecCode: storageSpec,
		BrokerNum:       minBroker,
		ArchType:        details.ArchTypes[0],
		AccessUser:      dmsUser,
		Password:        dmsUserPassword,
		VpcID:           vpcID,
		SecurityGroupID: defaultSgID,
		SubnetID:        subnetID,
		AvailableZones:  []string{az},
		ProductID:       details.ProductId,
		SslEnable:       &sslEnable,
	}

	dmsInstance, err := instances.Create(client, createOpts)
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, dmsInstance.InstanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 instance successfully created: %s", dmsInstance.InstanceID)

	return dmsInstance.InstanceID
}

func deleteDmsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete DMSv2.1 instance: %s", instanceID)

	err := instances.Delete(client, instanceID)
	th.AssertNoErr(t, err)

	err = waitForInstanceDelete(client, 600, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 instance deleted successfully: %s", instanceID)
}

func updateDmsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to update DMSv2.1 instance: %s", instanceID)

	emptyDescription := ""
	updateOpts := instances.UpdateOpts{
		Description: &emptyDescription,
	}

	err := instances.Update(client, instanceID, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("DMSv2.1 instance updated successfully: %s", instanceID)
}

func getDmsInstanceSpecification(t *testing.T, client *golangsdk.ServiceClient, engine string) *products.EngineProduct {
	pd, err := products.Get(client, engine)
	th.AssertNoErr(t, err)

	for _, v := range pd.Products {
		if v.Type == kafkaClusterSmall {
			return &v
		}
	}

	return nil
}

func getDmsInstanceAz(t *testing.T, client *golangsdk.ServiceClient) string {
	az, err := availablezones.Get(client)
	th.AssertNoErr(t, err)

	return az.AvailableZones[0].ID
}

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		dmsInstance, err := instances.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		if dmsInstance.Status == dmsTargetStatus {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDelete(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := instances.Get(client, instanceID)
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
	t.Logf("Attempting to create DMSv2.1 Topic")
	topicName := tools.RandomString("dms-topic-", 8)

	createOpts := topics.CreateOpts{
		Name:             topicName,
		Partition:        10,
		Replication:      2,
		SyncReplication:  true,
		RetentionTime:    100,
		SyncMessageFlush: true,
	}
	dmsTopic, err := topics.Create(client, instanceId, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2.1 Topic successfully created: %s", dmsTopic.Name)

	return dmsTopic.Name
}

func updateDmsTopic(t *testing.T, client *golangsdk.ServiceClient, instanceId string, topicName string) error {
	t.Logf("Attempting to update DMSv2.1 Topic")
	partition := 12
	retention := 70
	updateOpts := topics.UpdateOpts{
		Topics: []topics.UpdateItem{
			{
				Name:          topicName,
				Partition:     &partition,
				RetentionTime: &retention,
			},
		},
	}
	return topics.Update(client, instanceId, updateOpts)
}

func deleteTopic(t *testing.T, client *golangsdk.ServiceClient, instanceId string, name string) *topics.DeleteResponse {
	t.Logf("Attempting to delete DMSv2.1 Topic")
	dmsTopic, err := topics.Delete(client, instanceId, []string{
		name,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, dmsTopic.Topics[0].Success)

	t.Logf("DMSv2.1 Topic successfully deleted")

	return dmsTopic
}
