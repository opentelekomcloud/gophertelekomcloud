package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Whether a proxy is used for the domain name.
	// If your website has no layer-7 proxy server such as CDN and cloud
	// acceleration service deployed in front of WAF and uses only layer-4 load balancers (or NAT),
	// set Proxy Configured to No. Otherwise, Proxy Configured must be set to Yes.
	// This ensures that WAF obtains real IP addresses of website visitors and
	// takes protective actions configured in protection policies.
	Proxy *bool `json:"proxy"`
	// HTTPS certificate ID. It can be obtained by calling the ListCertificates API.
	CertificateId string `json:"certificateid"`
	// HTTPS certificate name. It can be obtained by calling the ListCertificates API.
	// Certifacteid and certificatename are required at the same.
	// If certificateid does not match certificatename, an error is reported.
	CertificateName string `json:"certificatename"`
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
	// WAF status of the protected domain name.
	// -1: Bypassed. Requests are directly sent to the backend servers without passing through WAF.
	// 0: Suspended. WAF only forwards requests for the domain name but does not detect attacks.
	// -1: Enabled. WAF detects attacks based on the configured policy.
	ProtectStatus int `json:"protect_status"`
	// Alarm configuration page.
	BlockPage *BlockPage `json:"block_page"`
}

type BlockPage struct {
	// Template name
	Template string `json:"template" required:"true"`
	// Custom alarm page
	CustomPage *CustomPage `json:"custom_page"`
	// Redirection URL
	RedirectUrl string `json:"redirect_url"`
}

type CustomPage struct {
	// Status Codes
	StatusCode string `json:"status_code" required:"true"`
	// Content type of alarm page
	ContentType string `json:"content_type" required:"true"`
	// Page content
	Content string `json:"content" required:"true"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Host, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/premium-waf/host
	raw, err := client.Put(client.ServiceURL("host", id), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}

	var res Host
	return &res, extract.Into(raw.Body, &res)
}
