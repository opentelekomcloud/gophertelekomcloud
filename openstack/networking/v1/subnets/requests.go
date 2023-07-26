package subnets

import (
	"encoding/json"
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	// ID is the unique identifier for the subnet.
	ID string `json:",omitempty"`

	// Name is the human readable name for the subnet. It does not have to be
	// unique.
	Name string `json:",omitempty"`

	// Specifies the network segment on which the subnet resides.
	CIDR string `json:",omitempty"`

	// Status indicates whether or not a subnet is currently operational.
	Status string `json:",omitempty"`

	// Specifies the gateway of the subnet.
	GatewayIP string `json:",omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet.
	PrimaryDNS string `json:",omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet.
	SecondaryDNS string `json:",omitempty"`

	// Identifies the availability zone (AZ) to which the subnet belongs.
	AvailabilityZone string `json:",omitempty"`

	// Specifies the ID of the VPC to which the subnet belongs.
	VpcID string `json:",omitempty"`
}

// List returns collection of
// subnets. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those subnets that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Subnet, error) {
	url := rootURL(client)
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return SubnetPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}.AllPages()
	if err != nil {
		return nil, err
	}

	allSubnets, err := ExtractSubnets(pages)
	if err != nil {
		return nil, err
	}

	return FilterSubnets(allSubnets, opts)
}

func FilterSubnets(subnets []Subnet, opts ListOpts) ([]Subnet, error) {
	matchOptsByte, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	var matchOpts map[string]any
	err = json.Unmarshal(matchOptsByte, &matchOpts)
	if err != nil {
		return nil, err
	}

	if len(matchOpts) == 0 {
		return subnets, nil
	}

	var refinedSubnets []Subnet
	for _, subnet := range subnets {
		if subnetMatchesFilter(&subnet, matchOpts) {
			refinedSubnets = append(refinedSubnets, subnet)
		}
	}
	return refinedSubnets, nil
}

func subnetMatchesFilter(subnet *Subnet, filter map[string]any) bool {
	for key, expectedValue := range filter {
		if getStructField(subnet, key) != expectedValue {
			return false
		}
	}
	return true
}

func getStructField(v *Subnet, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSubnetCreateMap() (map[string]any, error)
}

// CreateOpts contains all the values needed to create a new subnets. There are
// no required values.
type CreateOpts struct {
	Name             string         `json:"name" required:"true"`
	Description      string         `json:"description,omitempty"`
	CIDR             string         `json:"cidr" required:"true"`
	DNSList          []string       `json:"dnsList,omitempty"`
	GatewayIP        string         `json:"gateway_ip" required:"true"`
	EnableDHCP       *bool          `json:"dhcp_enable,omitempty"`
	PrimaryDNS       string         `json:"primary_dns,omitempty"`
	SecondaryDNS     string         `json:"secondary_dns,omitempty"`
	AvailabilityZone string         `json:"availability_zone,omitempty"`
	VpcID            string         `json:"vpc_id" required:"true"`
	ExtraDHCPOpts    []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

type ExtraDHCPOpt struct {
	OptName  string `json:"opt_name" required:"true"`
	OptValue string `json:"opt_value,omitempty"`
}

// ToSubnetCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToSubnetCreateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "subnet")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical subnets. When it is created, the subnets does not have an internal
// interface - it is not associated to any subnet.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSubnetCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get retrieves a particular subnets based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSubnetUpdateMap() (map[string]any, error)
}

// UpdateOpts contains the values used when updating a subnets.
type UpdateOpts struct {
	Name          string         `json:"name,omitempty"`
	Description   *string        `json:"description,omitempty"`
	EnableDHCP    *bool          `json:"dhcp_enable,omitempty"`
	PrimaryDNS    string         `json:"primary_dns,omitempty"`
	SecondaryDNS  string         `json:"secondary_dns,omitempty"`
	DNSList       []string       `json:"dnsList,omitempty"`
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

// ToSubnetUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToSubnetUpdateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "subnet")
}

// Update allows subnets to be updated. You can update the name, administrative
// state, and the external gateway.
func Update(client *golangsdk.ServiceClient, vpcID string, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSubnetUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, vpcID, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular subnets based on its unique ID.
func Delete(client *golangsdk.ServiceClient, vpcID string, id string) (r DeleteResult) {
	_, r.Err = client.Delete(updateURL(client, vpcID, id), nil)
	return
}
