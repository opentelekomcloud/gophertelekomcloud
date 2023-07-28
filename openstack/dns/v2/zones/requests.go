package zones

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToZoneListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
// https://developer.openstack.org/api-ref/dns/
type ListOpts struct {
	// Integer value for the limit of values to return.
	Limit int `q:"limit"`

	// UUID of the zone at which you want to set a marker.
	Marker string `q:"marker"`

	Description string `q:"description"`
	Email       string `q:"email"`
	Name        string `q:"name"`
	SortDir     string `q:"sort_dir"`
	SortKey     string `q:"sort_key"`
	Status      string `q:"status"`
	TTL         int    `q:"ttl"`
	Type        string `q:"type"`
	QueryTags   string `q:"tags"`

	Tags []tags.ResourceTag
}

// ToZoneListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToZoneListQuery() (string, error) {
	var tagList []string
	for _, tag := range opts.Tags {
		tagList = append(tagList, fmt.Sprintf("%s,%s", tag.Key, tag.Value))
	}
	opts.QueryTags = strings.Join(tagList, "|")

	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List implements a zone List request.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(client)
	if opts != nil {
		query, err := opts.ToZoneListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ZonePage{pagination.LinkedPageBase{PageResult: r}}
		},
	}
}

// Get returns information about a zone, given its ID.
func Get(client *golangsdk.ServiceClient, zoneID string) (r GetResult) {
	_, r.Err = client.Get(zoneURL(client, zoneID), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToZoneCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies the attributes used to create a zone.
type CreateOpts struct {
	// Email contact of the zone.
	Email string `json:"email,omitempty"`

	// Description of the zone.
	Description string `json:"description,omitempty"`

	// Name of the zone.
	Name string `json:"name" required:"true"`

	// TTL is the time to live of the zone.
	TTL int `json:"ttl,omitempty"`

	ZoneType string `json:"zone_type,omitempty"`

	Router *RouterOpts `json:"router,omitempty"`

	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

// ToZoneCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToZoneCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Create implements a zone create request.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToZoneCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client), &b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201, 202},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToZoneUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts specifies the attributes to update a zone.
type UpdateOpts struct {
	// Email contact of the zone.
	Email string `json:"email,omitempty"`

	// TTL is the time to live of the zone.
	TTL int `json:"ttl,omitempty"`

	// Description of the zone.
	Description string `json:"description,omitempty"`
}

// ToZoneUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToZoneUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update implements a zone update request.
func Update(client *golangsdk.ServiceClient, zoneID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToZoneUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(zoneURL(client, zoneID), &b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete implements a zone delete request.
func Delete(client *golangsdk.ServiceClient, zoneID string) (r DeleteResult) {
	_, r.Err = client.Delete(zoneURL(client, zoneID), &golangsdk.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: &r.Body,
	})
	return
}

// RouterOptsBuilder allows adding parameters to the associate/disassociate Zone request.
type RouterOptsBuilder interface {
	ToRouterMap() (map[string]interface{}, error)
}

// RouterOpts specifies the required information to associate/disassociate a Router with a Zone.
type RouterOpts struct {
	// Router ID
	RouterID string `json:"router_id" required:"true"`

	// Router Region
	RouterRegion string `json:"router_region,omitempty"`
}

// ToRouterMap constructs a request body from RouterOpts.
func (opts RouterOpts) ToRouterMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "router")
}

// AssociateZone associate a Router with a Zone.
func AssociateZone(client *golangsdk.ServiceClient, zoneID string, opts RouterOptsBuilder) (r AssociateResult) {
	b, err := opts.ToRouterMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(associateURL(client, zoneID), b, nil, nil)
	return
}

// DisassociateZone disassociate a Router with a Zone.
func DisassociateZone(client *golangsdk.ServiceClient, zoneID string, opts RouterOptsBuilder) (r DisassociateResult) {
	b, err := opts.ToRouterMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(disassociateURL(client, zoneID), b, nil, nil)
	return
}
