package shares

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToShareCreateMap() (map[string]any, error)
}

// CreateOpts contains the options for create an SFS Turbo. This object is
// passed to shares.Create().
type CreateOpts struct {
	// Defines the SFS Turbo file system name
	Name string `json:"name" required:"true"`
	// Defines the SFS Turbo file system protocol to use, the valid value is NFS.
	ShareProto string `json:"share_proto,omitempty"`
	// ShareType defines the file system type. the valid values are STANDARD and PERFORMANCE.
	ShareType string `json:"share_type" required:"true"`
	// Size in GB, range from 500 to 32768.
	Size int `json:"size" required:"true"`
	// The availability zone of the SFS Turbo file system
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// The VPC ID
	VpcID string `json:"vpc_id" required:"true"`
	// The subnet ID
	SubnetID string `json:"subnet_id" required:"true"`
	// The security group ID
	SecurityGroupID string `json:"security_group_id" required:"true"`
	// The enterprise project ID
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// The backup ID
	BackupID string `json:"backup_id,omitempty"`
	// Share description
	Description string `json:"description,omitempty"`
	// The metadata information
	Metadata Metadata `json:"metadata,omitempty"`
}

// Metadata specifies the metadata information
type Metadata struct {
	CryptKeyID string `json:"crypt_key_id,omitempty"`
}

// ToShareCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToShareCreateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "share")
}

// Create will create a new SFS Turbo file system based on the values in CreateOpts. To extract
// the Share object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToShareCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

type ListOptsBuilder interface {
	ToShareListQuery() (string, error)
}

type ListOpts struct {
	Limit  string `q:"limit"`
	Offset string `q:"offset"`
}

func (opts ListOpts) ToShareListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// SFS Turbo resources.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToShareListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return TurboPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}

// Get will get a single SFS Turbo file system with given UUID
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// Delete will delete an existing SFS Turbo file system with the given UUID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// ExpandOptsBuilder allows extensions to add additional parameters to the
// Expand request.
type ExpandOptsBuilder interface {
	ToShareExpandMap() (map[string]any, error)
}

// ExpandOpts contains the options for expanding a SFS Turbo. This object is
// passed to shares.Expand().
type ExpandOpts struct {
	// Specifies the extend object.
	Extend ExtendOpts `json:"extend" required:"true"`
}

type ExtendOpts struct {
	// Specifies the post-expansion capacity (GB) of the shared file system.
	NewSize int `json:"new_size" required:"true"`
}

// ToShareExpandMap assembles a request body based on the contents of a
// ExpandOpts.
func (opts ExpandOpts) ToShareExpandMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Expand will expand a SFS Turbo based on the values in ExpandOpts.
func Expand(client *golangsdk.ServiceClient, shareID string, opts ExpandOptsBuilder) (r ExpandResult) {
	b, err := opts.ToShareExpandMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, shareID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// ChangeSGOptsBuilder allows extensions to change the security group
// bound to a SFS Turbo file system
type ChangeSGOptsBuilder interface {
	ToShareSGMap() (map[string]any, error)
}

// ChangeSGOpts contains the options for changing security group to a SFS Turbo
type ChangeSGOpts struct {
	// Specifies the change_security_group object.
	ChangeSecurityGroup SecurityGroupOpts `json:"change_security_group" required:"true"`
}

type SecurityGroupOpts struct {
	// Specifies the ID of the security group to be modified.
	SecurityGroupID string `json:"security_group_id" required:"true"`
}

// ToShareExpandMap assembles a request body based on the contents of a
// ChangeSGOpts.
func (opts ChangeSGOpts) ToShareSGMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ChangeSG will change security group to a SFS Turbo based on the values in ChangeSGOpts.
func ChangeSG(client *golangsdk.ServiceClient, shareID string, opts ChangeSGOptsBuilder) (r ExpandResult) {
	b, err := opts.ToShareSGMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, shareID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
