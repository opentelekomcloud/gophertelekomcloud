package vaults

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOptsBuilder interface {
	ToVaultCreateMap() (map[string]interface{}, error)
}

type BillingCreateExtraInfo struct {
	// ID of the application for creating vaults in combination.
	// This parameter is mandatory when creating vaults in combination.
	CombinedOrderID string `json:"combined_order_id,omitempty"`
	// Number of items in the application for creating vaults in the combination mode.
	// This parameter is mandatory when creating vaults in the combination mode.
	CombinedOrderECSNum int `json:"combined_order_ecs_num,omitempty"`
}

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

type ResourceExtraInfoIncludeVolumes struct {
	// EVS disk ID. Only UUID is supported.
	ID string `json:"id"`
	// OS type
	OSVersion string `json:"os_version,omitempty"`
}

type ResourceExtraInfo struct {
	// ID of the disk that is excluded from the backup.
	// This parameter is used only when there are VM disk backups.
	ExcludeVolumes []string `json:"exclude_volumes,omitempty"`
	// Disk to be backed up
	IncludeVolumes []ResourceExtraInfoIncludeVolumes `json:"include_volumes,omitempty"`
}

type ResourceCreate struct {
	// ID of the resource to be backed up
	ID string `json:"id"`
	// Type of the resource to be backed up.
	// Possible values are `OS::Nova::Server` and `OS::Cinder::Volume`
	Type string `json:"type"`
	// Resource name
	Name string `json:"name,omitempty"`
	// Extra information of the resource
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`
}

type VaultBindRules struct {
	// Filters automatically associated resources by tag.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
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

func (opts CreateOpts) ToVaultCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "vault")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToVaultCreateMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create vault create map: %s", err)
		return
	}
	_, err = client.Post(rootURL(client), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	r.Err = err
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(vaultURL(client, id), nil)
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(vaultURL(client, id), &r.Body, nil)
	return
}

type UpdateResult struct {
	vaultResult
}

type UpdateOptsBuilder interface {
	ToVaultUpdateMap() (map[string]interface{}, error)
}

type BillingUpdate struct {
	Size int `json:"size,omitempty"`
}

type UpdateOpts struct {
	Billing    *BillingUpdate  `json:"billing,omitempty"`
	Name       string          `json:"name,omitempty"`
	AutoBind   *bool           `json:"auto_bind,omitempty"`
	BindRules  *VaultBindRules `json:"bind_rules,omitempty"`
	AutoExpand *bool           `json:"auto_expand,omitempty"`
}

func (opts UpdateOpts) ToVaultUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "vault")
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToVaultUpdateMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create vault update map: %s", err)
		return
	}
	_, r.Err = client.Put(vaultURL(client, id), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type BindPolicyOptsBuilder interface {
	ToBindPolicyMap() (map[string]interface{}, error)
}

type BindPolicyOpts struct {
	PolicyID string `json:"policy_id"`
}

func (opts BindPolicyOpts) ToBindPolicyMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func BindPolicy(client *golangsdk.ServiceClient, vaultID string, opts BindPolicyOptsBuilder) (r BindPolicyResult) {
	reqBody, err := opts.ToBindPolicyMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create bind policy map: %s", err)
		return
	}
	_, r.Err = client.Post(bindPolicyURL(client, vaultID), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func UnbindPolicy(client *golangsdk.ServiceClient, vaultID string, opts BindPolicyOptsBuilder) (r UnbindPolicyResult) {
	reqBody, err := opts.ToBindPolicyMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create bind policy map: %s", err)
		return
	}
	_, r.Err = client.Post(unbindPolicyURL(client, vaultID), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type AssociateResourcesOptsBuilder interface {
	ToAssociateResourcesMap() (map[string]interface{}, error)
}
type AssociateResourcesOpts struct {
	Resources []ResourceCreate `json:"resources"`
}

func (opts AssociateResourcesOpts) ToAssociateResourcesMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func AssociateResources(client *golangsdk.ServiceClient, vaultID string, opts AssociateResourcesOptsBuilder) (r AssociateResourcesResult) {
	reqBody, err := opts.ToAssociateResourcesMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create associate resource map: %s", err)
		return
	}
	_, r.Err = client.Post(addResourcesURL(client, vaultID), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type DissociateResourcesOptsBuilder interface {
	ToDissociateResourcesMap() (map[string]interface{}, error)
}
type DissociateResourcesOpts struct {
	ResourceIDs []string `json:"resource_ids"`
}

func (opts DissociateResourcesOpts) ToDissociateResourcesMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func DissociateResources(client *golangsdk.ServiceClient, vaultID string, opts DissociateResourcesOptsBuilder) (r DissociateResourcesResult) {
	reqBody, err := opts.ToDissociateResourcesMap()
	if err != nil {
		r.Err = fmt.Errorf("failed to create dissociate resource map: %s", err)
		return
	}
	_, r.Err = client.Post(removeResourcesURL(client, vaultID), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
