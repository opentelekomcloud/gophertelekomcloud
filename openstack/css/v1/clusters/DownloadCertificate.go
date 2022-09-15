package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

func DownloadCertificate(client *golangsdk.ServiceClient, clusterID string) (r CertificateResult) {
	raw, err = client.Get(client.ServiceURL("clusters", clusterID, "sslCert"), nil, nil)
	return
}
