package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToConfigCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new configuration.
type CreateOpts struct {
	// Configuration Name
	Name string `json:"name" required:"true"`
	// Configuration Description
	Description string `json:"description,omitempty"`
	// Configuration Values
	Values map[string]string `json:"values,omitempty"`
	// Database Object
	DataStore DataStore `json:"datastore" required:"true"`
}

type DataStore struct {
	// DB Engine
	Type string `json:"type" required:"true"`
	// DB version
	Version string `json:"version" required:"true"`
}

// ToConfigCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToConfigCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToConfigUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a Configuration.
type UpdateOpts struct {
	// Configuration Name
	Name string `json:"name,omitempty"`
	// Configuration Description
	Description string `json:"description,omitempty"`
	// Configuration Values
	Values map[string]string `json:"values,omitempty"`
}

// ToConfigUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToConfigUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ApplyOptsBuilder allows extensions to update instance templates to the
// Apply request.
type ApplyOptsBuilder interface {
	ToInstanceTemplateApplyMap() (map[string]interface{}, error)
}

// ApplyOpts contains all the instances needed to apply another template.
type ApplyOpts struct {
	// Specifies the DB instance ID list object.
	InstanceIDs []string `json:"instance_ids" required:"true"`
}

// ToInstanceTemplateApplyMap builds an apply request body from ApplyOpts.
func (opts ApplyOpts) ToInstanceTemplateApplyMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new Config based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToConfigCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders, JSONBody: nil,
	})
	return
}

// Update accepts a UpdateOpts struct and uses the values to update a Configuration.The response code from api is 200
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToConfigUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.Put(resourceURL(c, id), b, nil, reqOpt)
	return
}

// Get retrieves a particular Configuration based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, openstack.StdRequestOpts())
	return
}

// GetForInstance retrieves Configuration applied to particular RDS instance
// configuration ID and Name will be empty
func GetForInstance(c *golangsdk.ServiceClient, instanceID string) (r GetResult) {
	_, r.Err = c.Get(instanceConfigURL(c, instanceID), &r.Body, openstack.StdRequestOpts())
	return
}

// Delete will permanently delete a particular Configuration based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}, MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.Delete(resourceURL(c, id), reqOpt)
	return
}

// List is used to obtain the parameter template list, including default
// parameter templates of all databases and those created by users.
func List(client *golangsdk.ServiceClient) (r ListResult) {
	_, r.Err = client.Get(rootURL(client), &r.Body, nil)
	return
}

// Apply is used to apply a parameter template to one or more DB instances.
func Apply(client *golangsdk.ServiceClient, id string, opts ApplyOptsBuilder) (r ApplyResult) {
	b, err := opts.ToInstanceTemplateApplyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(applyURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
