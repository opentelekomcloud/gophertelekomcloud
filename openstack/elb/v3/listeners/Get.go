package listeners

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get retrieves a particular Listeners based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Listener, error) {
	// GET /v3/{project_id}/elb/listeners/{listener_id}
	raw, err := client.Get(client.ServiceURL("listeners", id), nil, nil)
	return extra(err, raw)
}

// Listener is the primary load balancing configuration object that specifies
// the loadbalancer and port on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type Listener struct {
	// Specifies the listener ID.
	ID string `json:"id"`
	// Specifies the administrative status of the listener. The value can only be true.
	// This parameter is unsupported. Please do not use it.
	AdminStateUp bool `json:"admin_state_up"`
	// Specifies the ID of the CA certificate used by the listener. This parameter is available only when type is set to client.
	CAContainerRef string `json:"client_ca_tls_container_ref"`
	// Specifies the maximum number of connections that the load balancer can establish with backend servers. The value -1 indicates that the number of connections is not limited.
	//
	// This parameter is unsupported. Please do not use it.
	ConnectionLimit int `json:"connection_limit"`
	// Specifies the time when the listener was created, in the format of yyyy-MM-dd''T''HH:mm:ss''Z'', for example, 2021-07-30T12:03:44Z.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the listener was updated.
	UpdatedAt string `json:"updated_at"`
	// Specifies the ID of the default backend server group. If there is no matched forwarding policy, requests are forwarded to the default backend server.
	DefaultPoolID string `json:"default_pool_id"`
	// Specifies the ID of the server certificate used by the listener.
	DefaultTlsContainerRef string `json:"default_tls_container_ref"`
	// Provides supplementary information about the listener.
	Description string `json:"description"`
	// Specifies whether to use HTTP/2 if you want the clients to use HTTP/2 to communicate with the load balancer. However, connections between the load balancer and backend servers still use HTTP/1.x by default.
	//
	// This parameter is available only for HTTPS listeners. If you configure this parameter for listeners with other protocols, it will not take effect.
	Http2Enable bool `json:"http2_enable"`
	// Specifies the ID of the load balancer that the listener is added to. A listener can be added to only one load balancer.
	Loadbalancers []structs.ResourceRef `json:"loadbalancers"`
	// Specifies the Listener Name.
	Name string `json:"name"`
	// Specifies the ProjectID where the listener is used.
	ProjectID string `json:"project_id"`
	// Specifies the protocol used by the listener.
	//
	// The value can be TCP, HTTP, UDP, HTTPS, or TERMINATED_HTTPS.
	//
	// Note:
	//
	// Protocol used by HTTPS listeners added to a shared load balancer can only be set to TERMINATED_HTTPS. If HTTPS is passed, the value will be automatically changed to TERMINATED_HTTPS.
	//
	// Protocol used by HTTPS listeners added to a dedicated load balancer can only be set to HTTPS. If TERMINATED_HTTPS is passed, the value will be automatically changed to HTTPS.
	Protocol string `json:"protocol"`
	// The port on which to listen to client traffic that is associated with the
	// Loadbalancer. A valid value is from 0 to 65535.
	ProtocolPort int `json:"protocol_port"`
	// Specifies the IDs of SNI certificates (server certificates with domain names) used by the listener.
	//
	// Note:
	//
	// The domain names of all SNI certificates must be unique.
	//
	// The total number of domain names of all SNI certificates cannot exceed 30.
	SniContainerRefs []string `json:"sni_container_refs"`
	// Specifies how wildcard domain name matches with the SNI certificates used by the listener.
	//
	// longest_suffix indicates longest suffix match. wildcard indicates wildcard match.
	//
	// The default value is wildcard.
	SniMatchAlgo string `json:"sni_match_algo"`
	// Lists the Tags.
	Tags []tags.ResourceTag `json:"tags"`
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
	TlsCiphersPolicy string `json:"tls_ciphers_policy"`
	// Specifies the ID of the custom security policy.
	//
	// Note:
	//
	// This parameter is available only for HTTPS listeners added to a dedicated load balancer.
	//
	// If both security_policy_id and tls_ciphers_policy are specified, only security_policy_id will take effect.
	//
	// The priority of the encryption suite from high to low is: ecc suite: ecc suite, rsa suite, tls 1.3 suite (supporting both ecc and rsa).
	SecurityPolicy string `json:"security_policy_id"`
	// Specifies whether to enable health check retries for backend servers. The value can be true (enable health check retries) or false (disable health check retries). The default value is true. Note:
	//
	// If a shared load balancer is associated, this parameter is available only when protocol is set to HTTP or TERMINATED_HTTPS.
	//
	// If a dedicated load balancer is associated, this parameter is available only when protocol is set to HTTP, or HTTPS.
	EnableMemberRetry bool `json:"enable_member_retry"`
	// Specifies the idle timeout duration, in seconds. If there are no requests reaching the load balancer after the idle timeout duration elapses, the load balancer will disconnect the connection with the client and establish a new connection when there is a new request.
	//
	// For TCP listeners, the value ranges from 10 to 4000, and the default value is 300.
	//
	// For HTTP and HTTPS listeners, the value ranges from 1 to 4000, and the default value is 60.
	//
	// For UDP listeners, this parameter does not take effect.
	KeepAliveTimeout int `json:"keepalive_timeout"`
	// Specifies the timeout duration for waiting for a response from a client, in seconds. There are two situations:
	//
	// If the client fails to send a request header to the load balancer within the timeout duration, the request will be interrupted.
	//
	// If the interval between two consecutive request bodies reaching the load balancer is greater than the timeout duration, the connection will be disconnected.
	//
	// The value ranges from 1 to 300, and the default value is 60.
	//
	// This parameter is available only for HTTP and HTTPS listeners.
	ClientTimeout int `json:"client_timeout"`
	// Specifies the timeout duration for waiting for a response from a backend server, in seconds. If the backend server fails to respond after the timeout duration elapses, the load balancer will stop waiting and return HTTP 504 Gateway Timeout to the client.
	//
	// The value ranges from 1 to 300, and the default value is 60.
	//
	// This parameter is available only for HTTP and HTTPS listeners.
	MemberTimeout int `json:"member_timeout"`
	// Specifies the IP address group associated with the listener.
	IpGroup IpGroup `json:"ipgroup"`
	// Specifies the HTTP header fields that can transmit required information to backend servers. For example, the X-Forwarded-ELB-IP header field can transmit the EIP of the load balancer to backend servers.
	InsertHeaders InsertHeaders `json:"insert_headers"`
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
	TransparentClientIP bool `json:"transparent_client_ip_enable"`
	// Specifies whether to enable advanced forwarding. The value can be true (enable advanced forwarding) or false (disable advanced forwarding), and the default value is false.
	//
	// If this function is enabled, action can be set to REDIRECT_TO_URL (requests will be redirected to another URL) or Fixed_RESPONSE (a fixed response body will be returned to clients).
	//
	// Parameters priority, redirect_url_config, and fixed_response_config can be specified in a forwarding policy.
	//
	// Parameter type can be set to METHOD, HEADER, QUERY_STRING, or SOURCE_IP for a forwarding rule .
	//
	// If type is set to HOST_NAME for a forwarding rule, the value parameter of the forwarding rule supports wildcard asterisks (*).
	//
	// The conditions parameter can be specified for forwarding rules.
	//
	// This parameter is not available in eu-nl region. Please do not use it.
	//
	// Default: false
	EnhanceL7policy bool `json:"enhance_l7policy_enable"`
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
	QuicConfig QuicConfigOption `json:"quic_config"`
}
