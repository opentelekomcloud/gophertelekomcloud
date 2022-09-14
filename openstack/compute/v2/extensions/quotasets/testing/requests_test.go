package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/quotasets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(th.TeardownHTTP)
	HandleGetSuccessfully(t)
	actual, err := quotasets.Get(client.ServiceClient(), FirstTenantID)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstQuotaSet, actual)
}

func TestGetDetail(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(th.TeardownHTTP)
	HandleGetDetailSuccessfully(t)
	actual, err := quotasets.GetDetail(client.ServiceClient(), FirstTenantID)
	th.CheckDeepEquals(t, &FirstQuotaDetailsSet, actual)
	th.AssertNoErr(t, err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(th.TeardownHTTP)
	HandlePutSuccessfully(t)
	actual, err := quotasets.Update(client.ServiceClient(), FirstTenantID, UpdatedQuotaSet)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstQuotaSet, actual)
}

func TestPartialUpdate(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(th.TeardownHTTP)
	HandlePartialPutSuccessfully(t)
	opts := quotasets.UpdateOpts{Cores: golangsdk.IntToPointer(200), Force: true}
	actual, err := quotasets.Update(client.ServiceClient(), FirstTenantID, opts)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstQuotaSet, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(th.TeardownHTTP)
	HandleDeleteSuccessfully(t)
	err := quotasets.Delete(client.ServiceClient(), FirstTenantID)
	th.AssertNoErr(t, err)
}
