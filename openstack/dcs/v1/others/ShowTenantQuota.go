package others

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowTenantQuota(client *golangsdk.ServiceClient) (*ShowQuotaOfTenantResponse, error) {
	raw, err := client.Get(client.ServiceURL("quota"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowQuotaOfTenantResponse
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}

type ShowQuotaOfTenantResponse struct {
	// List of quotas.
	Resources []Resources `json:"resources"`
	// Information about a resource tenant
	ResourceUser ResourceUser `json:"resource_user"`
}

type Resources struct {
	// Resource unit.
	// When type is set to instance, no value is returned.
	// When type is set to ram, GB is returned.
	Unit string `json:"unit"`
	// Indicates the minimum limit of instance quota when type is set to instance.
	// Indicates the minimum limit of memory quota when type is set to ram.
	Min int32 `json:"min"`
	// Indicates the maximum limit of instance quota when type is set to instance.
	// Indicates the maximum limit of memory quota when type is set to ram.
	Max int32 `json:"max"`
	// Maximum number of instances that can be created and maximum allowed total memory.
	Quota int32 `json:"quota"`
	// Number of created instances and used memory.
	Used int32 `json:"used"`
	// Values:
	// instances: indicates the instance quota.
	// ram: indicates the memory quota.
	Type string `json:"type"`
}

type ResourceUser struct {
	// Resource tenant ID
	TenantId string `json:"tenant_id"`
	// Resource tenant name
	TenantName string `json:"tenant_name"`
}
