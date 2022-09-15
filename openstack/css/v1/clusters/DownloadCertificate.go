package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func DownloadCertificate(client *golangsdk.ServiceClient, clusterID string) (string, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "sslCert"), nil, nil)
	if err != nil {
		return "", err
	}

	var res struct {
		CertBase64 string `json:"certBase64"`
	}
	err = extract.Into(raw.Body, &res)
	return res.CertBase64, err
}
