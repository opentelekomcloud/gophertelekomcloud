package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1/instances"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDDMQueryInstances(t *testing.T) {
	client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)

	queryOpts := instances.QueryInstancesOpts{}
	_, err = instances.QueryInstances(client, queryOpts)
	th.AssertNoErr(t, err)
}

func TestDDMInstancesLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	if subnetID == "" || vpcID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but are required for DDM instances test")
	}

	// CREATE CLIENT
	client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := CreateDDMInstance(t, client)
	t.Logf("Creating Second Security Group")
	newSecGroupId := openstack.CreateSecurityGroup(t)
	t.Cleanup(func() {
		t.Logf("Deleting DDM Instance: %s", ddmInstance.Id)
		DeleteDDMInstance(t, client, ddmInstance.Id)
		openstack.DeleteSecurityGroup(t, newSecGroupId)
	})
	// RENAME DDM INSTANCE
	newInstanceName := tools.RandomString("ddm-instance-renamed-", 3)
	t.Logf("Renaming DDM Instance '%s' to '%s'", ddmInstance.Id, newInstanceName)
	renamedInstanceName, err := instances.Rename(client, ddmInstance.Id, newInstanceName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, newInstanceName, *renamedInstanceName)
	t.Logf("Renamed DDM Instance '%s' to '%s'", ddmInstance.Id, newInstanceName)

	// MODIFY SECURITY GROUP
	modifySecGroupOpts := instances.ModifySecurityGroupOpts{
		SecurityGroupId: newSecGroupId,
	}
	modifiedSecGroupID, err := instances.ModifySecurityGroup(client, ddmInstance.Id, modifySecGroupOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, newSecGroupId, *modifiedSecGroupID)

	// QUERY NODES
	t.Logf("Quering DDM Nodes")
	queryNodeOpts := instances.QueryNodesOpts{}
	nodeList, err := instances.QueryNodes(client, ddmInstance.Id, queryNodeOpts)
	th.AssertEquals(t, len(nodeList), 2)
	th.AssertNoErr(t, err)

	// QUERY NODE DETAILS
	t.Logf("Quering DDM Node details: %s", nodeList[0].NodeID)
	nodeDetails, err := instances.QueryNodeDetails(client, ddmInstance.Id, nodeList[0].NodeID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, nodeDetails.NodeID, nodeList[0].NodeID)
}
