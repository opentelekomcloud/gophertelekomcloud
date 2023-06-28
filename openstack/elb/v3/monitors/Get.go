package monitors

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
)

// Get retrieves a particular Health Monitor based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Monitor, error) {
	raw, err := client.Get(client.ServiceURL("healthmonitors", id), nil, nil)
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Monitor, error) {
	if err != nil {
		return nil, err
	}

	var res Monitor
	err = extract.IntoStructPtr(raw.Body, res, "healthmonitor")
	return &res, err
}

// Monitor represents a load balancer health monitor. A health monitor is used
// to determine whether back-end members of the VIP's pool are usable
// for processing a request. A pool can have several health monitors associated
// with it. There are different types of health monitors supported:
//
// PING: used to ping the members using ICMP.
// TCP: used to connect to the members using TCP.
// HTTP: used to send an HTTP request to the member.
// HTTPS: used to send a secure HTTP request to the member.
//
// When a pool has several monitors associated with it, each member of the pool
// is monitored by all these monitors. If any monitor declares the member as
// unhealthy, then the member status is changed to INACTIVE and the member
// won't participate in its pool's load balancing. In other words, ALL monitors
// must declare the member to be healthy for it to stay ACTIVE.
type Monitor struct {
	// Specifies the administrative status of the health check.
	//
	// true(default) indicates that the health check is enabled.
	//
	// false indicates that the health check is disabled.
	AdminStateUp *bool `json:"admin_state_up"`
	// Specifies the interval between health checks, in seconds. The value ranges from 1 to 50.
	//
	// Minimum: 1
	//
	// Maximum: 50
	Delay *int `json:"delay"`
	// Specifies the domain name that HTTP requests are sent to during the health check.
	//
	// The value can contain only digits, letters, hyphens (-), and periods (.) and must start with a digit or letter.
	//
	// The value is left blank by default, indicating that the virtual IP address of the load balancer is used as the destination address of HTTP requests.
	//
	// This parameter is available only when type is set to HTTP.
	DomainName string `json:"domain_name"`
	// Specifies the expected HTTP status code. This parameter will take effect only when type is set to HTTP or HTTPS.
	//
	// The value options are as follows:
	//
	// A specific value, for example, 200
	//
	// A list of values that are separated with commas (,), for example, 200, 202
	//
	// A value range, for example, 200-204
	//
	// The default value is 200. Multiple status codes can be queried in the format of expected_codes=xxx&expected_codes=xxx.
	ExpectedCodes string `json:"expected_codes"`
	// Specifies the HTTP method. The value can be GET, HEAD, POST, PUT, DELETE, TRACE, OPTIONS, CONNECT, or PATCH. The default value is GET.
	//
	// This parameter is available when type is set to HTTP or HTTPS.
	//
	// This parameter is unsupported. Please do not use it.
	HttpMethod string `json:"http_method"`
	// Specifies the health check ID.
	Id string `json:"id"`
	// Specifies the number of consecutive health checks when the health check result of a backend server changes from OFFLINE to ONLINE.
	//
	// The value ranges from 1 to 10
	//
	// Minimum: 1
	//
	// Maximum: 10
	MaxRetries *int `json:"max_retries"`
	// Specifies the number of consecutive health checks when the health check result of a backend server changes from ONLINE to OFFLINE.
	//
	// The value ranges from 1 to 10, and the default value is 3.
	//
	// Minimum: 1
	//
	// Maximum: 10
	MaxRetriesDown *int `json:"max_retries_down"`
	// Specifies the port used for the health check. If this parameter is left blank, a port of the backend server will be used by default. The port number ranges from 1 to 65535.
	//
	// Minimum: 1
	//
	// Maximum: 65535
	MonitorPort *int `json:"monitor_port"`
	// Specifies the health check name.
	Name string `json:"name"`
	// Lists the IDs of backend server groups for which the health check is configured. Only one ID will be returned.
	Pools []structs.ResourceRef `json:"pools"`
	// Specifies the project ID.
	ProjectId string `json:"project_id"`
	// Specifies the maximum time required for waiting for a response from the health check, in seconds. It is recommended that you set the value less than that of parameter delay.
	//
	// Minimum: 1
	//
	// Maximum: 50
	Timeout *int `json:"timeout"`
	// Specifies the health check protocol. The value can be TCP, UDP_CONNECT, HTTP, or HTTPS.
	//
	// Note:
	//
	// If the protocol of the backend server is QUIC, the value can only be UDP_CONNECT.
	//
	// If the protocol of the backend server is UDP, the value can only be UDP_CONNECT.
	//
	// If the protocol of the backend server is TCP, the value can only be TCP, HTTP, or HTTPS.
	//
	// If the protocol of the backend server is HTTP, the value can only be TCP, HTTP, or HTTPS.
	//
	// If the protocol of the backend server is HTTPS, the value can only be TCP, HTTP, or HTTPS.
	//
	// QUIC protocol is not supported in eu-nl region.
	Type string `json:"type"`
	// Specifies the HTTP request path for the health check. The value must start with a slash (/), and the default value is /. Note: This parameter is available only when type is set to HTTP.
	UrlPath string `json:"url_path"`
	// Specifies the time when the health check was configured. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the health check was updated. The format is yyyy-MM-dd'T'HH:mm:ss'Z' (UTC time).
	//
	// This is a new field in this version, and it will not be returned for resources associated with existing dedicated load balancers and for resources associated with existing and new shared load balancers.
	UpdatedAt string `json:"updated_at"`
}
