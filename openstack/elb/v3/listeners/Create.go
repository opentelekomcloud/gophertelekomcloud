package listeners

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts represents options for creating a listener.
type CreateOpts struct {
	// Specifies the administrative status of the listener. The value can only be true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the ID of the CA certificate used by the listener. This parameter is available only when type is set to client.
	//
	// Minimum: 1
	//
	// Maximum: 128
	CAContainerRef string `json:"client_ca_tls_container_ref,omitempty"`
	// Specifies the ID of the default backend server group. If there is no matched forwarding policy, requests will be forwarded to the default backend server for processing.
	//
	// Minimum: 1
	//
	// Maximum: 36
	DefaultPoolID string `json:"default_pool_id,omitempty"`
	// Specifies the ID of the server certificate used by the listener. This parameter is available only when the listener's protocol is HTTPS and type is set to server.
	//
	// Minimum: 1
	//
	// Maximum: 128
	DefaultTlsContainerRef string `json:"default_tls_container_ref,omitempty"`
	// Provides supplementary information about the listener.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
	// Specifies whether to use HTTP/2 if you want the clients to use HTTP/2 to communicate with the load balancer. However, connections between the load balancer and backend servers still use HTTP/1.x by default.
	//
	// This parameter is available only for HTTPS listeners. If you configure this parameter for listeners with other protocols, it will not take effect.
	Http2Enable *bool `json:"http2_enable,omitempty"`
	// Specifies the ID of the load balancer that the listener is added to. Note: A listener can be added to only one load balancer.
	//
	// Minimum: 1
	//
	// Maximum: 36
	LoadbalancerID string `json:"loadbalancer_id" required:"true"`
	// Specifies the listener name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the project ID.
	//
	// Minimum: 1
	//
	// Maximum: 32
	ProjectID string `json:"project_id,omitempty"`
	// Specifies the protocol used by the listener.
	//
	// The value can be TCP, HTTP, UDP, HTTPS, or TERMINATED_HTTPS.
	//
	// Note:
	//
	// Protocol used by HTTPS listeners added to a shared load balancer can only be set to TERMINATED_HTTPS. If HTTPS is passed, the value will be automatically changed to TERMINATED_HTTPS.
	//
	// Protocol used by HTTPS listeners added to a dedicated load balancer can only be set to HTTPS. If TERMINATED_HTTPS is passed, the value will be automatically changed to HTTPS.
	Protocol string `json:"protocol" required:"true"`
	// Specifies the protocol used by the listener.
	//
	// The value can be TCP, HTTP, UDP, HTTPS, or TERMINATED_HTTPS.
	//
	// Note:
	//
	// Protocol used by HTTPS listeners added to a shared load balancer can only be set to TERMINATED_HTTPS. If HTTPS is passed, the value will be automatically changed to TERMINATED_HTTPS.
	//
	// Protocol used by HTTPS listeners added to a dedicated load balancer can only be set to HTTPS. If TERMINATED_HTTPS is passed, the value will be automatically changed to HTTPS.
	//
	// Minimum: 1
	//
	// Maximum: 65535
	ProtocolPort int `json:"protocol_port" required:"true"`
	// Specifies the IDs of SNI certificates (server certificates with domain names) used by the listener.
	//
	// Note:
	//
	// The domain names of all SNI certificates must be unique.
	//
	// The total number of domain names of all SNI certificates cannot exceed 30.
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`
	// Specifies how wildcard domain name matches with the SNI certificates used by the listener.
	//
	// longest_suffix indicates longest suffix match. wildcard indicates wildcard match.
	//
	// The default value is wildcard.
	SniMatchAlgo string `json:"sni_match_algo,omitempty"`
	// A list of Tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Specifies the security policy used by the listener.
	//
	// Values: tls-1-0-inherit,tls-1-0, tls-1-1, tls-1-2,tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0, and tls-1-0 (default).
	//
	// Note:
	//
	// This parameter will take effect only for HTTPS listeners added to a dedicated load balancer.
	//
	// If both security_policy_id and tls_ciphers_policy are specified, only security_policy_id will take effect.
	//
	// The priority of the encryption suite from high to low is: ecc suite, rsa suite, tls 1.3 suite (supporting both ecc and rsa).
	TlsCiphersPolicy string `json:"tls_ciphers_policy,omitempty"`
	// Specifies the ID of the custom security policy.
	//
	// Note:
	//
	// This parameter is available only for HTTPS listeners added to a dedicated load balancer.
	//
	// If both security_policy_id and tls_ciphers_policy are specified, only security_policy_id will take effect.
	//
	// The priority of the encryption suite from high to low is: ecc suite: ecc suite, rsa suite, tls 1.3 suite (supporting both ecc and rsa).
	//
	// Minimum: 1
	//
	// Maximum: 36
	SecurityPolicy string `json:"security_policy_id,omitempty"`
	// Specifies whether to enable health check retries for backend servers. The value can be true (enable health check retries) or false (disable health check retries). The default value is true. Note:
	//
	// If a shared load balancer is associated, this parameter is available only when protocol is set to HTTP or TERMINATED_HTTPS.
	//
	// If a dedicated load balancer is associated, this parameter is available only when protocol is set to HTTP, or HTTPS.
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`
	// Specifies the idle timeout duration, in seconds. If there are no requests reaching the load balancer after the idle timeout duration elapses, the load balancer will disconnect the connection with the client and establish a new connection when there is a new request.
	//
	// For TCP listeners, the value ranges from 10 to 4000, and the default value is 300.
	//
	// For HTTP and HTTPS listeners, the value ranges from 1 to 4000, and the default value is 60.
	//
	// For UDP listeners, this parameter does not take effect.
	KeepAliveTimeout int `json:"keepalive_timeout,omitempty"`
	// Specifies the timeout duration for waiting for a response from a client, in seconds. There are two situations:
	//
	// If the client fails to send a request header to the load balancer within the timeout duration, the request will be interrupted.
	//
	// If the interval between two consecutive request bodies reaching the load balancer is greater than the timeout duration, the connection will be disconnected.
	//
	// The value ranges from 1 to 300, and the default value is 60.
	//
	// This parameter is available only for HTTP and HTTPS listeners.
	//
	// Minimum: 1
	//
	// Maximum: 300
	//
	// Default: 60
	ClientTimeout int `json:"client_timeout,omitempty"`
	// Specifies the timeout duration for waiting for a response from a backend server, in seconds. If the backend server fails to respond after the timeout duration elapses, the load balancer will stop waiting and return HTTP 504 Gateway Timeout to the client.
	//
	// The value ranges from 1 to 300, and the default value is 60.
	//
	// This parameter is available only for HTTP and HTTPS listeners.
	//
	// Minimum: 1
	//
	// Maximum: 300
	//
	// Default: 60
	MemberTimeout int `json:"member_timeout,omitempty"`
	// Specifies the IP address group associated with the listener.
	IpGroup *IpGroup `json:"ipgroup,omitempty"`
	// Specifies the HTTP header fields that can transmit required information to backend servers. For example, the X-Forwarded-ELB-IP header field can transmit the EIP of the load balancer to backend servers.
	InsertHeaders *InsertHeaders `json:"insert_headers,omitempty"`
	// Specifies whether to pass source IP addresses of the clients to backend servers.
	//
	// TCP or UDP listeners of shared load balancers: The value can be true or false, and the default value is false if this parameter is not passed.
	//
	// HTTP or HTTPS listeners of shared load balancers: The value can only be true, and the default value is true if this parameter is not passed.
	//
	// All listeners of dedicated load balancers: The value can only be true, and the default value is true if this parameter is not passed.
	//
	// Note:
	//
	// If this function is enabled, the load balancer communicates with backend servers using their real IP addresses. Ensure that security group rules and access control policies are correctly configured.
	//
	// If this function is enabled, a server cannot serve as both a backend server and a client.
	//
	// If this function is enabled, backend server specifications cannot be changed.
	TransparentClientIP *bool `json:"transparent_client_ip_enable,omitempty"`
	// Specifies whether to enable advanced forwarding. If advanced forwarding is enabled, more flexible forwarding policies and rules are supported. The value can be true (enable advanced forwarding) or false (disable advanced forwarding), and the default value is false.
	//
	// The following scenarios are supported:
	//
	// action can be set to REDIRECT_TO_URL (requests will be redirected to another URL) or Fixed_RESPONSE (a fixed response body will be returned to clients).
	//
	// Parameters priority, redirect_url_config, and fixed_response_config can be specified in a forwarding policy.
	//
	// Parameter type can be set to METHOD, HEADER, QUERY_STRING, or SOURCE_IP for a forwarding rule.
	//
	// If type is set to HOST_NAME for a forwarding rule, the value parameter of the forwarding rule supports wildcard asterisks (*).
	//
	// The conditions parameter can be specified for forwarding rules. This parameter is not available in eu-nl region. Please do not use it.
	EnhanceL7policy *bool `json:"enhance_l7policy_enable,omitempty"`
	// Specifies the QUIC configuration for the current listener. This parameter is valid only when protocol is set to HTTPS.
	//
	// For a TCP/UDP/HTTP/QUIC listener, if this parameter is not left blank, an error will be reported.
	//
	// Note
	//
	// The client sends a normal HTTP request that contains information indicating that the QUIC protocol is supported.
	//
	// If QUIC upgrade is enabled for the listeners, QUIC port and version information will be added to the response header.
	//
	// When the client sends both HTTPS and QUIC requests to the server, if the QUIC request is successfully sent, QUIC protocol will be used for subsequent communications.
	//
	// QUIC protocol is not supported.
	QuicConfig QuicConfigOption `json:"quic_config,omitempty"`
}

type IpGroup struct {
	// Specifies the ID of the IP address group associated with the listener.
	//
	// If ip_list is set to an empty array [] and type to whitelist, no IP addresses are allowed to access the listener.
	//
	// If ip_list is set to an empty array [] and type to blacklist, any IP address is allowed to access the listener.
	//
	// Minimum: 1
	//
	// Maximum: 36
	IpGroupID string `json:"ipgroup_id" required:"true"`
	// Specifies whether to enable access control.
	//
	// true (default): Access control will be enabled.
	//
	// false: Access control will be disabled.
	Enable *bool `json:"enable_ipgroup,omitempty"`
	// Specifies how access to the listener is controlled.
	//
	// white (default): A whitelist will be configured. Only IP addresses in the whitelist can access the listener.
	//
	// black: A blacklist will be configured. IP addresses in the blacklist are not allowed to access the listener.
	Type string `json:"type,omitempty"`
}

type InsertHeaders struct {
	// Specifies whether to transparently transmit the load balancer EIP to backend servers. If X-Forwarded-ELB-IP is set to true, the load balancer EIP will be stored in the HTTP header and passed to backend servers.
	//
	// Default: false
	ForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`
	// Specifies whether to transparently transmit the listening port of the load balancer to backend servers. If X-Forwarded-Port is set to true, the listening port of the load balancer will be stored in the HTTP header and passed to backend servers.
	//
	// Default: false
	ForwardedPort *bool `json:"X-Forwarded-Port,omitempty"`
	// Specifies whether to transparently transmit the source port of the client to backend servers. If X-Forwarded-For-Port is set to true, the source port of the client will be stored in the HTTP header and passed to backend servers.
	//
	// Default: false
	ForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`
	// Specifies whether to rewrite the X-Forwarded-Host header. If X-Forwarded-Host is set to true, X-Forwarded-Host in the request header from the clients can be set to Host in the request header sent from the load balancer to backend servers.
	//
	// Default: true
	ForwardedHost *bool `json:"X-Forwarded-Host" required:"true"`
}

type QuicConfigOption struct {
	// Specifies the ID of the QUIC listener. Specifies the specified listener. The specified quic_listener_id must exist. The listener protocol must be QUIC and cannot be set to null, otherwise, it will conflict with enable_quic_upgrade.
	//
	// QUIC protocol is not supported.
	QuicListenerId string `json:"quic_listener_id" required:"true"`
	//
	// Specifies whether to enable QUIC upgrade. True: QUIC upgrade is enabled. False (default): QUIC upgrade is disabled. HTTPS listeners can be upgraded to QUIC listeners.
	//
	// QUIC protocol is not supported.
	//
	// Default: false
	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty"`
}

// Create is an operation which provisions a new Listeners based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Listeners on behalf of other tenants by
// specifying a TenantID attribute different from their own.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Listener, error) {
	b, err := build.RequestBody(opts, "listener")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/elb/listeners
	raw, err := client.Post(client.ServiceURL("listeners"), b, nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Listener, error) {
	if err != nil {
		return nil, err
	}

	var res Listener
	err = extract.IntoStructPtr(raw.Body, &res, "listener")
	return &res, err
}
