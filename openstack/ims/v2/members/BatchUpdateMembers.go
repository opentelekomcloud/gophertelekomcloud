package members

type BatchUpdateMembersOpts struct {
	// Specifies the image IDs.
	Images []string `json:"images" required:"true"`
	// Specifies the project ID.
	ProjectId string `json:"project_id" required:"true"`
	// Specifies whether a shared image will be accepted or declined.
	//
	// The value can be one of the following:
	//
	// accepted: indicates that a shared image is accepted. After an image is accepted, the image is displayed in the image list. You can use the image to create ECSs.
	//
	// rejected: indicates that a shared image is declined. After an image is declined, the image is not displayed in the image list. However, you can still use the image to create ECSs.
	Status string `json:"status" required:"true"`
	// Specifies the ID of a vault.
	//
	// This parameter is mandatory if you want to accept a shared full-ECS image created from a CBR backup.
	//
	// You can obtain the vault ID from the CBR console or section "Querying the Vault List" in Cloud Backup and Recovery API Reference.
	VaultId string `json:"vault_id,omitempty"`
}

// PUT /v1/cloudimages/members

// 200 Job ID
