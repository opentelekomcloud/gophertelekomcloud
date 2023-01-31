package tags

type DeleteImageTagOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-"`
	// Specifies the key of the tag to be deleted.
	Key string `json:"-"`
}

// DELETE /v2/{project_id}/images/{image_id}/tags/{key}

// 204
