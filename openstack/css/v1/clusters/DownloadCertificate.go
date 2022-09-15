package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

func DownloadCertificate(client *golangsdk.ServiceClient, clusterID string) (r CertificateResult) {
	_, r.Err = client.Get(client.ServiceURL("clusters", clusterID, "sslCert"), &r.Body, nil)
	return
}
