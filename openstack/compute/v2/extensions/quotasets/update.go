package quotasets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts - options for Updating the quotas of a Tenant.
// All int-values are pointers, so they can be nil if they are not needed.
// You can use golangsdk.IntToPointer() for convenience
type UpdateOpts struct {
	// FixedIPs is number of fixed ips alloted this quota_set.
	FixedIPs *int `json:"fixed_ips,omitempty"`
	// FloatingIPs is number of floating ips alloted this quota_set.
	FloatingIPs *int `json:"floating_ips,omitempty"`
	// InjectedFileContentBytes is content bytes allowed for each injected file.
	InjectedFileContentBytes *int `json:"injected_file_content_bytes,omitempty"`
	// InjectedFilePathBytes is allowed bytes for each injected file path.
	InjectedFilePathBytes *int `json:"injected_file_path_bytes,omitempty"`
	// InjectedFiles is injected files allowed for each project.
	InjectedFiles *int `json:"injected_files,omitempty"`
	// KeyPairs is number of ssh keypairs.
	KeyPairs *int `json:"key_pairs,omitempty"`
	// MetadataItems is number of metadata items allowed for each instance.
	MetadataItems *int `json:"metadata_items,omitempty"`
	// RAM is megabytes allowed for each instance.
	RAM *int `json:"ram,omitempty"`
	// SecurityGroupRules is rules allowed for each security group.
	SecurityGroupRules *int `json:"security_group_rules,omitempty"`
	// SecurityGroups security groups allowed for each project.
	SecurityGroups *int `json:"security_groups,omitempty"`
	// Cores is number of instance cores allowed for each project.
	Cores *int `json:"cores,omitempty"`
	// Instances is number of instances allowed for each project.
	Instances *int `json:"instances,omitempty"`
	// Number of ServerGroups allowed for the project.
	ServerGroups *int `json:"server_groups,omitempty"`
	// Max number of Members for each ServerGroup.
	ServerGroupMembers *int `json:"server_group_members,omitempty"`
	// Force will update the quotaset even if the quota has already been used
	// and the reserved quota exceeds the new quota.
	Force bool `json:"force,omitempty"`
}

// Update updates the quotas for the given tenantID and returns the new QuotaSet.
func Update(client *golangsdk.ServiceClient, tenantID string, opts UpdateOpts) (*QuotaSet, error) {
	reqBody, err := build.RequestBody(opts, "quota_set")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("os-quota-sets", tenantID), reqBody, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}
