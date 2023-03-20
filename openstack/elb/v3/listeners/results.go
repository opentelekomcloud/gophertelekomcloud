package listeners

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Listener is the primary load balancing configuration object that specifies
// the loadbalancer and port on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type Listener struct {
	// The unique ID for the Listener.
	ID string `json:"id"`

	// The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	AdminStateUp bool `json:"admin_state_up"`

	// The ID of the CA certificate used by the listener.
	CAContainerRef string `json:"client_ca_tls_container_ref"`

	// The maximum number of connections allowed for the Loadbalancer.
	// Default is -1, meaning no limit.
	ConnectionLimit int `json:"connection_limit"`

	// Specifies the time when the listener was created.
	CreatedAt string `json:"created_at"`

	// Specifies the time when the listener was updated.
	UpdatedAt string `json:"updated_at"`

	// The UUID of default pool. Must have compatible protocol with listener.
	DefaultPoolID string `json:"default_pool_id"`

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef string `json:"default_tls_container_ref"`

	// Provides supplementary information about the Listener.
	Description string `json:"description"`

	// whether to use HTTP2.
	Http2Enable bool `json:"http2_enable"`

	// A list of load balancer IDs.
	Loadbalancers []structs.ResourceRef `json:"loadbalancers"`

	// Specifies the Listener Name.
	Name string `json:"name"`

	// Specifies the ProjectID where the listener is used.
	ProjectID string `json:"project_id"`

	// The protocol to loadbalancer. A valid value is TCP, HTTP, or HTTPS.
	Protocol string `json:"protocol"`

	// The port on which to listen to client traffic that is associated with the
	// Loadbalancer. A valid value is from 0 to 65535.
	ProtocolPort int `json:"protocol_port"`

	// The list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs"`

	// SNI certificates wildcard.
	SniMatchAlgo string `json:"sni_match_algo"`

	// Lists the Tags.
	Tags []tags.ResourceTag `json:"tags"`

	// Specifies the security policy used by the listener.
	TlsCiphersPolicy string `json:"tls_ciphers_policy"`

	SecurityPolicy string `json:"security_policy_id"`

	// Whether enable member retry
	EnableMemberRetry bool `json:"enable_member_retry"`

	// The keepalive timeout of the Listener.
	KeepAliveTimeout int `json:"keepalive_timeout"`

	// The client timeout of the Listener.
	ClientTimeout int `json:"client_timeout"`

	// The member timeout of the Listener.
	MemberTimeout int `json:"member_timeout"`

	// The ipgroup of the Listener.
	IpGroup IpGroup `json:"ipgroup"`

	// The http insert headers of the Listener.
	InsertHeaders InsertHeaders `json:"insert_headers"`

	// Transparent client ip enable
	TransparentClientIP bool `json:"transparent_client_ip_enable"`

	// Enhance L7policy enable
	EnhanceL7policy bool `json:"enhance_l7policy_enable"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a listener.
func (r commonResult) Extract() (*Listener, error) {
	s := new(Listener)
	err := r.ExtractIntoStructPtr(s, "listener")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Listener.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Listener.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Listener.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

type ListenerPage struct {
	pagination.PageWithInfo
}

func (p ListenerPage) IsEmpty() (bool, error) {
	l, err := ExtractListeners(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

func ExtractListeners(r pagination.Page) ([]Listener, error) {
	var s []Listener
	err := (r.(ListenerPage)).ExtractIntoSlicePtr(&s, "listeners")
	if err != nil {
		return nil, err
	}
	return s, nil
}
