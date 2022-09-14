package servers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ServerResult struct {
	golangsdk.Result
}

func ExtractSer(err error, raw *http.Response) (*Server, error) {
	if err != nil {
		return nil, err
	}

	var res Server
	err = extract.IntoStructPtr(raw.Body, &res, "server")
	return &res, err
}

func (r ServerResult) ExtractInto(v interface{}) error {
	return extract.IntoStructPtr(r.BodyReader(), v, "server")
}

// Server represents a server/instance in the OpenStack cloud.
type Server struct {
	// ID uniquely identifies this server amongst all other servers,
	// including those not accessible to the current tenant.
	ID string `json:"id"`
	// TenantID identifies the tenant owning this server resource.
	TenantID string `json:"tenant_id"`
	// UserID uniquely identifies the user account owning the tenant.
	UserID string `json:"user_id"`
	// Name contains the human-readable name for the server.
	Name string `json:"name"`
	// Updated and Created contain ISO-8601 timestamps of when the state of the
	// server last changed, and when it was created.
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	// HostID is the host where the server is located in the cloud.
	HostID string `json:"hostid"`
	// Status contains the current operational status of the server,
	// such as IN_PROGRESS or ACTIVE.
	Status string `json:"status"`
	// Progress ranges from 0..100.
	// A request made against the server completes only once Progress reaches 100.
	Progress int `json:"progress"`
	// AccessIPv4 and AccessIPv6 contain the IP addresses of the server,
	// suitable for remote access for administration.
	AccessIPv4 string `json:"accessIPv4"`
	AccessIPv6 string `json:"accessIPv6"`
	// Image refers to a JSON object, which itself indicates the OS image used to deploy the server.
	Image map[string]interface{} `json:"-"`
	// Flavor refers to a JSON object, which itself indicates the hardware
	// configuration of the deployed server.
	Flavor map[string]interface{} `json:"flavor"`
	// Addresses includes a list of all IP addresses assigned to the server, keyed by pool.
	Addresses map[string]interface{} `json:"addresses"`
	// Metadata includes a list of all user-specified key-value pairs attached to the server.
	Metadata map[string]string `json:"metadata"`
	// Links includes HTTP references to itself, useful for passing along to
	// other APIs that might want a server reference.
	Links []interface{} `json:"links"`
	// KeyName indicates which public key was injected into the server on launch.
	KeyName string `json:"key_name"`
	// AdminPass will generally be empty ("").  However, it will contain the
	// administrative password chosen when provisioning a new server without a
	// set AdminPass setting in the first place.
	// Note that this is the ONLY time this field will be valid.
	AdminPass string `json:"adminPass"`
	// SecurityGroups includes the security groups that this instance has applied to it.
	SecurityGroups []map[string]interface{} `json:"security_groups"`
	// Fault contains failure information about a server.
	Fault Fault `json:"fault"`
	// VolumeAttached includes the volumes that attached to the server.
	VolumesAttached []map[string]string `json:"os-extended-volumes:volumes_attached"`
}

type Fault struct {
	Code    int       `json:"code"`
	Created time.Time `json:"created"`
	Details string    `json:"details"`
	Message string    `json:"message"`
}

func (r *Server) UnmarshalJSON(b []byte) error {
	type tmp Server
	var s struct {
		tmp
		Image interface{} `json:"image"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Server(s.tmp)

	switch t := s.Image.(type) {
	case map[string]interface{}:
		r.Image = t
	case string:
		switch t {
		case "":
			r.Image = nil
		}
	}

	return err
}

type NIC struct {
	PortState  string    `json:"port_state"`
	FixedIPs   []FixedIP `json:"fixed_ips"`
	PortID     string    `json:"port_id"`
	NetID      string    `json:"net_id"`
	MACAddress string    `json:"mac_addr"`
}

type FixedIP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

// ServerPage abstracts the raw results of making a List() request against
// the API. As OpenStack extensions may freely alter the response bodies of
// structures returned to the client, you may only safely access the data
// provided through the ExtractServers call.
type ServerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no Server results.
func (r ServerPage) IsEmpty() (bool, error) {
	s, err := ExtractServers(r)
	return len(s) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (r ServerPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &res, "servers_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// ExtractServers interprets the results of a single page from a List() call, producing a slice of Server entities.
func ExtractServers(r pagination.Page) ([]Server, error) {
	var s []Server
	err := ExtractServersInto(r, &s)
	return s, err
}

func ExtractServersInto(r pagination.Page, v interface{}) error {
	return extract.IntoSlicePtr(r.(ServerPage).BodyReader(), v, "servers")
}

func extraTum(err error, raw *http.Response) (map[string]string, error) {
	if err != nil {
		return nil, err
	}

	var res struct {
		Metadatum map[string]string `json:"meta"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Metadatum, err
}

func extraMet(err error, raw *http.Response) (map[string]string, error) {
	if err != nil {
		return nil, err
	}

	var res struct {
		Metadata map[string]string `json:"metadata"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Metadata, err
}

// Address represents an IP address.
type Address struct {
	Version int    `json:"version"`
	Address string `json:"addr"`
}

// MarshalJSON marshals the escaped file, base64 encoding the contents.
func (f *File) MarshalJSON() ([]byte, error) {
	file := struct {
		Path     string `json:"path"`
		Contents string `json:"contents"`
	}{
		Path:     f.Path,
		Contents: base64.StdEncoding.EncodeToString(f.Contents),
	}
	return json.Marshal(file)
}
