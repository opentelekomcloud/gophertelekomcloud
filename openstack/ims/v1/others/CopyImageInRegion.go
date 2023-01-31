package others

type CopyImageInRegionOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the image name.
	Name string `json:"name" required:"true"`
	// Specifies the encryption key. This parameter is left blank by default.
	CmkId string `json:"cmk_id,omitempty"`
	// Provides supplementary information about the image. For detailed description, see Image Attributes. The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
}

// POST /v1/cloudimages/{image_id}/copy

// 200 job id
