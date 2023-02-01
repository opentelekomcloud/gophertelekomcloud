package members

type BatchDeleteMembersOpts struct {
	// Specifies the image IDs.
	Images []string `json:"images"`
	// Specifies the project IDs.
	Projects []string `json:"projects"`
}

// DELETE /v1/cloudimages/members

// 200 Job ID
