package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UnbindCertFromDomain unbinds an SSL certificate from domain names
func UnbindCertFromDomain(client *golangsdk.ServiceClient, opts AttachDomainOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("apigw", "certificates", opts.CertificateID, "domains", "detach"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
