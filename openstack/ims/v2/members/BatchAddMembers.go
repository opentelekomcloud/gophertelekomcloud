package members

type BatchAddMembersOpts struct {
	// Specifies the image IDs.
	Images []string `json:"images"`
	// Specifies the project IDs.
	Projects []string `json:"projects"`
}

// POST /v1/cloudimages/members

// 200 Job ID
