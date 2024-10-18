package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	ddmhelper "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/ddm/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v2/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQueryEngineInfoAndNodeClasses(t *testing.T) {
	// CREATE V2 ClIENT
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
	//CREATE V1 CLIENT
	ddmv1client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE V2 ClIENT
	ddmv2client, err := clients.NewDDMV2Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := ddmhelper.CreateDDMInstance(t, ddmv1client)
	t.Cleanup(func() {
		ddmhelper.DeleteDDMInstance(t, ddmv1client, ddmInstance.Id)
	})

	// SCALE OUT
	t.Logf("Scaling out DDM Instance %s", ddmInstance.Id)
	scaleOutOpts := instances.ScaleOutOpts{
		FlavorId:   "941b5a6d-3485-329e-902c-ffd49d352f16",
		NodeNumber: 1,
	}
	_, err = instances.ScaleOut(ddmv2client, ddmInstance.Id, scaleOutOpts)
	th.AssertNoErr(t, err)
	err = ddmhelper.WaitForInstanceInRunningState(ddmv1client, ddmInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("Scaled out DDM Instance %s to n+1 nodes", ddmInstance.Id)

	// SCALE IN
	t.Logf("Scaling in DDM Instance %s", ddmInstance.Id)
	scaleInOpts := instances.ScaleInOpts{
		NodeNumber: 1,
	}
	_, err = instances.ScaleIn(ddmv2client, ddmInstance.Id, scaleInOpts)
	th.AssertNoErr(t, err)
	err = ddmhelper.WaitForInstanceInRunningState(ddmv1client, ddmInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("Scaled in DDM Instance %s to n-1 nodes", ddmInstance.Id)

}

