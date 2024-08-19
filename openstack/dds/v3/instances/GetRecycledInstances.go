package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListRecycledOpts struct {
	// Index offset. If offset is set to N, the resource query starts from the N+1 piece of data.
	// The default value is 0, indicating that the query starts from the first piece of data. The value must be a positive integer.
	Offset int `q:"offset"`
	// Number of records displayed on each page. The default value is 100.
	Limit int `q:"limit"`
}

func ListRecycledInstances(client *golangsdk.ServiceClient, opts ListRecycledOpts) (*ListRecycledResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("recycle-instances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/recycle-instances
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListRecycledResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListRecycledResponse struct {
	Instances  []RecycledInstanceResponse `json:"instances"`
	TotalCount int                        `json:"total_count"`
}

type RecycledInstanceResponse struct {
	// Indicates the DB instance ID.
	Id string `json:"id"`
	// Indicates the DB instance name.
	Name string `json:"name"`
	// Instance type. Cluster, replica set, and single node instances are supported. The value can be:
	// Sharding
	// ReplicaSet
	// Single
	Mode string `json:"mode"`
	// Database information.
	Datastore DataStore `json:"datastore"`
	// Billing mode.
	// 0: indicates the instance is billed on a pay-per-use basis.
	// 1: indicates the instance is billed based on a yearly/monthly basis.
	PayMode string `json:"pay_mode"`
	// Backup ID.
	BackupId string `json:"backup_id"`
	// Creation time.
	CreatedAt string `json:"created_at"`
	// Deletion time.
	DeletedAt string `json:"deleted_at"`
	// Retention end time.
	RetainedUntil string `json:"retained_until"`
	// Instance backup recycling status.
	Status string `json:"status"`
}
