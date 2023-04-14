package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves a particular Load balancer based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Certificate, error) {
	raw, err := client.Get(client.ServiceURL("certificates", id), nil, nil)
	return extra(err, raw)
}

type Certificate struct {
	// Specifies a certificate ID
	ID string `json:"id"`
	// Specifies the project ID.
	ProjectID string `json:"project_id"`
	// Specifies the certificate name.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Name string `json:"name"`
	// Provides supplementary information about the certificate.
	//
	// Minimum: 1
	//
	// Maximum: 255
	Description string `json:"description"`
	// Specifies the certificate type. The value can be server or client. server indicates server certificates, and client indicates CA certificates. The default value is server.
	Type string `json:"type"`
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
	//
	// Minimum: 1
	//
	// Maximum: 1024
	Domain string `json:"domain"`
	// Specifies the private key of the server certificate. The value must be PEM encoded.
	//
	// This parameter will be ignored if type is set to client. A CA server can still be created and used normally. This parameter will be left blank even if you enter a private key that is not PEM encoded.
	//
	// This parameter is valid and mandatory only when type is set to server. If you enter an invalid private key, an error is returned.
	PrivateKey string `json:"private_key"`
	// Specifies the private key of the certificate. The value must be PEM encoded.
	Certificate string `json:"certificate"`
	// Specifies the administrative status of the certificate.
	//
	// This parameter is unsupported. Please do not use it.
	AdminStateUp bool `json:"admin_state_up"`
	// Specifies the time when the certificate was created.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the certificate was updated.
	UpdatedAt string `json:"updated_at"`
	// Specifies the time when the certificate expires.
	ExpireTime string `json:"expire_time"`
}
