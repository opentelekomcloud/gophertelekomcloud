package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/extensions/quotasets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	uriQueryParms := map[string]string{}
	HandleSuccessfulRequest(t, "GET", "/os-quota-sets/"+FirstTenantID, getExpectedJSONBody, uriQueryParms)
	actual, err := quotasets.Get(client.ServiceClient(), FirstTenantID)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &getExpectedQuotaSet, actual)
}

func TestGetUsage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	uriQueryParms := map[string]string{"usage": "true"}
	HandleSuccessfulRequest(t, "GET", "/os-quota-sets/"+FirstTenantID, getUsageExpectedJSONBody, uriQueryParms)
	actual, err := quotasets.GetUsage(client.ServiceClient(), FirstTenantID)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &getUsageExpectedQuotaSet, actual)
}

func TestFullUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	uriQueryParms := map[string]string{}
	HandleSuccessfulRequest(t, "PUT", "/os-quota-sets/"+FirstTenantID, fullUpdateExpectedJSONBody, uriQueryParms)
	actual, err := quotasets.Update(client.ServiceClient(), FirstTenantID, fullUpdateOpts)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &fullUpdateExpectedQuotaSet, actual)
}

func TestPartialUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	uriQueryParms := map[string]string{}
	HandleSuccessfulRequest(t, "PUT", "/os-quota-sets/"+FirstTenantID, partialUpdateExpectedJSONBody, uriQueryParms)
	actual, err := quotasets.Update(client.ServiceClient(), FirstTenantID, partialUpdateOpts)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &partiualUpdateExpectedQuotaSet, actual)
}

func TestErrorInToBlockStorageQuotaUpdateMap(t *testing.T) {
	opts := quotasets.UpdateOpts{}
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSuccessfulRequest(t, "PUT", "/os-quota-sets/"+FirstTenantID, "", nil)
	_, err := quotasets.Update(client.ServiceClient(), FirstTenantID, opts)
	if err == nil {
		t.Fatal("Error handling failed")
	}
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := quotasets.Delete(client.ServiceClient(), FirstTenantID)
	th.AssertNoErr(t, err)
}
