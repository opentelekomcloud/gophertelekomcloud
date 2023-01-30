package repositories

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	// Organization name.
	Namespace string `json:"-"`
	// Image repository name.
	// Enter 1 to 128 characters, starting and ending with a lowercase letter or digit. Only lowercase letters, digits, periods (.), slashes (/), underscores (_), and hyphens (-) are allowed. Periods, slashes, underscores, and hyphens cannot be placed next to each other. A maximum of two consecutive underscores are allowed.
	Repository string `json:"repository"`
	// Repository type.
	// The value can be app_server, linux, framework_app, database, lang, other, windows or arm.
	Category string `json:"category,omitempty"`
	// Brief description of the image repository.
	Description string `json:"description,omitempty"`
	// Whether the repository is a public repository. When the value is true, it indicates the repository is public. When the value is false, it indicates the repository is private.
	IsPublic bool `json:"is_public"`
}

// Create new repository in the organization (namespace)
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v2/manage/namespaces/{namespace}/repos
	_, err = client.Post(client.ServiceURL("manage", "namespaces", opts.Namespace, "repos"), b, nil, nil)
	return
}
