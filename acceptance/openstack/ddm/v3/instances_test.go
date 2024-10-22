package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	ddmhelper "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/ddm/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDDMInstancesTestV3(t *testing.T) {
	// CREATE V1 CLIENT
	ddmV1Client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE V3 CLIENT
	ddmV3Client, err := clients.NewDDMV3Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := ddmhelper.CreateDDMInstance(t, ddmV1Client)
	t.Cleanup(func() {
		ddmhelper.DeleteDDMInstance(t, ddmV1Client, ddmInstance.Id)
	})

	// QUERY PARAMETERS
	t.Logf("Listing parameters for DDM instance  %s", ddmInstance.Id)
	_, err = instances.QueryParameters(ddmV3Client, ddmInstance.Id, instances.QueryParametersOpts{})
	th.AssertNoErr(t, err)

	// MODIFY PARAMETERS
	t.Logf("Modifying parameters for DDM instance  %s", ddmInstance.Id)
	modifyParametersOpts := instances.ModifyParametersOpts{
		Values: instances.Values{
			MaxConnections: "30000",
		},
	}
	_, err = instances.ModifyParameters(ddmV3Client, ddmInstance.Id, modifyParametersOpts)
	th.AssertNoErr(t, err)

	// CHANGING NODE CLASS
	t.Logf("Modifying node class for DDM instance  %s", ddmInstance.Id)
	changeNodeClassOpts := instances.ChangeNodeClassOpts{
		SpecCode: "ddm.8xlarge.2",
	}
	_, err = instances.ChangeNodeClass(ddmV3Client, ddmInstance.Id, changeNodeClassOpts)
	th.AssertNoErr(t, err)
	err = ddmhelper.WaitForInstanceInRunningState(ddmV1Client, ddmInstance.Id)
	th.AssertNoErr(t, err)
	t.Logf("Modified node class for DDM instance  %s", ddmInstance.Id)
}
