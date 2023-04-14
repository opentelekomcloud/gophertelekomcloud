package listeners

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get retrieves a particular Listeners based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("listeners", id), &r.Body, nil)
	return
}

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
	//
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
