package v1

import (
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/kibana"
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

func TestClusterPublicAccess(t *testing.T) {
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
		HttpsEnabled:     "true",
		AuthorityEnabled: true,
		AdminPassword:    "Test123!@#",
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

	managePublicOpts := clusters.ManagePublicAccessOpts{
		ClusterId: got.ID,
		Size:      5,
	}

	res, err := clusters.EnablePublicAccess(client, managePublicOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, res, "bindZone")

	err = clusters.WaitForCluster(client, managePublicOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)

	managePublicOpts.Size = 10
	err = clusters.UpdatePublicAccess(client, managePublicOpts)
	th.AssertNoErr(t, err)

	err = clusters.WaitForCluster(client, managePublicOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)

	res, err = clusters.DisablePublicAccess(client, managePublicOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, res, "unbindZone")

	err = clusters.WaitForCluster(client, managePublicOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)
}

func TestClusterKibanaPublicAccess(t *testing.T) {
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
		HttpsEnabled:     "true",
		AuthorityEnabled: true,
		AdminPassword:    "Test123!@#",
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

	clusterId := got.ID

	manageKibanaOpts := kibana.ManageOpts{
		ClusterId:       clusterId,
		EipSize:         5,
		EnableWhiteList: true,
		Whitelist:       "192.168.1.1",
	}

	err = kibana.Enable(client, manageKibanaOpts)
	th.AssertNoErr(t, err)

	err = clusters.WaitForCluster(client, manageKibanaOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)

	err = kibana.Update(client, clusterId, 10)
	th.AssertNoErr(t, err)

	err = clusters.WaitForCluster(client, manageKibanaOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)

	err = kibana.Disable(client, manageKibanaOpts)
	th.AssertNoErr(t, err)

	err = clusters.WaitForCluster(client, manageKibanaOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)

	err = kibana.UpdateAccess(client, clusterId, manageKibanaOpts.Whitelist)
	th.AssertNoErr(t, err)

	err = clusters.WaitForCluster(client, manageKibanaOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)

	err = kibana.DisableAccess(client, clusterId)
	th.AssertNoErr(t, err)

	err = clusters.WaitForCluster(client, manageKibanaOpts.ClusterId, timeout)
	th.AssertNoErr(t, err)
}
