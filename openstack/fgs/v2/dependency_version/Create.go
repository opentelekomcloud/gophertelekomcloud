package dependency_version

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	DependFile  string `json:"depend_file,omitempty"`
	DependLink  string `json:"depend_link,omitempty"`
	DependType  string `json:"depend_type" required:"true"`
	Runtime     string `json:"runtime" required:"true"`
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*DepVersionResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "dependencies", "version"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DepVersionResp
	return &res, extract.Into(raw.Body, &res)
}

type DepVersionResp struct {
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	Link         string `json:"link"`
	Runtime      string `json:"runtime"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	FileName     string `json:"file_name"`
	Version      int    `json:"version"`
	DepId        string `json:"dep_id"`
	LastModified int    `json:"last_modified"`
	DownloadLink string `json:"download_link"`
	IsShared     bool   `json:"is_shared"`
}
