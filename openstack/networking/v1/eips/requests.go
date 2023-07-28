package eips

import (
	"encoding/json"
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToEipListQuery() (string, error)
}

// ListOpts filters the results returned by the List() function.
type ListOpts struct {
	// ID is the unique identifier for the ElasticIP.
	ID string `json:",omitempty"`

	// Status indicates whether or not a ElasticIP is currently operational.
	Status string `json:",omitempty"`

	// PrivateAddress of the resource with assigned ElasticIP.
	PrivateAddress string `json:",omitempty"`

	// PublicAddress of the ElasticIP.
	PublicAddress string `json:",omitempty"`

	// PortID of the resource with assigned ElasticIP.
	PortID string `json:",omitempty"`

	// BandwidthID of the ElasticIP.
	BandwidthID string `json:",omitempty"`
}

// ToEipListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToEipListQuery() (string, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List instructs OpenStack to provide a list of flavors.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]PublicIp, error) {
	url := rootURL(client)
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			p := EipPage{pagination.MarkerPageBase{PageResult: r}}
			p.MarkerPageBase.Owner = p
			return p
		},
	}.AllPages()
	if err != nil {
		return nil, err
	}

	allPublicIPs, err := ExtractEips(pages)
	if err != nil {
		return nil, err
	}

	return FilterPublicIPs(allPublicIPs, opts)
}

func FilterPublicIPs(publicIPs []PublicIp, opts ListOpts) ([]PublicIp, error) {
	matchOptsByte, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	var matchOpts map[string]interface{}
	err = json.Unmarshal(matchOptsByte, &matchOpts)
	if err != nil {
		return nil, err
	}

	if len(matchOpts) == 0 {
		return publicIPs, nil
	}

	var refinedPublicIPs []PublicIp
	for _, publicIP := range publicIPs {
		if publicIPMatchesFilter(&publicIP, matchOpts) {
			refinedPublicIPs = append(refinedPublicIPs, publicIP)
		}
	}
	return refinedPublicIPs, nil
}

func publicIPMatchesFilter(publicIP *PublicIp, filter map[string]interface{}) bool {
	for key, expectedValue := range filter {
		if getStructField(publicIP, key) != expectedValue {
			return false
		}
	}
	return true
}

func getStructField(v *PublicIp, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

// ApplyOptsBuilder is an interface by which can build the request body of public ip
// application
type ApplyOptsBuilder interface {
	ToPublicIpApplyMap() (map[string]interface{}, error)
}

// ApplyOpts is a struct which is used to create public ip
type ApplyOpts struct {
	IP        PublicIpOpts  `json:"publicip" required:"true"`
	Bandwidth BandwidthOpts `json:"bandwidth" required:"true"`
}

type PublicIpOpts struct {
	Type    string `json:"type" required:"true"`
	Address string `json:"ip_address,omitempty"`
}

type BandwidthOpts struct {
	Name       string `json:"name,omitempty"`
	Size       int    `json:"size,omitempty"`
	Id         string `json:"id,omitempty"`
	ShareType  string `json:"share_type" required:"true"`
	ChargeMode string `json:"charge_mode,omitempty"`
}

func (opts ApplyOpts) ToPublicIpApplyMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

// Apply is a method by which can access to apply the public ip
func Apply(client *golangsdk.ServiceClient, opts ApplyOptsBuilder) (r ApplyResult) {
	b, err := opts.ToPublicIpApplyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get is a method by which can get the detailed information of public ip
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// Delete is a method by which can be able to delete a private ip
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// UpdateOptsBuilder is an interface by which can be able to build the request
// body
type UpdateOptsBuilder interface {
	ToPublicIpUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is a struct which represents the request body of update method
type UpdateOpts struct {
	PortID string `json:"port_id,omitempty"`
}

func (opts UpdateOpts) ToPublicIpUpdateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "publicip")
}

// Update is a method which can be able to update the port of public ip
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPublicIpUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
