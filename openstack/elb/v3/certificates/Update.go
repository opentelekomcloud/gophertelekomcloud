package certificates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Specifies the certificate name.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the certificate.
	//
	// Minimum: 0
	//
	// Maximum: 255
	Description string `json:"description,omitempty"`
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
	// Minimum: 0
	//
	// Maximum: 1024
	Domain string `json:"domain,omitempty"`
	// Specifies the private key of the server certificate. The value must be PEM encoded.
	//
	// This parameter will be ignored if type is set to client. A CA server can still be created and used normally. This parameter will be left blank even if you enter a private key that is not PEM encoded.
	//
	// This parameter is valid and mandatory only when type is set to server. If you enter an invalid private key, an error is returned.
	PrivateKey string `json:"private_key,omitempty"`
	// Specifies the private key of the certificate. The value must be PEM encoded.
	Certificate string `json:"certificate,omitempty"`
}

// Update is an operation which modifies the attributes of the specified Certificate.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Certificate, error) {
	b, err := build.RequestBody(opts, "certificate")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("certificates", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
