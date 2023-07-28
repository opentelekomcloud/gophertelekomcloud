package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CountOpts struct {
	// Query based on field status is supported.
	// Value range: waiting_protect, protecting, available, waiting_restore, restoring, error, waiting_delete, deleting, and deleted
	Status string `q:"status"`
	// Whether to query the backup of all tenants. Only administrators can query the backup of all tenants.
	AllTenants string `q:"all_tenants"`
	// Supports query by backup name.
	Name string `q:"name"`
	// AZ-based filtering is supported.
	Az string `q:"az"`
	// Filtering based on the backup object ID is supported.
	ResourceId string `q:"resource_id"`
	// Filtering based on the backup object name is supported.
	ResourceName string `q:"resource_name"`
	// Filtering based on the backup time is supported. This is the backup start time. For example, 2017-04-15T04:25:38
	StartTime string `q:"start_time"`
	// Filtering based on the backup time is supported. This is the backup end time. For example, 2017-04-15T04:25:38
	EndTime string `q:"end_time"`
	// Supports filtering by backup image type. This parameter can be used only when images are created using backups.
	// The image type can be obtained from Image Management Service.
	ImageType string `q:"image_type"`
	// Filtering based on policy_id is supported.
	PolicyId string `q:"policy_id"`
	// Searching based on the VM's IP address is supported.
	Ip string `q:"ip"`
	// Filtering based on checkpoint_id is supported.
	CheckpointId string `q:"checkpoint_id"`
	// Type of the backup object. For example, OS::Nova::Server
	ResourceType string `q:"resource_type"`
}

func CountBackups(client *golangsdk.ServiceClient, opts CountOpts) (*int, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{endpoint}/v1/{project_id}/checkpoint_items/count
	raw, err := client.Get(client.ServiceURL("checkpoint_items", "count")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Count int `json:"count"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Count, err
}
