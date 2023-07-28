package repositories

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ListTagsOpts struct {
	// Organization name
	Namespace string `json:"-" required:"true"`
	// Image repository name.
	Repository string `json:"-" required:"true"`
	// Start index.
	Offset int `q:"offset"`
	// Number of responses.
	Limit int `q:"limit"`
	// Sorting by column. You can set this parameter to updated_at (sorting by update time).
	OrderColumn string `q:"order_column"`
	// Sorting type. You can set this parameter to desc (descending sort) and asc (ascending sort).
	OrderType string `q:"order_type"`
	// Image tag name.
	Tag string `q:"tag"`
}

func ListTags(client *golangsdk.ServiceClient, opts ListTagsOpts) ([]TagsResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v2/manage/namespaces/{namespace}/repos/{repository}/tags
	url := client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "tags") + q.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []TagsResponse
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return res, err
}

type TagsResponse struct {
	// ID
	Id int64 `json:"id"`
	// Repository ID.
	RepoId int64 `json:"repo_id"`
	// Image tag name.
	Tag string `json:"Tag"`
	// Image ID.
	ImageId string `json:"image_id"`
	// Image manifest.
	Manifest string `json:"manifest"`
	// SHA value of an image.
	Digest string `json:"digest"`
	// Docker protocol version. The version can be 1 or 2.
	Schema int64 `json:"schema"`
	// External image pull address. The format is {Repository address}/{Namespace name}/{Repository name}:{Tag name}.
	Path string `json:"path"`
	// Internal image pull address. The format is {Repository address}/{Namespace name}/{Repository name}:{Tag name}.
	InternalPath string `json:"internal_path"`
	// Image size.
	// Value range: 0 to 9223372036854775807
	// Unit: byte
	Size int64 `json:"size"`
	// By default, the value is false.
	IsTrusted bool `json:"is_trusted"`
	// Time when an image is created. It is the UTC standard time. Users need to calculate the offset based on the local time. For example, GMT+8 is 8 hours ahead the GMT time.
	Created string `json:"created"`
	// Time when an image is updated. It is the UTC standard time. Users need to calculate the offset based on the local time. For example, GMT+8 is 8 hours ahead the GMT time.
	Updated string `json:"updated"`
	// Time when an image was deleted.
	Deleted string `json:"deleted"`
	// Account ID.
	DomainId string `json:"domain_id"`
	// 0: manifest. 1: manifest list.
	TagType int64 `json:"tag_type"`
}
