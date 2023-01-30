package repositories

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteByTagOpts struct {
	// Organization name
	Namespace string `json:"-" required:"true"`
	// Image repository name
	Repository string `json:"-" required:"true"`
	// Image tag name
	Tag string `json:"-" required:"true"`
}

func DeleteByTag(client *golangsdk.ServiceClient, opts DeleteByTagOpts) (err error) {
	// DELETE /v2/manage/namespaces/{namespace}/repos/{repository}/tags/{tag}
	url := client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "tags", opts.Tag)
	_, err = client.Delete(url, nil)
	return
}
