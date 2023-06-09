package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Host, error) {
	// GET /v1/{project_id}/premium-waf/host
	raw, err := client.Get(client.ServiceURL("host", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Host
	return &res, extract.Into(raw.Body, &res)
}

type Host struct {
	// Domain name ID
	ID string `json:"id"`
	// ID of the policy initially used to the domain name.
	// It can be obtained by calling the API described in 2.1.1 Querying Protection Policies.
	PolicyId string `json:"policyid"`
	// Domain name added to cloud WAF.
	Hostname string `json:"hostname"`
	// User domain ID.
	DomainId string `json:"domainid"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// HTTP protocol.
	Protocol string `json:"protocol"`
	// Minimum TLS version supported.
	// TLS v1.0 is used by default.
	// The value can be:TLS v1.0TLS v1.1TLS v1.2TLS v1.3
	Tls string `json:"tls"`
	// Cipher suite. The value can be:
	// cipher_1: ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH
	// cipher_2: EECDH+AESGCM:EDH+AESGCM
	// cipher_3: ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH
	// cipher_4. ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH n - cipher_default: ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher string `json:"cipher"`
	// Origin server details
	Server []ServerResponse `json:"server"`
	// HTTPS certificate ID.
	// It can be obtained by calling the ListCertificates API.
	// n - This parameter is not required when the client protocol is HTTP.
	// n - This parameter is mandatory when the client protocol is HTTPS.
	CertificateId string `json:"certificateid"`
	// Certificate name.
	// n - This parameter is not required when the client protocol is HTTP.
	// n - This parameter is mandatory when the client protocol is HTTPS.
	CertificateName string `json:"certificatename"`
	// Whether the proxy is enabled
	Proxy bool `json:"proxy"`
	// Lock status. This parameter is redundant and can be ignored. Default value: 0
	Locked int `json:"locked"`
	// WAF status of the protected domain name. The value can be:
	// -1: Bypassed. Requests are directly sent to the backend servers without passing through WAF.
	// 0: Suspended. WAF only forwards requests for the domain name but does not detect attacks.
	// 1: Enabled. WAF detects attacks based on the configured policy.
	ProtectStatus int `json:"protect_status"`
	// Whether a domain name is connected to WAF.
	// 0: The domain name is not connected to the engine instance.
	// 1: The domain name is connected to the engine instance.
	AccessStatus int `json:"access_status"`
	// Time a domain name is added to WAF
	CreatedAt int `json:"timestamp"`
	// Special domain name identifier, which is used to store additional domain name configurations
	Flag *FlagResponse `json:"flag"`
	// Alarm configuration page
	BlockPage *BlockPageResponse `json:"block_page"`
	// Extended attribute
	Extend map[string]string `json:"extend"`
	// WAF mode. The value is premium, indicating
	// the dedicated WAF engine
	WafType string `json:"waf_type"`
}
