package ssl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func DownloadCert(client *golangsdk.ServiceClient, id string) (*DownloadCertResp, error) {
	raw, err := client.Post(client.ServiceURL("instances", id, "ssl-certs", "download"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}
	var res DownloadCertResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DownloadCertResp struct {
	FileName   string `json:"file_name"`
	Link       string `json:"link"`
	BucketName string `json:"bucket_name"`
}
