package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/aggregates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := aggregates.List(client.ServiceClient())
	th.AssertNoErr(t, err)

	if len(actual) != 2 {
		t.Fatalf("Expected 2 aggregates, got %d", len(actual))
	}
	th.CheckDeepEquals(t, FirstFakeAggregate, actual[0])
	th.CheckDeepEquals(t, SecondFakeAggregate, actual[1])
}

func TestCreateAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	expected := CreatedAggregate

	opts := aggregates.CreateOpts{
		Name:             "name",
		AvailabilityZone: "london",
	}

	actual, err := aggregates.Create(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestDeleteAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := aggregates.Delete(client.ServiceClient(), AggregateIDtoDelete)
	th.AssertNoErr(t, err)
}

func TestGetAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	expected := SecondFakeAggregate

	actual, err := aggregates.Get(client.ServiceClient(), AggregateIDtoGet)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestUpdateAggregate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	expected := UpdatedAggregate

	opts := aggregates.UpdateOpts{
		Name:             "test-aggregates2",
		AvailabilityZone: "nova2",
	}

	actual, err := aggregates.Update(client.ServiceClient(), expected.ID, opts)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestAddHostAggregate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddHostSuccessfully(t)

	expected := AggregateWithAddedHost

	opts := aggregates.AddHostOpts{
		Host: "cmp1",
	}

	actual, err := aggregates.AddHost(client.ServiceClient(), expected.ID, opts)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestRemoveHostAggregate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRemoveHostSuccessfully(t)

	expected := AggregateWithRemovedHost

	opts := aggregates.RemoveHostOpts{
		Host: "cmp1",
	}

	actual, err := aggregates.RemoveHost(client.ServiceClient(), expected.ID, opts)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestSetMetadataAggregate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSetMetadataSuccessfully(t)

	expected := AggregateWithUpdatedMetadata

	opts := aggregates.SetMetadataOpts{
		Metadata: map[string]interface{}{"key": "value"},
	}

	actual, err := aggregates.SetMetadata(client.ServiceClient(), expected.ID, opts)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}
