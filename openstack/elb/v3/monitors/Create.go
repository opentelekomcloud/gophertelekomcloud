package monitors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts is the common options' struct used in this package's Create operation.
type CreateOpts struct {
	// Specifies the administrative status of the health check.
	//
	// true (default): Health check is enabled.
	//
	// false: Health check is disabled.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the interval between health checks, in seconds. The value ranges from 1 to 50.
	Delay *int `json:"delay" required:"true"`
	// Specifies the domain name that HTTP requests are sent to during the health check.
	//
	// The value can contain only digits, letters, hyphens (-), and periods (.) and must start with a digit or letter.
	//
	// The value is left blank by default, indicating that the virtual IP address of the load balancer is used as the destination address of HTTP requests.
	//
	// This parameter is available only when type is set to HTTP.
	//
	// Minimum: 1
	//
	// Maximum: 100
	DomainName string `json:"domain_name,omitempty"`
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
	//
	// Minimum: 1
	//
	// Maximum: 64
	ExpectedCodes string `json:"expected_codes,omitempty"`
	// Specifies the HTTP method. The value can be GET, HEAD, POST, PUT, DELETE, TRACE, OPTIONS, CONNECT, or PATCH. The default value is GET.
	//
	// This parameter is available when type is set to HTTP or HTTPS.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Minimum: 1
	//
	// Maximum: 16
	HttpMethod string `json:"http_method,omitempty"`
	// Specifies the number of consecutive health checks when the health check result of a backend server changes from OFFLINE to ONLINE.
	//
	// The value ranges from 1 to 10.
	//
	// Minimum: 1
	//
	// Maximum: 10
	MaxRetries *int `json:"max_retries" required:"true"`
	// Specifies the number of consecutive health checks when the health check result of a backend server changes from ONLINE to OFFLINE.
	//
	// The value ranges from 1 to 10, and the default value is 3.
	//
	// Minimum: 1
	//
	// Maximum: 10
	//
	// Default: 3
	MaxRetriesDown *int `json:"max_retries_down,omitempty"`
	// Specifies the port used for the health check. If this parameter is left blank, a port of the backend server will be used by default. The port number ranges from 1 to 65535.
	//
	// Minimum: 1
	//
	// Maximum: 65535
	MonitorPort *int `json:"monitor_port,omitempty"`
	// Specifies the health check name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Specifies the ID of the backend server group for which the health check is configured.
	PoolId string `json:"pool_id" required:"true"`
	// Specifies the project ID.
	//
	// Minimum: 1
	//
	// Maximum: 32
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the maximum time required for waiting for a response from the health check, in seconds. It is recommended that you set the value less than that of parameter delay.
	//
	// Minimum: 1
	//
	// Maximum: 50
	Timeout *int `json:"timeout" required:"true"`
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
	Type string `json:"type" required:"true"`
	// Specifies the HTTP request path for the health check. The value must start with a slash (/), and the default value is /. Note: This parameter is available only when type is set to HTTP.
	//
	// Default: /
	//
	// Minimum: 1
	//
	// Maximum: 80
	UrlPath string `json:"url_path,omitempty"`
}

// Create is an operation which provisions a new Health Monitor. There are
// different types of Monitor you can provision: PING, TCP or HTTP(S). Below
// are examples of how to create each one.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Monitor, error) {
	b, err := build.RequestBody(opts, "healthmonitor")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("healthmonitors"), b, nil, nil)
	return extra(err, raw)
}
