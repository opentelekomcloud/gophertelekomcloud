package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	ddmhelper "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/ddm/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v2/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQueryEngineInfoAndNodeClasses(t *testing.T) {
	// CREATE V2 CLIENT
	client, err := clients.NewDDMV2Client()
	th.AssertNoErr(t, err)

	engineGroups, err := instances.QueryEngineInfo(client, instances.QueryEngineOpts{})
	th.AssertNoErr(t, err)

	queryNodeClassesOpts := instances.QueryNodeClassesOpts{
		EngineId: engineGroups[0].ID,
	}
	_, err = instances.QueryNodeClasses(client, queryNodeClassesOpts)
	th.AssertNoErr(t, err)
}

func TestDDMInstancesV2Scaling(t *testing.T) {
	// CREATE V1 CLIENT
	ddmV1Client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE V2 CLIENT
	ddmV2Client, err := clients.NewDDMV2Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := ddmhelper.CreateDDMInstance(t, ddmV1Client)
	t.Cleanup(func() {
		ddmhelper.DeleteDDMInstance(t, ddmV1Client, ddmInstance.Id)
	})

	// SCALE OUT
	engineId := "367b68a3-b48b-3d8a-b3a1-4c463a75a4b4"
	clientV2, err := clients.NewDDMV2Client()
	th.AssertNoErr(t, err)
	engines, err := instances.QueryEngineInfo(clientV2, instances.QueryEngineOpts{})
	th.AssertNoErr(t, err)
	if len(engines) != 0 {
		engineId = engines[0].ID
	}

	flavorId := "941b5a6d-3485-329e-902c-ffd49d352f16"
	classes, err := instances.QueryNodeClasses(clientV2, instances.QueryNodeClassesOpts{
		EngineId: engineId,
	})
	th.AssertNoErr(t, err)
	if len(classes.ComputeFlavorGroups) != 0 {
		flavorId = classes.ComputeFlavorGroups[0].ComputeFlavors[0].ID
	}
	t.Logf("Scaling out DDM Instance %s", ddmInstance.Id)
	scaleOutOpts := instances.ScaleOutOpts{
		FlavorId:   flavorId,
		NodeNumber: 1,
	}
	_, err = instances.ScaleOut(ddmV2Client, ddmInstance.Id, scaleOutOpts)
	th.AssertNoErr(t, err)
	err = ddmhelper.WaitForInstanceInRunningState(ddmV1Client, ddmInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("Scaled out DDM Instance %s to n+1 nodes", ddmInstance.Id)

	// SCALE IN
	t.Logf("Scaling in DDM Instance %s", ddmInstance.Id)
	scaleInOpts := instances.ScaleInOpts{
		NodeNumber: 1,
	}
	_, err = instances.ScaleIn(ddmV2Client, ddmInstance.Id, scaleInOpts)
	th.AssertNoErr(t, err)
	err = ddmhelper.WaitForInstanceInRunningState(ddmV1Client, ddmInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("Scaled in DDM Instance %s to n-1 nodes", ddmInstance.Id)

}

// MANUAL TEST
func TestDDMInstancesV2ModifyReadPolicy(t *testing.T) {
	t.Skip("Impossible to run with CI")

	// CREATE V2 ClIENT
	ddmv2client, err := clients.NewDDMV2Client()
	th.AssertNoErr(t, err)

	ddmInstanceId := "6f24dd2d4f83469b870fcf9fbc7a5180in09"
	modifyDbReadPolicyOpts := instances.ModifyDbReadPolicyOpts{
		ReadWeight: map[string]int{
			"55d93e249b77461b81f990fa805db3f3in01": 60,
		},
	}

	modifyDbReadPolicy, err := instances.ModifyDbReadPolicy(ddmv2client, ddmInstanceId, modifyDbReadPolicyOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, modifyDbReadPolicy.Success, true)

}
