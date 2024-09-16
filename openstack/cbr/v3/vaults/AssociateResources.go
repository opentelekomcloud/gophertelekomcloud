package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

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

type AssociateResourcesOpts struct {
	Resources []ResourceCreate `json:"resources"`
}

func AssociateResources(client *golangsdk.ServiceClient, vaultID string, opts AssociateResourcesOpts) ([]string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("vaults", vaultID, "addresources"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []string
	return res, extract.IntoSlicePtr(raw.Body, &res, "add_resource_ids")
}
