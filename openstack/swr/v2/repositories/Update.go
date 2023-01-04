package repositories

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Organization name.
	Namespace string `json:"-"`
	// Image repository name.
	Repository string `json:"-"`
	// Repository type.
	// The value can be app_server, linux, framework_app, database, lang, other, windows or arm.
	Category string `json:"category,omitempty"`
	// Repository description.
	Description string `json:"description,omitempty"`
	// Whether the repository is a public repository. The value can be either true or false.
	IsPublic bool `json:"is_public"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PATCH /v2/manage/namespaces/{namespace}/repos/{repository}
	_, err = client.Patch(client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
