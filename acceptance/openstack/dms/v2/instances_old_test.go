package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDmsList(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	listOpts := lifecycle.ListOpts{}
	dmsInstances, err := lifecycle.List(client, listOpts)
	th.AssertNoErr(t, err)
	for _, val := range dmsInstances.Instances {
		tools.PrintResource(t, val)
	}
}

func TestDmsOldLifeCycle(t *testing.T) {
	t.Skip("DMS Creation takes too long to complete")
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createOldDmsInstance(t, client)
	defer deleteDmsInstance(t, client, instanceID)

	dmsInstance, err := lifecycle.Get(client, instanceID)
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
	th.AssertEquals(t, delTopic.Topics[0].Name, dmsTopic)

	updateDmsInstance(t, client, instanceID)
	dmsInstance, err = lifecycle.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", dmsInstance.Description)
}

func createOldDmsInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create DMSv2 instance")
	dmsName := tools.RandomString("dms-acc-", 8)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but DMS test requires using existing network")
	}

	defaultSgID := openstack.DefaultSecurityGroup(t)
	az := getDmsInstanceAz(t, client)

	sslEnable := true

	createOpts := lifecycle.CreateDeprOpts{
		Name:            dmsName,
		Description:     "some interesting description",
		Engine:          "kafka",
		EngineVersion:   "2.3.0",
		StorageSpace:    600,
		Password:        "5ecuredPa55w0rd!",
		AccessUser:      "root",
		VpcID:           vpcID,
		SecurityGroupID: defaultSgID,
		SubnetID:        subnetID,
		AvailableZones:  []string{az},
		ProductID:       "00300-30308-0--0",
		SslEnable:       &sslEnable,
		StorageSpecCode: "dms.physical.storage.high",
	}
	dmsInstance, err := lifecycle.CreateDepr(client, createOpts)
	th.AssertNoErr(t, err)
	err = waitForInstanceAvailable(client, 600, dmsInstance.InstanceID)
	th.AssertNoErr(t, err)
	t.Logf("DMSv2 instance successfully created: %s", dmsInstance.InstanceID)

	return dmsInstance.InstanceID
}
