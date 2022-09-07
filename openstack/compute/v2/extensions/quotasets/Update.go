package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update updates the quotas for the given tenantID and returns the new QuotaSet.
func Update(client *golangsdk.ServiceClient, tenantID string, opts UpdateOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToComputeQuotaUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(updateURL(client, tenantID), reqBody, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return
}
