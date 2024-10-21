package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/query"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSpecificResourceTypeList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListSpecificOpts{}
	resources, err := query.ListSpecificType(
		client,
		client.DomainID,
		"ecs",
		"cloudservers",
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.AssertLengthGreaterThan(t, resources, 1)
}

func TestRecorderResourceList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListRecordedOpts{}
	resources, err := query.ListRecordedResources(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(resources))
}

func TestServicesList(t *testing.T) {
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListServicesOpts{}
	resources, err := query.ListServices(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.AssertLengthGreaterThan(t, resources, 1)
}

func TestGetSpecificByIdResource(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListSpecificOpts{}
	resources, err := query.ListSpecificType(
		client,
		client.DomainID,
		"ecs",
		"cloudservers",
		listOpts,
	)
	th.AssertNoErr(t, err)
	if len(resources) != 0 {
		res, err := query.GetResource(
			client,
			client.DomainID,
			"ecs",
			"cloudservers",
			resources[0].ID,
		)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, resources[0].ID, res.ID)
	}
}

func TestRecordedResourcesTagsList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListRecordedTagsOpts{}
	resources, err := query.ListRecordedResourcesTags(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(resources))
}

func TestRecordedResourcesSummaryList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListRecordedSummaryOpts{}
	resources, err := query.ListRecordedResourcesSummary(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(resources))
}

func TestAllResourcesList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListAllOpts{}
	resources, err := query.ListAllResources(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.AssertLengthGreaterThan(t, resources, 1)
}

func TestAnyResourceById(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListAllOpts{}
	resources, err := query.ListAllResources(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	if len(resources) != 0 {
		res, err := query.GetAnyResource(
			client,
			client.DomainID,
			resources[0].ID,
		)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, resources[0].ID, res.ID)
	}
}

func TestGetTagsFromAnyResource(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListAllTagsOpts{}
	tags, err := query.ListAllResourcesTags(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.AssertLengthGreaterThan(t, tags, 1)
}

func TestGetCountResources(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.CountOpts{}
	count, err := query.GetCount(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, count)
}

func TestResourcesSummaryList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	listOpts := query.ListSummaryOpts{}
	resources, err := query.ListResourcesSummary(
		client,
		client.DomainID,
		listOpts,
	)
	th.AssertNoErr(t, err)
	tools.AssertLengthGreaterThan(t, resources, 1)
}
