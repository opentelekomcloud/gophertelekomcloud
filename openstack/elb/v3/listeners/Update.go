package listeners

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts represents options for updating a Listener.
type UpdateOpts struct {
	// Specifies the administrative status of the listener. The value can only be true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the ID of the CA certificate used by the listener. This parameter is available only when type is set to client.
	CAContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`
	// Specifies the ID of the default backend server group. If there is no matched forwarding policy, requests are forwarded to the default backend server.
	//
	// Minimum: 1
	//
	// Maximum: 36
	DefaultPoolID string `json:"default_pool_id,omitempty"`
	// Specifies the ID of the server certificate used by the listener. This parameter is available only when the listener's protocol is HTTPS and type is set to server.
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// Provides supplementary information about the listener.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description *string `json:"description,omitempty"`
	// Specifies whether to use HTTP/2 if you want the clients to use HTTP/2 to communicate with the load balancer. However, connections between the load balancer and backend servers still use HTTP/1.x by default.
	//
	// This parameter is available only for HTTPS listeners. If you configure this parameter for listeners with other protocols, it will not take effect.
	Http2Enable *bool `json:"http2_enable,omitempty"`
	// Specifies the listener name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name *string `json:"name,omitempty"`
	// Specifies the IDs of SNI certificates (server certificates with domain names) used by the listener.
	//
	// Note:
	//
	// The domain names of all SNI certificates must be unique.
	//
	// The total number of domain names of all SNI certificates cannot exceed 30.
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`
	// Specifies how wildcard domain name matches with the SNI certificates used by the listener.
	//
	// longest_suffix indicates longest suffix match. wildcard indicates wildcard match.
	//
	// The default value is wildcard.
	SniMatchAlgo string `json:"sni_match_algo,omitempty"`
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
	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`
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
	// For TCP listeners, the value ranges from 10 to 4000.
	//
	// For HTTP and HTTPS listeners, the value ranges from 1 to 4000.
	//
	// For UDP listeners, this parameter does not take effect.
	KeepAliveTimeout int `json:"keepalive_timeout,omitempty"`
	// Specifies the timeout duration for waiting for a response from a client, in seconds.
	//
	// This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300.
	//
	// Minimum: 1
	//
	// Maximum: 300
	ClientTimeout int `json:"client_timeout,omitempty"`
	// Specifies the timeout duration for waiting for a response from a backend server, in seconds. If the backend server fails to respond after the timeout duration elapses, the load balancer will stop waiting and return HTTP 504 Gateway Timeout to the client.
	//
	// The value ranges from 1 to 300.
	//
	// This parameter is available only for HTTP and HTTPS listeners.
	//
	// Minimum: 1
	//
	// Maximum: 300
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
	QuicConfig QuicConfigOption `json:"quic_config"`
}

// Update is an operation which modifies the attributes of the specified Listener.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Listener, error) {
	// PUT /v3/{project_id}/elb/listeners/{listener_id}
	b, err := build.RequestBody(opts, "listener")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("listeners", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extra(err, raw)
}
