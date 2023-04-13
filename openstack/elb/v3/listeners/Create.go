package listeners

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options for creating a listener.
type CreateOpts struct {
	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// the ID of the CA certificate used by the listener.
	CAContainerRef string `json:"client_ca_tls_container_ref,omitempty"`

	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID string `json:"default_pool_id,omitempty"`

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef string `json:"default_tls_container_ref,omitempty"`

	// Provides supplementary information about the Listener.
	Description string `json:"description,omitempty"`

	// whether to use HTTP2.
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// The load balancer on which to provision this listener.
	LoadbalancerID string `json:"loadbalancer_id" required:"true"`

	// Specifies the Listener name.
	Name string `json:"name,omitempty"`

	// ProjectID is only required if the caller has an admin role and wants
	// to create a pool for another project.
	ProjectID string `json:"project_id,omitempty"`

	// The protocol - can either be TCP, HTTP or HTTPS.
	Protocol Protocol `json:"protocol" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`

	// A list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`

	// Specifies how wildcard domain name matches with the SNI certificates used by the listener.
	// longest_suffix indicates longest suffix match. wildcard indicates wildcard match.
	// The default value is wildcard.
	SniMatchAlgo string `json:"sni_match_algo,omitempty"`

	// A list of Tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`

	// Specifies the security policy used by the listener.
	TlsCiphersPolicy string `json:"tls_ciphers_policy,omitempty"`

	// Specifies the ID of the custom security policy.
	// Note:
	// This parameter is available only for HTTPS listeners added to a dedicated load balancer.
	// If both security_policy_id and tls_ciphers_policy are specified, only security_policy_id will take effect.
	// The priority of the encryption suite from high to low is: ecc suite: ecc suite, rsa suite, tls 1.3 suite (supporting both ecc and rsa).
	SecurityPolicy string `json:"security_policy_id,omitempty"`

	// Whether enable member retry
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	// The keepalive timeout of the Listener.
	KeepAliveTimeout int `json:"keepalive_timeout,omitempty"`

	// The client timeout of the Listener.
	ClientTimeout int `json:"client_timeout,omitempty"`

	// The member timeout of the Listener.
	MemberTimeout int `json:"member_timeout,omitempty"`

	// The IpGroup of the Listener.
	IpGroup *IpGroup `json:"ipgroup,omitempty"`

	// The http insert headers of the Listener.
	InsertHeaders *InsertHeaders `json:"insert_headers,omitempty"`

	// Transparent client ip enable
	TransparentClientIP *bool `json:"transparent_client_ip_enable,omitempty"`

	// Enhance L7policy enable
	EnhanceL7policy *bool `json:"enhance_l7policy_enable,omitempty"`
}

// ToListenerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "listener")
}

// Create is an operation which provisions a new Listeners based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Listeners on behalf of other tenants by
// specifying a TenantID attribute different from their own.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("listeners"), b, &r.Body, nil)
	return
}
