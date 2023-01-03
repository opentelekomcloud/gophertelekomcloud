package v1

import (
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const timeout = 1200

func TestClusterWorkflow(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")

	if vpcID == "" || subnetID == "" {
		t.Skip("Both `VPC_ID` and `NETWORK_ID` need to be defined")
	}

	sgID := openstack.DefaultSecurityGroup(t)

	opts := clusters.CreateOpts{
		Name: tools.RandomString("css-cluster-", 4),
		Instance: &clusters.InstanceSpec{
			Flavor: "css.medium.8",

			Volume: &clusters.Volume{
				Type: "COMMON",
				Size: 40,
			},
			Nics: &clusters.Nics{
				VpcID:           vpcID,
				SubnetID:        subnetID,
				SecurityGroupID: sgID,
			},
			AvailabilityZone: "eu-de-02",
		},
		InstanceNum: 1,
		DiskEncryption: &clusters.DiskEncryption{
			Encrypted: "0",
		},
		Datastore: &clusters.Datastore{
			Version: "7.6.2",
			Type:    "elasticsearch",
		},
	}
	created, err := clusters.Create(client, opts)
	th.AssertNoErr(t, err)

	defer func() {
		err = clusters.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	}()

	got, err := clusters.Get(client, created.ID)
	th.AssertNoErr(t, err)

	log.Printf("Creating cluster, ID: %s", got.ID)
	th.AssertEquals(t, created.ID, got.ID)
	th.AssertEquals(t, created.Name, got.Name)

	th.CheckNoErr(t, clusters.WaitForClusterOperationSucces(client, created.ID, timeout))

	list, err := clusters.List(client)
	th.AssertNoErr(t, err)

	found := false
	for _, one := range list {
		if one.ID == created.ID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("cluster %s is not found in list", created.ID)
	}

	_, err = clusters.ExtendCluster(client, created.ID, []clusters.ClusterExtendSpecialOpts{
		{
			Type:     "ess",
			NodeSize: 0,
			DiskSize: 160,
		},
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterToExtend(client, created.ID, timeout))

	_, err = clusters.ExtendCluster(client, created.ID, clusters.ClusterExtendCommonOpts{
		ModifySize: 1,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterToExtend(client, created.ID, timeout))
}
