package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeStorageTypeOpts struct {
	InstanceId string `json:"-"`
	// Specification code of the target storage type.
	// Values for RDS for MySQL:
	// -- If the target storage type is ESSD:
	// * rds.mysql.volume.essd.ha: ESSD specification code for primary/standby DB instances
	// * rds.mysql.volume.essd.rr: ESSD specification code for read replicas
	// -- If the target storage type is cloud SSD:
	// * rds.mysql.volume.cloudssd.ha: cloud SSD specification code for primary/standby DB instances
	// * rds.mysql.volume.cloudssd.rr: cloud SSD specification code for read replicas
	// * rds.mysql.volume.cloudssd: cloud SSD specification code for single-node DB instances
	// Values for RDS for PostgreSQL:
	// -- If the target storage type is ESSD:
	// * rds.pg.volume.essd.ha: ESSD specification code for primary/standby DB instances
	// * rds.pg.volume.essd.rr: ESSD specification code for read replicas
	// -- If the target storage type is cloud SSD:
	// * rds.pg.volume.cloudssd.ha: cloud SSD specification code for primary/standby DB instances
	// * rds.pg.volume.cloudssd.rr: cloud SSD specification code for read replicas
	// * rds.pg.volume.cloudssd: cloud SSD specification code for single-node DB instances
	// Values for RDS for SQL Server:
	// -- If the target storage type is ESSD:
	// * rds.mssql.volume.essd.ha: ESSD specification code for primary/standby DB instances
	// * rds.mssql.volume.essd.rr: ESSD specification code for read replicas
	// * rds.mssql.volume.essd: ESSD specification code for single-node DB instances
	// -- If the target storage type is cloud SSD:
	// * rds.mssql.volume.cloudssd.ha: cloud SSD specification code for primary/standby DB instances
	// * rds.mssql.volume.cloudssd.rr: cloud SSD specification code for read replicas
	// * rds.mssql.volume.cloudssd: cloud SSD specification code for single-node DB instances
	VolumeCode string `json:"volume_code"`
}

func ChangeStorageType(client *golangsdk.ServiceClient, opts ChangeStorageTypeOpts) (*string, error) {
	b, err := build.RequestBody(opts, "change_volume")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
