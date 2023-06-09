package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// HTTPS certificate ID.
	// It can be obtained by calling the ListCertificates API.
	// This parameter is not required when the client protocol is HTTP,
	// but it is mandatory when the client protocol is HTTPS.
	CertificateId string `json:"certificateid"`
	// Certificate name.
	// Certifacteid and certificatename are required at the same.
	// If certificateid does not match certificatename, an error is reported.
	// This parameter is not required when the client protocol is HTTP,
	// but it is mandatory when the client protocol is HTTPS.
	CertificateName string `json:"certificatename"`
	// Protected domain name or IP address (port allowed)
	Hostname string `json:"hostname" required:"true"`
	// Whether a proxy is used for the domain name.
	// If your website has no layer-7 proxy server
	// such as CDN and cloud acceleration service deployed
	// in front of WAF and uses only layer-4 load balancers
	// (or NAT), set Proxy Configured to No. Otherwise,
	// Proxy Configured must be set to Yes.
	// This ensures that WAF obtains real IP addresses of website
	// visitors and takes protective actions configured in
	// protection policies.
	Proxy *bool `json:"proxy" required:"true"`
	// ID of the policy initially used to the domain name.
	// It can be obtained by calling the API described in 2.1.1
	// Querying Protection Policies.
	PolicyId string `json:"policyid"`
	// Server configuration in dedicated mode
	Server []PremiumWafServer `json:"server" required:"true"`
}

type PremiumWafServer struct {
	// Client protocol
	// Enumeration values:
	// HTTP
	// HTTPS
	FrontProtocol string `json:"front_protocol" required:"true"`
	// Server protocol
	// Enumeration values:
	// HTTP
	// HTTPS
	BackProtocol string `json:"back_protocol" required:"true"`
	// IP address or domain name of the origin server that the client accesses.
	Address string `json:"address" required:"true"`
	// Server port
	Port int `json:"port" required:"true"`
	// The origin server address is an IPv4 or IPv6 address. Default value: ipv4
	// Enumeration values:
	// ipv4
	// ipv6
	Type string `json:"type" required:"true"`
	// VPC ID. Perform the following steps to obtain the VPC ID:
	// 1.Find the name of the VPC where the dedicated engine is located. The VPC name is in the VPC\Subnet column. Log in to the WAF console and choose Instance Management > Dedicated Engine > VPC\Subnet.
	// Log in to the VPC console and click the VPC name. On the page displayed, copy the VPC ID in the VPC Information area.
	VpcId string `json:"vpc_id" required:"true"`
}

// Create will create a new Protected Domain Name on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*HostResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/premium-waf/host
	raw, err := client.Post(client.ServiceURL("host"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res HostResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type HostResponse struct {
	// Protected domain name ID
	ID string `json:"id"`
	// Policy ID
	PolicyId string `json:"policyid"`
	// Protected domain name
	Hostname string `json:"hostname"`
	// Tenant ID
	DomainId string `json:"domainid"`
	// Project ID
	ProjectId string `json:"projectid"`
	// HTTP protocol
	Protocol string `json:"protocol"`
	// WAF status of the protected domain name.
	// -1: Bypassed. Requests are directly sent to the backend servers without passing through WAF.
	// 0: Suspended. WAF only forwards requests for the domain name but does not detect attacks.
	// -1: Enabled. WAF detects attacks based on the configured policy.
	ProtectStatus int `json:"protect_status"`
	// Whether a domain name is connected to WAF.
	// 0: disconnected
	// 1: connected
	AccessStatus int `json:"access_status"`
	// Whether a proxy is used.
	// true: The proxy is enabled.
	// false: The proxy is disabled.
	Proxy bool `json:"proxy"`
	// Origin server list
	Server []ServerResponse `json:"server"`
	// Special domain name identifier, which is used to store additional domain name configuration.
	Flag *FlagResponse `json:"flag"`
	// Alarm configuration page
	BlockPage *BlockPageResponse `json:"block_page"`
	// Not described
	Extend map[string]string `json:"extend"`
	// Creation time.
	CreatedAt int `json:"timestamp"`
}

type ServerResponse struct {
	// Client protocol
	// Enumeration values:
	// HTTP
	// HTTPS
	FrontProtocol string `json:"front_protocol"`
	// Server protocol
	// Enumeration values:
	// HTTP
	// HTTPS
	BackProtocol string `json:"back_protocol"`
	// IP address or domain name of the origin server that the client accesses.
	Address string `json:"address"`
	// Server port
	Port int `json:"port"`
	// The origin server address is an IPv4 or IPv6 address. Default value: ipv4
	// Enumeration values:
	// ipv4
	// ipv6
	Type string `json:"type"`
	// VPC ID. Perform the following steps to obtain the VPC ID:
	// 1.Find the name of the VPC where the dedicated engine is located. The VPC name is in the VPC\Subnet column. Log in to the WAF console and choose Instance Management > Dedicated Engine > VPC\Subnet.
	// Log in to the VPC console and click the VPC name. On the page displayed, copy the VPC ID in the VPC Information area.
	VpcId string `json:"vpc_id"`
}

type FlagResponse struct {
	// Whether PCI 3DS certification check is enabled for the domain name. Currently, this function is not supported. The default value is false. You can ignore this parameter.
	// true: PCI 3DS check is enabled.
	// false: PCI 3DS check is disabled.
	Pci3ds string `json:"pci_3ds"`
	// Whether PCI DDS certification check is enabled for the domain name.
	// true: PCI DDS check is enabled.
	// false: PCI DDS check is disabled.
	PciDss string `json:"pci_dss"`
}
type BlockPageResponse struct {
	// Template name
	Template string `json:"template"`
	// Custom alarm page
	CustomPage *CustomPageResponse `json:"custom_page"`
	// Redirection URL
	RedirectUrl string `json:"redirect_url"`
}

type CustomPageResponse struct {
	StatusCode  string `json:"status_code"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
}
