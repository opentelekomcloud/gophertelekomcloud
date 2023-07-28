package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type BillingCreate struct {
	// Cloud platform. Enumeration values:
	//    public
	//    hybrid
	CloudType string `json:"cloud_type,omitempty"`
	// Backup specifications. The default value is `crash_consistent`
	ConsistentLevel string `json:"consistent_level"`
	// Object type
	ObjectType string `json:"object_type"`
	// Operation type. Enumeration values:
	//    backup
	//    replication
	ProtectType string `json:"protect_type"`
	// Capacity, in GB. Minimum: `1`. Maximum: `10485760`
	Size int `json:"size"`
	// Billing mode. Possible values are `post_paid` (pay-per-use) or `pre_paid` (yearly/monthly packages).
	// The value defaults to post_paid.
	ChargingMode string `json:"charging_mode,omitempty"`
	// Package type. This parameter is mandatory if charging_mode is set to pre_paid.
	// Possible values are `year` (yearly) or `month`(monthly).
	PeriodType string `json:"period_type,omitempty"`
	// Required duration for the package. This parameter is mandatory if charging_mode is set to `pre_paid`.
	PeriodNum int `json:"period_num,omitempty"`
	// Whether to automatically renew the subscription after expiration. By default, it is not renewed.
	IsAutoRenew bool `json:"is_auto_renew,omitempty"`
	// Whether the fee is automatically deducted from the customer's account balance after an order is submitted.
	// The non-automatic payment mode is used by default.
	IsAutoPay bool `json:"is_auto_pay,omitempty"`
	// Redirection URL
	ConsoleURL string `json:"console_url,omitempty"`
	// Extended information for creating a vault
	ExtraInfo *BillingCreateExtraInfo `json:"extra_info,omitempty"`
}

type BillingCreateExtraInfo struct {
	// ID of the application for creating vaults in combination.
	// This parameter is mandatory when creating vaults in combination.
	CombinedOrderID string `json:"combined_order_id,omitempty"`
	// Number of items in the application for creating vaults in the combination mode.
	// This parameter is mandatory when creating vaults in the combination mode.
	CombinedOrderECSNum int `json:"combined_order_ecs_num,omitempty"`
}

type CreateOpts struct {
	// Backup policy ID. If the value of this parameter is missing,
	// automatic backup is not performed.
	BackupPolicyID string `json:"backup_policy_id,omitempty"`
	// Parameter information for billing creation
	Billing *BillingCreate `json:"billing"`
	// User-defined vault description
	Description string `json:"description,omitempty"`
	// Vault name
	Name string `json:"name"`
	// Associated resources. Set this parameter to [] if no resources are associated when creating a vault.
	Resources []ResourceCreate `json:"resources"`
	// Tags - Tag list.
	// This list cannot be an empty list.
	// The list can contain up to 10 keys.
	// Keys in this list must be unique.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Enterprise project ID. The default value is 0.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
	// Whether automatic association is supported
	AutoBind bool `json:"auto_bind,omitempty"`
	// Rules for automatic association
	BindRules *VaultBindRules `json:"bind_rules,omitempty"`
	// Whether to automatically expand the vault capacity.
	// Only pay-per-use vaults support this function.
	AutoExpand bool `json:"auto_expand,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Vault, error) {
	reqBody, err := build.RequestBodyMap(opts, "vault")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("vaults"), reqBody, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Vault
	return &res, extract.IntoStructPtr(raw.Body, &res, "vault")
}
