package v2

import (
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/availablezones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/products"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDmsList(t *testing.T) {
	client, err := clients.NewDmsV2Client()
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

	err = instances.ChangePassword(client, instanceID, instances.PasswordOpts{
		NewPassword: "5ecuredPa55w0rd!-not",
	})
	th.AssertNoErr(t, err)
	t.Logf("DMSv2 Instance password updated")

	// updateDMScrossVpc(t, client, instanceID)
	dmsTopic := createTopic(t, client, instanceID)

	err = updateDmsTopic(t, client, instanceID, dmsTopic)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2 Topic updated")

	listTopics, err := topics.List(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listTopics.Topics[0].Name, dmsTopic)

	getTopic, err := topics.Get(client, instanceID, dmsTopic)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, dmsTopic, getTopic.Name)

	delTopic := deleteTopic(t, client, instanceID, dmsTopic)
	th.AssertEquals(t, delTopic.Name, dmsTopic)

	updateDmsInstance(t, client, instanceID)
	dmsInstance, err = instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", dmsInstance.Description)
}

func createDmsInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create DMSv2 instance")
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

	sslEnable := true

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
		SslEnable:       &sslEnable,
		Specification:   details.VMSpecification,
		StorageSpecCode: details.IOs[0].StorageSpecCode,
	}
	dmsInstance, err := instances.Create(client, createOpts)
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, dmsInstance.InstanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2 instance successfully created: %s", dmsInstance.InstanceID)

	return dmsInstance.InstanceID
}

func deleteDmsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete DMSv2 instance: %s", instanceID)

	err := instances.Delete(client, instanceID)
	th.AssertNoErr(t, err)

	err = waitForInstanceDelete(client, 600, instanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv1 instance deleted successfully: %s", instanceID)
}

func updateDmsInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to update DMSv2 instance: %s", instanceID)

	emptyDescription := ""
	updateOpts := instances.UpdateOpts{
		Description: &emptyDescription,
	}

	_, err := instances.Update(client, instanceID, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("DMSv2 instance updated successfully: %s", instanceID)
}

// func updateDMScrossVpc(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
// 	t.Logf("Attempting to modify crossVPC for DMSv2 instance: %s", instanceID)
//
// 	crossVpcOpts := instances.CrossVpcUpdateOpts{
// 		Contents: map[string]string{
// 			"192.168.1.27":  "192.168.1.27",
// 			"192.168.1.238": "192.168.1.238",
// 			"192.168.1.11":  "192.168.1.12",
// 		},
// 	}
//
// 	crossVpc, err := instances.UpdateCrossVpc(client, instanceID, crossVpcOpts)
// 	th.AssertNoErr(t, err)
// 	th.AssertEquals(t, true, crossVpc.Success)
//
// 	t.Logf("DMSv2 instance crossVPC modified successfully: %s", instanceID)
// }

func getDmsInstanceSpecification(t *testing.T, client *golangsdk.ServiceClient) products.Detail {
	v, err := products.Get(client, "kafka")
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
		if dmsInstance.Status == "RUNNING" {
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
	t.Logf("Attempting to create DMSv2 Topic")
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
	t.Logf("DMSv2 Topic successfully created: %s", dmsTopic.Name)

	return dmsTopic.Name
}

func updateDmsTopic(t *testing.T, client *golangsdk.ServiceClient, instanceId string, topicName string) error {
	t.Logf("Attempting to update DMSv2 Topic")
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
	t.Logf("Attempting to delete DMSv2 Topic")
	dmsTopic, err := topics.Delete(client, instanceId, []string{
		name,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, dmsTopic.Success, true)

	t.Logf("DMSv2 Topic successfully deleted")

	return dmsTopic
}
