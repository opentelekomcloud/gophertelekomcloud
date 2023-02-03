package quotasets

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// QuotaSet is a set of operational limits that allow for control of compute usage.
type QuotaSet struct {
	// ID is tenant associated with this QuotaSet.
	ID string `json:"id"`
	// FixedIPs is number of fixed ips alloted this QuotaSet.
	FixedIPs int `json:"fixed_ips"`
	// FloatingIPs is number of floating ips alloted this QuotaSet.
	FloatingIPs int `json:"floating_ips"`
	// InjectedFileContentBytes is the allowed bytes for each injected file.
	InjectedFileContentBytes int `json:"injected_file_content_bytes"`
	// InjectedFilePathBytes is allowed bytes for each injected file path.
	InjectedFilePathBytes int `json:"injected_file_path_bytes"`
	// InjectedFiles is the number of injected files allowed for each project.
	InjectedFiles int `json:"injected_files"`
	// KeyPairs is number of ssh keypairs.
	KeyPairs int `json:"key_pairs"`
	// MetadataItems is number of metadata items allowed for each instance.
	MetadataItems int `json:"metadata_items"`
	// RAM is megabytes allowed for each instance.
	RAM int `json:"ram"`
	// SecurityGroupRules is number of security group rules allowed for each security group.
	SecurityGroupRules int `json:"security_group_rules"`
	// SecurityGroups is the number of security groups allowed for each project.
	SecurityGroups int `json:"security_groups"`
	// Cores is number of instance cores allowed for each project.
	Cores int `json:"cores"`
	// Instances is number of instances allowed for each project.
	Instances int `json:"instances"`
	// ServerGroups is the number of ServerGroups allowed for the project.
	ServerGroups int `json:"server_groups"`
	// ServerGroupMembers is the number of members for each ServerGroup.
	ServerGroupMembers int `json:"server_group_members"`
}

func extra(err error, raw *http.Response) (*QuotaSet, error) {
	if err != nil {
		return nil, err
	}

	var res QuotaSet
	err = extract.IntoStructPtr(raw.Body, &res, "quota_set")
	return &res, err
}
