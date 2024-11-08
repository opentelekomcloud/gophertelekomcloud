package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type Billing struct {
	Allocated       int    `json:"allocated"`
	ChargingMode    string `json:"charging_mode"`
	CloudType       string `json:"cloud_type"`
	ConsistentLevel string `json:"consistent_level"`
	ObjectType      string `json:"object_type"`
	OrderID         string `json:"order_id"`
	ProductID       string `json:"product_id"`
	ProtectType     string `json:"protect_type"`
	Size            int    `json:"size"`
	SpecCode        string `json:"spec_code"`
	Status          string `json:"status"`
	StorageUnit     string `json:"storage_unit"`
	Used            int    `json:"used"`
	FrozenScene     string `json:"frozen_scene"`
}

type ResourceResp struct {
	ExtraInfo     ResourceExtraInfo `json:"extra_info"`
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	ProtectStatus string            `json:"protect_status"`
	Size          int               `json:"size"`
	Type          string            `json:"type"`
	BackupSize    int               `json:"backup_size"`
	BackupCount   int               `json:"backup_count"`
}

type Vault struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Billing     Billing            `json:"billing"`
	Description string             `json:"description"`
	ProjectID   string             `json:"project_id"`
	ProviderID  string             `json:"provider_id"`
	Resources   []ResourceResp     `json:"resources"`
	Tags        []tags.ResourceTag `json:"tags"`
	AutoBind    bool               `json:"auto_bind"`
	BindRules   VaultBindRules     `json:"bind_rules"`
	UserID      string             `json:"user_id"`
	CreatedAt   string             `json:"created_at"`
	AutoExpand  bool               `json:"auto_expand"`
}
