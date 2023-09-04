package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is the common options' struct used in this package's Create
// operation.
type CreateOpts struct {
	// Specifies the IP address group name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the IP address group.
	Description string `json:"description,omitempty"`

	// Specifies the project ID of the IP address group.
	ProjectId string `json:"project_id,omitempty"`

	// Specifies the IP addresses or CIDR blocks in the IP address group. [] indicates any IP address.
	IpList *[]IpGroupOption `json:"ip_list,omitempty"`
}

type IpGroupOption struct {
	// Specifies the IP addresses in the IP address group.
	Ip string `json:"ip" required:"true"`

	// Provides remarks about the IP address group.
	Description string `json:"description"`
}

type DataStore struct {
	// Specifies the DB engine. Its value can be any of the following and is case-insensitive:
	// MySQL
	// PostgreSQL
	// SQLServer
	Type string `json:"type" required:"true"`
	// Specifies the database version.
	// Example values:
	// MySQL: 8.0
	// PostgreSQL: 13
	// SQLServer: 2017_SE
	Version string `json:"version" required:"true"`
}

// Create is an operation which provisions a new IP address group based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// IpGroup will be returned.
func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*IpGroup, error) {
	b, err := build.RequestBodyMap(opts, "ipgroup")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("ipgroups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res IpGroup
	err = extract.IntoStructPtr(raw.Body, &res, "ipgroup")
	return &res, err
}
