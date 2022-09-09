package quotasets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetDetail returns detailed public data about a previously created QuotaSet.
func GetDetail(client *golangsdk.ServiceClient, tenantID string) (*QuotaDetailSet, error) {
	raw, err := client.Get(client.ServiceURL("os-quota-sets", tenantID, "detail"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuotaDetailSet
	err = extract.IntoStructPtr(raw.Body, &res, "quota_set")
	return &res, err
}

// QuotaDetailSet represents details of both operational limits of compute
// resources and the current usage of those resources.
type QuotaDetailSet struct {
	// ID is the tenant ID associated with this QuotaDetailSet.
	ID string `json:"id"`
	// FixedIPs is number of fixed ips alloted this QuotaDetailSet.
	FixedIPs QuotaDetail `json:"fixed_ips"`
	// FloatingIPs is number of floating ips alloted this QuotaDetailSet.
	FloatingIPs QuotaDetail `json:"floating_ips"`
	// InjectedFileContentBytes is the allowed bytes for each injected file.
	InjectedFileContentBytes QuotaDetail `json:"injected_file_content_bytes"`
	// InjectedFilePathBytes is allowed bytes for each injected file path.
	InjectedFilePathBytes QuotaDetail `json:"injected_file_path_bytes"`
	// InjectedFiles is the number of injected files allowed for each project.
	InjectedFiles QuotaDetail `json:"injected_files"`
	// KeyPairs is number of ssh keypairs.
	KeyPairs QuotaDetail `json:"key_pairs"`
	// MetadataItems is number of metadata items allowed for each instance.
	MetadataItems QuotaDetail `json:"metadata_items"`
	// RAM is megabytes allowed for each instance.
	RAM QuotaDetail `json:"ram"`
	// SecurityGroupRules is number of security group rules allowed for each
	// security group.
	SecurityGroupRules QuotaDetail `json:"security_group_rules"`
	// SecurityGroups is the number of security groups allowed for each project.
	SecurityGroups QuotaDetail `json:"security_groups"`
	// Cores is number of instance cores allowed for each project.
	Cores QuotaDetail `json:"cores"`
	// Instances is number of instances allowed for each project.
	Instances QuotaDetail `json:"instances"`
	// ServerGroups is the number of ServerGroups allowed for the project.
	ServerGroups QuotaDetail `json:"server_groups"`
	// ServerGroupMembers is the number of members for each ServerGroup.
	ServerGroupMembers QuotaDetail `json:"server_group_members"`
}

// QuotaDetail is a set of details about a single operational limit that allows
// for control of compute usage.
type QuotaDetail struct {
	// InUse is the current number of provisioned/allocated resources of the given type.
	InUse int `json:"in_use"`
	// Reserved is a transitional state when a claim against quota has been made
	// but the resource is not yet fully online.
	Reserved int `json:"reserved"`
	// Limit is the maximum number of a given resource that can be
	// allocated/provisioned.  This is what "quota" usually refers to.
	Limit int `json:"limit"`
}
