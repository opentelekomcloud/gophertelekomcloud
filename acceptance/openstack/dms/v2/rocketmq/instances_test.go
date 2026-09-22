package rocketmq

import (
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/availablezones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/products"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	rocketMQEngine        = "reliability"
	rocketMQEngineVersion = "5.x"
	rocketMQSingleNode    = "single.basic"
	rocketMQTargetStatus  = "RUNNING"
)

func TestRocketMQInstanceLifeCycle(t *testing.T) {
	if clients.EnvOS.GetEnv("RUN_DMS_ROCKETMQ") == "" {
		t.Skip("OS_RUN_DMS_ROCKETMQ env var is missing, RocketMQ instance creation takes ~30 minutes")
	}

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createRocketMQInstance(t, client)
	defer deleteRocketMQInstance(t, client, instanceID)

	instance, err := instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "some interesting description", instance.Description)
	th.AssertEquals(t, rocketMQEngineVersion, instance.EngineVersion)
	th.AssertEquals(t, rocketMQSingleNode, instance.Type)

	listed, err := instances.List(client, instances.ListOpts{
		Engine:     rocketMQEngine,
		InstanceID: instanceID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listed.Instances))
	th.AssertEquals(t, instanceID, listed.Instances[0].InstanceID)

	t.Logf("Attempting to update RocketMQ instance: %s", instanceID)
	err = instances.Update(client, instanceID, instances.UpdateOpts{
		Description: pointerto.String(""),
	})
	th.AssertNoErr(t, err)

	instance, err = instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", instance.Description)
}

func createRocketMQInstance(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create RocketMQ instance")

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || subnetID == "" {
		t.Skip("One of OS_VPC_ID or OS_NETWORK_ID env vars is missing but RocketMQ test requires using existing network")
	}
	var err error

	product := getRocketMQProduct(t, client)
	// RocketMQ flavors report an empty min_broker, unlike the Kafka ones.
	brokerNum := 1
	if product.Properties.MinBroker != "" {
		brokerNum, err = strconv.Atoi(product.Properties.MinBroker)
		th.AssertNoErr(t, err)
	}
	// Taken from the API rather than hardcoded: the API reports 300 GB as the
	// minimum for rocketmq.b1.large.1, while the docs claim the single-node
	// range starts at 100 GB.
	storagePerNode, err := strconv.Atoi(product.Properties.MinStoragePerNode)
	th.AssertNoErr(t, err)

	io := getRocketMQIO(t, product)

	instance, err := instances.Create(client, instances.CreateOpts{
		Name:            tools.RandomString("rocketmq-acc-", 8),
		Description:     "some interesting description",
		Engine:          rocketMQEngine,
		EngineVersion:   rocketMQEngineVersion,
		StorageSpace:    storagePerNode * brokerNum,
		StorageSpecCode: io.IOSpec,
		BrokerNum:       brokerNum,
		VpcID:           vpcID,
		SubnetID:        subnetID,
		SecurityGroupID: openstack.DefaultSecurityGroup(t),
		AvailableZones:  []string{getRocketMQAz(t, client, io.AvailableZones)},
		ProductID:       product.ProductId,
		SslEnable:       pointerto.Bool(false),
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, waitForRocketMQAvailable(client, 3600, instance.InstanceID))
	t.Logf("RocketMQ instance successfully created: %s", instance.InstanceID)

	return instance.InstanceID
}

func deleteRocketMQInstance(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	t.Logf("Attempting to delete RocketMQ instance: %s", instanceID)

	th.AssertNoErr(t, instances.Delete(client, instanceID))
	th.AssertNoErr(t, waitForRocketMQDelete(client, 1800, instanceID))

	t.Logf("RocketMQ instance deleted successfully: %s", instanceID)
}

// getRocketMQProduct returns the cheapest single-node RocketMQ flavor available.
func getRocketMQProduct(t *testing.T, client *golangsdk.ServiceClient) *products.EngineProduct {
	pd, err := products.List(client, products.ListOpts{Engine: rocketMQEngine})
	th.AssertNoErr(t, err)

	for _, p := range pd.Products {
		if p.Type == rocketMQSingleNode && len(p.IOS) > 0 {
			return &p
		}
	}

	t.Fatalf("no %s RocketMQ product found", rocketMQSingleNode)
	return nil
}

// getRocketMQIO picks a disk type the flavor accepts. The order of the ios list
// is not stable between calls, so the spec has to be selected by name rather
// than by taking the first entry: some of the listed disk types make the create
// call fail with HTTP 500, which made this test flaky when it used ios[0].
func getRocketMQIO(t *testing.T, product *products.EngineProduct) *products.EngineIOS {
	for _, preferred := range []string{"dms.physical.storage.ultra.v2", "dms.physical.storage.high.v2"} {
		for _, io := range product.IOS {
			if io.IOSpec == preferred && len(io.AvailableZones) > 0 {
				return &io
			}
		}
	}

	t.Fatalf("no supported disk type for product %s", product.ProductId)
	return nil
}

// getRocketMQAz maps one of the AZ codes supported by the flavor to its AZ ID.
func getRocketMQAz(t *testing.T, client *golangsdk.ServiceClient, codes []string) string {
	az, err := availablezones.Get(client)
	th.AssertNoErr(t, err)

	for _, code := range codes {
		for _, a := range az.AvailableZones {
			if a.Code == code && a.ResourceAvailability == "true" {
				return a.ID
			}
		}
	}

	t.Fatalf("none of the AZs %v is available", codes)
	return ""
}

func waitForRocketMQAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		instance, err := instances.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		return instance.Status == rocketMQTargetStatus, nil
	})
}

func waitForRocketMQDelete(client *golangsdk.ServiceClient, secs int, instanceID string) error {
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
