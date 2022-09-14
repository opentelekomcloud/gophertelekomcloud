package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ResourceBackupCapOpts contains the options for querying whether resources can be backed up. This object is
// passed to backup.QueryResourceBackupCapability().
type ResourceBackupCapOpts struct {
	CheckProtectable []ResourceCapQueryParams `json:"check_protectable" required:"true"`
}

type ResourceCapQueryParams struct {
	ResourceId   string `json:"resource_id" required:"true"`
	ResourceType string `json:"resource_type" required:"true"`
}

// QueryResourceBackupCapability will query whether resources can be backed up based on the values in ResourceBackupCapOpts. To extract
// the ResourceCap object from the response, call the ExtractQueryResponse method on the QueryResult.
func QueryResourceBackupCapability(client *golangsdk.ServiceClient, opts ResourceBackupCapOpts) ([]ResourceCapability, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("providers", providerID, "resources", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []ResourceCapability
	err = extract.IntoSlicePtr(raw.Body, &res, "protectable")
	return res, err
}

type ResourceCapability struct {
	Result       bool   `json:"result"`
	ResourceType string `json:"resource_type"`
	ErrorCode    string `json:"error_code"`
	ErrorMsg     string `json:"error_msg"`
	ResourceId   string `json:"resource_id"`
}
