package listeners

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	//
	// Default: 2000
	Limit int `q:"limit"`
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies whether to use reverse query. Values:
	//
	// true: Query the previous page.
	//
	// false (default): Query the next page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If page_reverse is set to true and you want to query the previous page, set the value of marker to the value of previous_marker.
	PageReverse bool `q:"page_reverse"`
	// Specifies the port used by the listener.
	//
	// Multiple ports can be queried in the format of protocol_port=xxx&protocol_port=xxx.
	ProtocolPort []int `q:"protocol_port"`
	// Specifies the protocol used by the listener. The value can be TCP, HTTP, UDP, HTTPS or TERMINATED_HTTPS. Note: TERMINATED_HTTPS is only available for the listeners of shared load balancers.
	//
	// Multiple protocols can be queried in the format of protocol=xxx&protocol=xxx.
	Protocol []string `q:"protocol"`
	// Provides supplementary information about the listener.
	//
	// Multiple descriptions can be queried in the format of description=xxx&description=xxx.
	Description []string `q:"description"`
	// Specifies the ID of the server certificate used by the listener.
	//
	// Multiple IDs can be queried in the format of default_tls_container_ref=xxx&default_tls_container_ref=xxx.
	DefaultTLSContainerRef []string `q:"default_tls_container_ref"`
	// Specifies the ID of the CA certificate used by the listener.
	//
	// Multiple IDs can be queried in the format of client_ca_tls_container_ref=xxx&client_ca_tls_container_ref=xxx.
	ClientCATLSContainerRef []string `q:"client_ca_tls_container_ref"`
	// Specifies the administrative status of the listener. The value can only be true.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp *bool `q:"admin_state_up"`
	// Specifies the maximum number of connections that the load balancer can establish with backend servers. The value -1 indicates that the number of connections is not limited.
	//
	// Multiple values can be queried in the format of connection_limit=xxx&connection_limit=xxx.
	//
	// This parameter is unsupported. Please do not use it.
	ConnectionLimit []int `q:"connection_limit"`
	// Specifies the ID of the default backend server group. If there is no matched forwarding policy, requests will be routed to the default backend server.
	//
	// Multiple IDs can be queried in the format of default_pool_id=xxx&default_pool_id=xxx.
	DefaultPoolID []string `q:"default_pool_id"`
	// Specifies the listener ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	ID []string `q:"id"`
	// Specifies the name of the listener added to the load balancer.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Specifies whether to use HTTP/2 if you want the clients to use HTTP/2 to communicate with the listener. However, connections between the load balancer and backend servers still use HTTP/1.x by default.
	//
	// This parameter is available only for HTTPS listeners. If you configure this parameter for listeners with other protocols, it will not take effect.
	Http2Enable *bool `q:"http2_enable"`
	// Specifies the ID of the load balancer that the listener is added to.
	//
	// Multiple IDs can be queried in the format of loadbalancer_id=xxx&loadbalancer_id=xxx.
	LoadBalancerID []string `q:"loadbalancer_id"`
	// Specifies the security policy used by the listener.
	//
	// Multiple security policies can be queried in the format of tls_ciphers_policy=xxx&tls_ciphers_policy=xxx.
	TLSCiphersPolicy []string `q:"tls_ciphers_policy"`
	// Specifies the private IP address bound to the backend server. This parameter is used only as a query condition and is not included in the response.
	//
	// Multiple IP addresses can be queried in the format of member_address=xxx&member_address=xxx.
	MemberAddress []string `q:"member_address"`
	// Specifies the ID of the cloud server that serves as a backend server. This parameter is used only as a query condition and is not included in the response.
	//
	// Multiple IDs can be queried in the format of member_device_id=xxx&member_device_id=xxx.
	MemberDeviceID []string `q:"member_device_id"`
	// Specifies whether to enable health check retries for backend servers.
	//
	// The value can be true (enable health check retries) or false (disable health check retries).
	EnableMemberRetry *bool `q:"enable_member_retry"`
	// Specifies the timeout duration for waiting for a response from a backend server, in seconds. If the backend server fails to respond after the timeout duration elapses, the load balancer will stop waiting and return HTTP 504 Gateway Timeout to the client.
	//
	// The value ranges from 1 to 300.
	//
	// Multiple durations can be queried in the format of member_timeout=xxx&member_timeout=xxx.
	MemberTimeout []int `q:"member_timeout"`
	// Specifies the timeout duration for waiting for a response from a client, in seconds. There are two situations:
	//
	// If the client fails to send a request header to the load balancer within the timeout duration, the request will be interrupted.
	//
	// If the interval between two consecutive request bodies reaching the load balancer is greater than the timeout duration, the connection will be disconnected.
	//
	// The value ranges from 1 to 300.
	//
	// Multiple durations can be queried in the format of client_timeout=xxx&client_timeout=xxx.
	ClientTimeout []int `q:"client_timeout"`
	// Specifies the idle timeout duration, in seconds. If there are no requests reaching the load balancer after the idle timeout duration elapses, the load balancer will disconnect the connection with the client and establish a new connection when there is a new request.
	//
	// For TCP listeners, the value ranges from 10 to 4000.
	//
	// For HTTP, HTTPS, and TERMINATED_HTTPS listeners, the value ranges from 1 to 4000.
	//
	// For UDP listeners, this parameter does not take effect.
	//
	// Multiple durations can be queried in the format of keepalive_timeout=xxx&keepalive_timeout=xxx.
	KeepAliveTimeout []int `q:"keepalive_timeout"`
	// Specifies whether to pass source IP addresses of the clients to backend servers.
	//
	// This parameter is only available for TCP or UDP listeners of shared load balancers.
	//
	// true: Source IP addresses will be passed to backend servers.
	//
	// false: Source IP addresses will not be passed to backend servers.
	TransparentClientIpEnable *bool `q:"transparent_client_ip_enable"`
	// Specifies whether to enable advanced forwarding. If you enable this function, you can configure more flexible forwarding policies and rules.
	//
	// true: Enable advanced forwarding.
	//
	// false: Disable advanced forwarding. This parameter is not available in eu-nl region. Please do not use it.
	EnhanceL7policyEnable *bool `q:"enhance_l7policy_enable"`
	// Specifies the backend server ID. This parameter is used only as a query condition and is not included in the response. Multiple IDs can be queried in the format of member_instance_id=xxx&member_instance_id=xxx.
	MemberInstanceId []string `q:"member_instance_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v3/{project_id}/elb/listeners
	return pagination.NewPager(client, client.ServiceURL("listeners")+q.String(), func(r pagination.PageResult) pagination.Page {
		return ListenerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
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
	var res []Listener
	err := extract.IntoSlicePtr(r.(ListenerPage).BodyReader(), &res, "listeners")
	return res, err
}
