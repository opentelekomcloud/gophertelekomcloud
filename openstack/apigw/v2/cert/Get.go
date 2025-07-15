package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves details of a specific SSL certificate
func Get(client *golangsdk.ServiceClient, certificateID string) (*CertificateResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "certificates", certificateID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res CertificateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
