package domain

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type CertificateOpts struct {
	GatewayID     string
	GroupID       string
	DomainID      string
	CertificateID string
}

func UnbindCertificate(client *golangsdk.ServiceClient, opts CertificateOpts) (err error) {
	// DELETE /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificate/{certificate_id}
	_, err = client.Delete(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "domains", opts.DomainID, "certificate", opts.CertificateID), nil)
	return
}
