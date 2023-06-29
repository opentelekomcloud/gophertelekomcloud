package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Specifies the name of the custom security policy.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the custom security policy.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
	// Lists the TLS protocols supported by the custom security policy. Value options: TLSv1, TLSv1.1, TLSv1.2, and TLSv1.3
	Protocols []string `json:"protocols,omitempty"`
	// Lists the cipher suites supported by the custom security policy. The following cipher suites are supported:
	//
	// ECDHE-RSA-AES256-GCM-SHA384,ECDHE-RSA-AES128-GCM-SHA256,ECDHE-ECDSA-AES256-GCM-SHA384,ECDHE-ECDSA-AES128-GCM-SHA256,AES128-GCM-SHA256,AES256-GCM-SHA384,ECDHE-ECDSA-AES128-SHA256,ECDHE-RSA-AES128-SHA256,AES128-SHA256,AES256-SHA256,ECDHE-ECDSA-AES256-SHA384,ECDHE-RSA-AES256-SHA384,ECDHE-ECDSA-AES128-SHA,ECDHE-RSA-AES128-SHA,ECDHE-RSA-AES256-SHA,ECDHE-ECDSA-AES256-SHA,AES128-SHA,AES256-SHA,CAMELLIA128-SHA,DES-CBC3-SHA,CAMELLIA256-SHA,ECDHE-RSA-CHACHA20-POLY1305,ECDHE-ECDSA-CHACHA20-POLY1305,TLS_AES_128_GCM_SHA256,TLS_AES_256_GCM_SHA384,TLS_CHACHA20_POLY1305_SHA256,TLS_AES_128_CCM_SHA256,TLS_AES_128_CCM_8_SHA256
	//
	// Note:
	//
	// The protocol and cipher suite must match. At least one cipher suite must match the protocol.
	//
	// You can match the protocol and cipher suite based on system security policy.
	Ciphers []string `json:"ciphers,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts, id string) (*SecurityPolicy, error) {
	b, err := build.RequestBody(opts, "security_policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("security-policies", id), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}
