package listeners

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Protocol represents a listener protocol.
type Protocol string

// Supported attributes for create/update operations.
const (
	ProtocolTCP   Protocol = "TCP"
	ProtocolUDP   Protocol = "UDP"
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolHTTPS Protocol = "HTTPS"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options for creating a listener.
type CreateOpts struct {
	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// the ID of the CA certificate used by the listener.
	CAContainerRef string `json:"client_ca_tls_container_ref,omitempty"`

	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID string `json:"default_pool_id,omitempty"`

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef string `json:"default_tls_container_ref,omitempty"`

	// Provides supplementary information about the Listener.
	Description string `json:"description,omitempty"`

	// whether to use HTTP2.
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// The load balancer on which to provision this listener.
	LoadbalancerID string `json:"loadbalancer_id" required:"true"`

	// Specifies the Listener name.
	Name string `json:"name,omitempty"`

	// ProjectID is only required if the caller has an admin role and wants
	// to create a pool for another project.
	ProjectID string `json:"project_id,omitempty"`

	// The protocol - can either be TCP, HTTP or HTTPS.
	Protocol Protocol `json:"protocol" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`

	// A list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`

	// Specifies how wildcard domain name matches with the SNI certificates used by the listener.
	// longest_suffix indicates longest suffix match. wildcard indicates wildcard match.
	// The default value is wildcard.
	SniMatchAlgo string `json:"sni_match_algo,omitempty"`

	// A list of Tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`

	// Specifies the security policy used by the listener.
	TlsCiphersPolicy string `json:"tls_ciphers_policy,omitempty"`

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
	IpGroup *IpGroup `json:"ipgroup,omitempty"`

	// The http insert headers of the Listener.
	InsertHeaders *InsertHeaders `json:"insert_headers,omitempty"`

	// Transparent client ip enable
	TransparentClientIP *bool `json:"transparent_client_ip_enable,omitempty"`

	// Enhance L7policy enable
	EnhanceL7policy *bool `json:"enhance_l7policy_enable,omitempty"`
}

type IpGroup struct {
	IpGroupID string `json:"ipgroup_id" required:"true"`
	Enable    *bool  `json:"enable_ipgroup,omitempty"`
	Type      string `json:"type,omitempty"`
}

type InsertHeaders struct {
	ForwardedELBIP   *bool `json:"X-Forwarded-ELB-IP,omitempty"`
	ForwardedPort    *bool `json:"X-Forwarded-Port,omitempty"`
	ForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`
	ForwardedHost    *bool `json:"X-Forwarded-Host" required:"true"`
}

// ToListenerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "listener")
}

// Create is an operation which provisions a new Listeners based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Listeners on behalf of other tenants by
// specifying a TenantID attribute different from their own.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("listeners"), b, &r.Body, nil)
	return
}

// Get retrieves a particular Listeners based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("listeners", id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToListenerUpdateMap() (map[string]interface{}, error)
}

type IpGroupUpdate struct {
	IpGroupId string `json:"ipgroup_id,omitempty"`
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
}

// ToListenerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToListenerUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "listener")
}

// Update is an operation which modifies the attributes of the specified
// Listener.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToListenerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("listeners", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete will permanently delete a particular Listeners based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("listeners", id), nil)
	return
}

type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

type ListOpts struct {
	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	PageReverse bool   `q:"page_reverse"`

	ProtocolPort            []int      `q:"protocol_port"`
	Protocol                []Protocol `q:"protocol"`
	Description             []string   `q:"description"`
	DefaultTLSContainerRef  []string   `q:"default_tls_container_ref"`
	ClientCATLSContainerRef []string   `q:"client_ca_tls_container_ref"`
	DefaultPoolID           []string   `q:"default_pool_id"`
	ID                      []string   `q:"id"`
	Name                    []string   `q:"name"`
	LoadBalancerID          []string   `q:"loadbalancer_id"`
	TLSCiphersPolicy        []string   `q:"tls_ciphers_policy"`
	MemberAddress           []string   `q:"member_address"`
	MemberDeviceID          []string   `q:"member_device_id"`
	MemberTimeout           []int      `q:"member_timeout"`
	ClientTimeout           []int      `q:"client_timeout"`
	KeepAliveTimeout        []int      `q:"keepalive_timeout"`
}

func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("listeners")
	if opts != nil {
		q, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
