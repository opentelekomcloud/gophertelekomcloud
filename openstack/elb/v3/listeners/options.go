package listeners

// Protocol represents a listener protocol.
type Protocol string

// Supported attributes for create/update operations.
const (
	ProtocolTCP             Protocol = "TCP"
	ProtocolUDP             Protocol = "UDP"
	ProtocolHTTP            Protocol = "HTTP"
	ProtocolHTTPS           Protocol = "HTTPS"
	ProtocolTerminatedHTTPS Protocol = "TERMINATED_HTTPS"
)

type IpGroup struct {
	IpGroupID string `json:"ipgroup_id" required:"true"`
	Enable    *bool  `json:"enable_ipgroup,omitempty"`
	Type      string `json:"type,omitempty"`
}

type InsertHeaders struct {
	ForwardedELBIP                        *bool  `json:"X-Forwarded-ELB-IP,omitempty"`
	ForwardedPort                         *bool  `json:"X-Forwarded-Port,omitempty"`
	ForwardedForPort                      *bool  `json:"X-Forwarded-For-Port,omitempty"`
	ForwardedHost                         *bool  `json:"X-Forwarded-Host" required:"true"`
	ForwardedProto                        *bool  `json:"X-Forwarded-Proto,omitempty"`
	RealIP                                *bool  `json:"X-Real-IP,omitempty"`
	ForwardedELBID                        *bool  `json:"X-Forwarded-ELB-ID,omitempty"`
	ForwardedTLSCertificateID             *bool  `json:"X-Forwarded-TLS-Certificate-ID,omitempty"`
	ForwardedTLSProtocol                  *bool  `json:"X-Forwarded-TLS-Protocol,omitempty"`
	ForwardedTLSCipher                    *bool  `json:"X-Forwarded-TLS-Cipher,omitempty"`
	ForwardedTLSProtocolAlias             string `json:"X-Forwarded-TLS-Protocol-alias,omitempty"`
	ForwardedTLSCipherAlias               string `json:"X-Forwarded-TLS-Cipher-alias,omitempty"`
	ForwardedForProcessingMode            string `json:"X-Forwarded-For-Processing-Mode,omitempty"`
	ForwardedClientCertSubjectDNEnable    *bool  `json:"X-Forwarded-Clientcert-subjectdn-enable,omitempty"`
	ForwardedClientCertSubjectDNAlias     string `json:"X-Forwarded-Clientcert-subjectdn-alias,omitempty"`
	ForwardedClientCertIssuerDNEnable     *bool  `json:"X-Forwarded-Clientcert-issuerdn-enable,omitempty"`
	ForwardedClientCertIssuerDNAlias      string `json:"X-Forwarded-Clientcert-issuerdn-alias,omitempty"`
	ForwardedClientCertFingerprintEnable  *bool  `json:"X-Forwarded-Clientcert-fingerprint-enable,omitempty"`
	ForwardedClientCertFingerprintAlias   string `json:"X-Forwarded-Clientcert-fingerprint-alias,omitempty"`
	ForwardedClientCertClientVerifyEnable *bool  `json:"X-Forwarded-Clientcert-clientverify-enable,omitempty"`
	ForwardedClientCertClientVerifyAlias  string `json:"X-Forwarded-Clientcert-clientverify-alias,omitempty"`
	ForwardedClientCertSerialNumberEnable *bool  `json:"X-Forwarded-Clientcert-serialnumber-enable,omitempty"`
	ForwardedClientCertSerialNumberAlias  string `json:"X-Forwarded-Clientcert-serialnumber-alias,omitempty"`
	ForwardedClientCertEnable             *bool  `json:"X-Forwarded-Clientcert-enable,omitempty"`
	ForwardedClientCertAlias              string `json:"X-Forwarded-Clientcert-alias,omitempty"`
	ForwardedClientCertCiphersEnable      *bool  `json:"X-Forwarded-Clientcert-ciphers-enable,omitempty"`
	ForwardedClientCertCiphersAlias       string `json:"X-Forwarded-Clientcert-ciphers-alias,omitempty"`
	ForwardedClientCertEndEnable          *bool  `json:"X-Forwarded-Clientcert-end-enable,omitempty"`
	ForwardedClientCertEndAlias           string `json:"X-Forwarded-Clientcert-end-alias,omitempty"`
	ForwardedTLSALPNProtocolEnable        *bool  `json:"X-Forwarded-Tls-Alpn-Protocol-enable,omitempty"`
	ForwardedTLSALPNProtocolAlias         string `json:"X-Forwarded-Tls-Alpn-Protocol-alias,omitempty"`
	ForwardedTLSSNIEnable                 *bool  `json:"X-Forwarded-Tls-Sni-enable,omitempty"`
	ForwardedTLSSNIAlias                  string `json:"X-Forwarded-Tls-Sni-alias,omitempty"`
	ForwardedTLSJA3Enable                 *bool  `json:"X-Forwarded-Tls-Ja3-enable,omitempty"`
	ForwardedTLSJA3Alias                  string `json:"X-Forwarded-Tls-Ja3-alias,omitempty"`
	ForwardedTLSJA4Enable                 *bool  `json:"X-Forwarded-Tls-Ja4-enable,omitempty"`
	ForwardedTLSJA4Alias                  string `json:"X-Forwarded-Tls-Ja4-alias,omitempty"`
}

type IpGroupUpdate struct {
	IpGroupId string `json:"ipgroup_id,omitempty"`
	Enable    *bool  `json:"enable_ipgroup,omitempty"`
	Type      string `json:"type,omitempty"`
}

// UpdateOpts represents options for updating a Listener.
type UpdateOpts struct {
	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// the ID of the CA certificate used by the listener.
	CAContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID string `json:"default_pool_id,omitempty"`

	// A reference to a container of TLS secrets.
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	// Provides supplementary information about the Listener.
	Description *string `json:"description,omitempty"`

	// whether to use HTTP2.
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// Specifies the Listener name.
	Name *string `json:"name,omitempty"`

	// A list of references to TLS secrets.
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	// Specifies how wildcard domain name matches with the SNI certificates used by the listener.
	// longest_suffix indicates longest suffix match. wildcard indicates wildcard match.
	// The default value is wildcard.
	SniMatchAlgo string `json:"sni_match_algo,omitempty"`

	// Specifies the security policy used by the listener.
	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	// Specifies the ID of the custom security policy.
	// Note:
	// This parameter is available only for HTTPS listeners added to a dedicated load balancer.
	// If both security_policy_id and tls_ciphers_policy are specified, only security_policy_id will take effect.
	// The priority of the encryption suite from high to low is: ecc suite: ecc suite, rsa suite, tls 1.3 suite (supporting both ecc and rsa).
	SecurityPolicy string `json:"security_policy_id,omitempty"`

	// Whether enable member retry
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	// The keepalive timeout of the Listener.
	KeepAliveTimeout int `json:"keepalive_timeout,omitempty"`

	// The client timeout of the Listener.
	ClientTimeout int `json:"client_timeout,omitempty"`

	// The member timeout of the Listener.
	MemberTimeout int `json:"member_timeout,omitempty"`

	// The IpGroup of the Listener.
	IpGroup *IpGroupUpdate `json:"ipgroup,omitempty"`

	// The http insert headers of the Listener.
	InsertHeaders *InsertHeaders `json:"insert_headers,omitempty"`

	// Transparent client ip enable
	TransparentClientIP *bool `json:"transparent_client_ip_enable,omitempty"`

	// Enhance L7policy enable
	EnhanceL7policy *bool `json:"enhance_l7policy_enable,omitempty"`

	ProtectionStatus                 string                          `json:"protection_status,omitempty"`
	ProtectionReason                 string                          `json:"protection_reason,omitempty"`
	AccessLogCustomizedHeadersConfig *AccessLogCustomizedHeadersOpts `json:"access_log_customized_headers_config,omitempty"`
}
