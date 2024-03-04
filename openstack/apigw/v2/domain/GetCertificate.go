package domain

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetCertificate(client *golangsdk.ServiceClient, opts CertificateOpts) (*Certificate, error) {
	// GET /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificate/{certificate_id}
	raw, err := client.Get(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "domains", opts.DomainID, "certificate"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Certificate
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Certificate struct {
	// Certificate ID.
	ID string `json:"id"`
	// Certificate name.
	Name string `json:"name"`
	// Certificate type. Options:
	// global: Global certificate.
	// instance: Gateway certificate.
	Type string `json:"type"`
	// Gateway ID.
	// If type is set to global, the default value is common.
	// If type is set to instance, a gateway ID is displayed.
	GatewayId string `json:"instance_id"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// Creation time.
	CreatedAt string `json:"create_time"`
	// Update time.
	UpdatedAt string `json:"update_time"`
	// Certificate domain name.
	CommonName string `json:"common_name"`
	// Subject alternative names.
	San []string `json:"san"`
	// Certificate version.
	Version int `json:"version"`
	// Company or organization.
	Organization []string `json:"organization"`
	// Department.
	OrganizationalUnit []string `json:"organizational_unit"`
	// City.
	City []string `json:"locality"`
	// State or province.
	State []string `json:"state"`
	// Country or region.
	Country []string `json:"country"`
	// Start time of the certificate validity period.
	NotBefore string `json:"not_before"`
	// End time of the certificate validity period.
	NotAfter string `json:"not_after"`
	// Serial No.
	SerialNumber string `json:"serial_number"`
	// Certificate issuer.
	Issuer []string `json:"issuer"`
	// Signature algorithm.
	SignatureAlgorithm string `json:"signature_algorithm"`
}
