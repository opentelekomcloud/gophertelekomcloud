package domains

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSharedReposOpts struct {
	// Image repository name.
	Name string `q:"name,omitempty"`
	// self: images shared by you.
	// thirdparty: images shared with you by others.
	Center string `q:"center,omitempty"`
	// Number of returned records. Ensure that the offset and limit parameters are used together.
	Limit int `q:"limit,omitempty"`
	// Start index. Ensure that the offset and limit parameters are used together.
	Offset int `q:"offset,omitempty"`
}

func ListSharedReposDetails(client *golangsdk.ServiceClient, opts ListSharedReposOpts) ([]AccessDomain, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v2/manage/shared-repositories
	raw, err := client.Get(client.ServiceURL("manage", "shared-repositories")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AccessDomain
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return res, err
}
