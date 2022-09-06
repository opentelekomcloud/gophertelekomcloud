package attachinterfaces

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type attachInterfaceResult struct {
	golangsdk.Result
}

// Extract interprets any attachInterfaceResult as an Interface, if possible.
func (raw attachInterfaceResult) Extract() (*Interface, error) {
	var res struct {
		Interface *Interface `json:"interfaceAttachment"`
	}
	err = extract.Into(raw, &res)
	return &res, err
}

// GetResult is the response from a Get operation. Call its Extract
// method to interpret it as an Interface.
type GetResult struct {
	attachInterfaceResult
}

// CreateResult is the response from a Create operation. Call its Extract
// method to interpret it as an Interface.
type CreateResult struct {
	attachInterfaceResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// FixedIP represents a Fixed IP Address.
// This struct is also used when creating an attachment,
// but it is not possible to specify a SubnetID.
type FixedIP struct {
	SubnetID  string `json:"subnet_id,omitempty"`
	IPAddress string `json:"ip_address"`
}

// Interface represents a network interface on a server.
type Interface struct {
	PortState string    `json:"port_state"`
	FixedIPs  []FixedIP `json:"fixed_ips"`
	PortID    string    `json:"port_id"`
	NetID     string    `json:"net_id"`
	MACAddr   string    `json:"mac_addr"`
}

// InterfacePage abstracts the raw results of making a List() request against
// the API.
//
// As OpenStack extensions may freely alter the response bodies of structures
// returned to the client, you may only safely access the data provided through
// the ExtractInterfaces call.
type InterfacePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if an InterfacePage contains no interfaces.
func (r InterfacePage) IsEmpty() (bool, error) {
	interfaces, err := ExtractInterfaces(r)
	return len(interfaces) == 0, err
}

// ExtractInterfaces interprets the results of a single page from a List() call,
// producing a slice of Interface structs.
func ExtractInterfaces(r pagination.Page) ([]Interface, error) {
	var res struct {
		Interfaces []Interface `json:"interfaceAttachments"`
	}
	err := (r.(InterfacePage)).ExtractInto(&res)
	return res, err
}
