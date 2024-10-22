package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	ddmhelper "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/ddm/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDDMInstancesTestV3(t *testing.T) {
	//CREATE V1 CLIENT
	ddmv1client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE V3 ClIENT
	ddmv3client, err := clients.NewDDMV3Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := ddmhelper.CreateDDMInstance(t, ddmv1client)
	t.Cleanup(func() {
		ddmhelper.DeleteDDMInstance(t, ddmv1client, ddmInstance.Id)
	})

	// QUERY PARAMETERS
	t.Logf("Listing parameters for DDM instance  %s", ddmInstance.Id)
	_, err = instances.QueryParameters(ddmv3client, ddmInstance.Id, instances.QueryParametersOpts{})
	th.AssertNoErr(t, err)

	// MODIFY PARAMETERS
	t.Logf("Modifying parameters for DDM instance  %s", ddmInstance.Id)
	modifyParametersOpts := instances.ModifyParametersOpts{
		Values: instances.Values{
			MaxConnections: "30000",
		},
	}
	_, err = instances.ModifyParameters(ddmv3client, ddmInstance.Id, modifyParametersOpts)
	th.AssertNoErr(t, err)

	// CHANGING NODE CLASS
	t.Logf("Modifying node class for DDM instance  %s", ddmInstance.Id)
	changeNodeClassOpts := instances.ChangeNodeClassOpts{
		SpecCode: "ddm.8xlarge.2",
	}
	_, err = instances.ChangeNodeClass(ddmv3client, ddmInstance.Id, changeNodeClassOpts)
	th.AssertNoErr(t, err)
	err = ddmhelper.WaitForInstanceInRunningState(ddmv1client, ddmInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("Modified node class for DDM instance  %s", ddmInstance.Id)
}
