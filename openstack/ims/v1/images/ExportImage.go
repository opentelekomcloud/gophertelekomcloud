package images

type ExportImageOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the URL of the image file in the format of Bucket name:File name.
	//
	// Note
	//
	// The storage class of the OBS bucket must be Standard.
	BucketUrl string `json:"bucket_url" required:"true"`
	// Specifies the file format. The value can be qcow2, vhd, zvhd, or vmdk.
	FileFormat string `json:"file_format" required:"true"`
	// Whether to enable fast export. The value can be true or false.
	//
	// Note
	//
	// If fast export is enabled, file_format cannot be specified.
	IsQuickExport *bool `json:"is_quick_export,omitempty"`
}

// POST /v1/cloudimages/{image_id}/file

// 200 JobId
