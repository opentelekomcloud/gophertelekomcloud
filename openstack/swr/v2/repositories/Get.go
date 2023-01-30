package repositories

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, organization, repository string) (*ImageRepository, error) {
	// GET /v2/manage/namespaces/{namespace}/repos/{repository}
	raw, err := client.Get(client.ServiceURL("manage", "namespaces", organization, "repos", repository), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ImageRepository
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return &res, err
}

type ImageRepository struct {
	// Image repository ID.
	ID int `json:"id"`
	// Organization ID.
	OrganizationID int `json:"ns_id"`
	// Image repository name.
	Name string `json:"name"`
	// Image repository type. The value can be app_server, linux, framework_app, database, lang, other, windows, or arm.
	Category string `json:"category"`
	// Brief description of the image repository.
	Description string `json:"description"`
	// Image repository creator ID.
	CreatorID string `json:"creator_id"`
	// Image repository creator.
	CreatorName string `json:"creator_name"`
	// Image repository size.
	Size int `json:"size"`
	// Whether the image repository is a public repository. The value can be true or false.
	IsPublic bool `json:"is_public"`
	// Number of images in an image repository.
	NumImages int `json:"num_images"`
	// Download times of an image repository.
	NumDownloads int `json:"num_downloads"`
	// URL of the image repository logo image. This field has been discarded and is left empty by default.
	URL string `json:"url"`
	// External image pull address. The format is {Repository address}/{Namespace name}/{Repository name}.
	Path string `json:"path"`
	// Internal image pull address. The format is {Repository address}/{Namespace name}/{Repository name}.
	InternalPath string `json:"internal_path"`
	// Time when an image repository is created. It is the UTC standard time.
	Created string `json:"created"`
	// Time when an image repository is updated. It is the UTC standard time.
	Updated string `json:"updated"`
	// Account ID.
	DomainID string `json:"domain_id"`
	// Image sorting priority.
	Priority int `json:"priority"`
}
