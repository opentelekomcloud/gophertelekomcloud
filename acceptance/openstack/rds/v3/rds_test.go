package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRdsList(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	listOpts := instances.ListRdsInstanceOpts{}
	allRdsPages, err := instances.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	rdsInstances, err := instances.ExtractRdsInstances(allRdsPages)
	th.AssertNoErr(t, err)

	for _, rds := range rdsInstances.Instances {
		tools.PrintResource(t, rds)
	}
}

func TestRdsLifecycle(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	// Create RDSv3 instance
	rds := createRDS(t, client, cc.RegionName)
	defer deleteRDS(t, client, rds.Id)
	th.AssertEquals(t, rds.Volume.Size, 100)

	tagList := []tags.ResourceTag{
		{
			Key:   "muh",
			Value: "value-create",
		},
		{
			Key:   "kuh",
			Value: "value-create",
		},
	}
	err = tags.Create(client, "instances", rds.Id, tagList).ExtractErr()
	th.AssertNoErr(t, err)

	err = updateRDS(t, client, rds.Id)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, rds)

	listOpts := instances.ListRdsInstanceOpts{
		Id: rds.Id,
	}
	allPages, err := instances.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)
	newRds, err := instances.ExtractRdsInstances(allPages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(newRds.Instances), 1)
	th.AssertEquals(t, newRds.Instances[0].Volume.Size, 200)
	th.AssertEquals(t, len(newRds.Instances[0].Tags), 2)
}
