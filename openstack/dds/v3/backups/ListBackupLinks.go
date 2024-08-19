package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BackupLinkOpts struct {
	// Instance id.
	InstanceId string `q:"instance_id" required:"true"`
	// Specifies the backup ID.
	BackupId string `q:"backup_id" required:"true"`
}

func ListBackupDownloadLinks(client *golangsdk.ServiceClient, opts BackupLinkOpts) (*LinkResponse, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/download-file
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("backups", "download-file").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res LinkResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type LinkResponse struct {
	Files  []Files `json:"files"`
	Bucket string  `json:"bucket"`
}

type Files struct {
	// Indicates the file name.
	Name string `json:"name"`
	// Indicates the file size in KB.
	Size int64 `json:"size"`
	// Indicates the link for downloading the backup file.
	DownloadLink string `json:"download_link"`
	// Indicates the link expiration time.
	LinkExpiredTime string `json:"link_expired_time"`
}
