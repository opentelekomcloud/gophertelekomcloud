package certificates

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Create operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type CreateOptsBuilder interface {
	ToCertificateCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	// Specifies the certificate name. Only letters, digits, underscores, and hyphens are allowed.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the certificate.
	//
	// Minimum: 0
	// Maximum: 255
	Description string `json:"description,omitempty"`
	// Specifies the certificate type.
	//
	// The value can be server or client. server indicates server certificates, and client indicates CA certificates. The default value is server.
	Type string `json:"type,omitempty"`
	// Specifies the domain names used by the server certificate.
	//
	// This parameter will take effect only when type is set to server, and its default value is "".
	//
	// This parameter will not take effect even if it is passed and type is set to client. However, domain names will still be verified.
	//
	// Note:
	//
	// The value can contain 0 to 1024 characters and consists of multiple common domain names or wildcard domain names separated by commas. A maximum of 30 domain names are allowed.
	//
	// A common domain name consists of several labels separated by periods (.). Each label can contain a maximum of 63 characters, including letters, digits, and hyphens (-), and must start and end with a letter or digit. Example: www.test.com
	//
	// A wildcard domain name is a domain name starts with an asterisk (*). Example: *.test.com
	Domain string `json:"domain,omitempty"`
	//
	// Specifies the private key of the server certificate. The value must be PEM encoded.
	//
	// This parameter will be ignored if type is set to client. A CA server can still be created and used normally. This parameter will be left blank even if you enter a private key that is not PEM encoded.
	//
	// This parameter is valid and mandatory only when type is set to server. If you enter an invalid private key, an error is returned.
	PrivateKey string `json:"private_key,omitempty"`
	// Specifies the private key of the certificate. The value must be PEM encoded.
	Certificate string `json:"certificate" required:"true"`
	// Specifies the administrative status of the certificate.
	//
	// This parameter is unsupported. Please do not use it.
	//
	// Default: true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// Specifies the ID of the project where the certificate is used.
	//
	// Minimum: 1
	// Maximum: 32
	ProjectId string `json:"project_id,omitempty"`
}

// ToCertificateCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToCertificateCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "certificate")
}

// Create is an operation which provisions a new loadbalancer based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create loadbalancers on behalf of other tenants by
// specifying a TenantID attribute different from their own.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCertificateCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	// POST /v3/{project_id}/elb/certificates
	_, r.Err = client.Post(client.ServiceURL("certificates"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}
