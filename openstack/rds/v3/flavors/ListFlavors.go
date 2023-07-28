package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ListOpts struct {
	DatabaseName string
	// Specifies the database version.
	VersionName string `q:"version_name"`
	// Specifies the specification code.
	// NOTICE
	// Only DB instances running Microsoft SQL Server 2017 EE support the creation of read replicas.
	// Microsoft SQL Server 2017 EE does not support the creation of single DB instances.
	// The format of the specification code is: {spec code}{instance mode}.
	// spec code can be obtained from Table 8-10.
	// instance mode can be any of the following:
	// For single DB instances, the value is null. Example spe_code: rds.mysql.s1.xlarge
	// For primary/standby DB instances, the value is .ha. Example spe_code: rds.mysql.s1.xlarge.ha
	// For read replicas, the value is .rr. Example spe_code: rds.mysql.s1.xlarge.rr
	SpecCode string `q:"spec_code"`
}

func ListFlavors(client *golangsdk.ServiceClient, opts ListOpts) ([]Flavor, error) {
	// GET https://{Endpoint}/v3/{project_id}/flavors/{database_name}
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("flavors", opts.DatabaseName)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Flavor
	err = extract.IntoSlicePtr(raw.Body, &res, "flavors")
	return res, err
}

type Flavor struct {
	// Indicates the CPU size. For example, the value 1 indicates 1 vCPU.
	VCPUs string `json:"vcpus"`
	// Indicates the memory size in GB.
	RAM int `json:"ram"`
	// Indicates the specification ID, which is unique.
	Id string `json:"id"`
	// Indicates the resource specification code. Use rds.mysql.m1.xlarge.rr as an example.
	// rds: indicates the RDS product.
	// mysql: indicates the DB engine.
	// m1.xlarge: indicates the high memory performance specifications.
	// rr: indicates the read replica (.ha indicates primary/standby DB instances).
	SpecCode string `json:"spec_code"`
	// Indicates the database version.
	// Example value for MySQL: ["5.6","5.7","8.0"]
	VersionName []string `json:"version_name"`
	// Indicates the DB instance type. Its value can be any of the following:
	// ha: indicates primary/standby DB instances.
	// replica: indicates read replicas.
	// single: indicates single DB instances.
	InstanceMode string `json:"instance_mode"`
	// Indicates the specification status in an AZ. Its value can be any of the following:
	// normal: indicates that the specifications in the AZ are available.
	// unsupported: indicates that the specifications are not supported by the AZ.
	// sellout: indicates that the specifications in the AZ are sold out.
	AzStatus map[string]string `json:"az_status"`
	// Indicates the description of the AZ to which the specifications belong.
	AzDesc map[string]string `json:"az_desc"`
	// Indicates the performance specifications. Its value can be any of the following:
	// normal: general-enhanced
	GroupType string `json:"group_type"`
}
