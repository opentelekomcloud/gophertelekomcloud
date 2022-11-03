package quotasets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetDefaults(client *golangsdk.ServiceClient, projectID string) (*QuotaSet, error) {
	raw, err := client.Get(client.ServiceURL("os-quota-sets", projectID, "defaults"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuotaSet
	err = extract.IntoStructPtr(raw.Body, &res, "quota_set")
	return &res, err
}
