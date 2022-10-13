package flavors

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// RemoveAccess removes/revokes a tenant/project access to a flavor.
func RemoveAccess(client *golangsdk.ServiceClient, id string, opts AccessOpts) ([]FlavorAccess, error) {
	b, err := build.RequestBody(opts, "removeTenantAccess")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("flavors", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraAcc(err, raw)
}

func extraAcc(err error, raw *http.Response) ([]FlavorAccess, error) {
	if err != nil {
		return nil, err
	}

	var res []FlavorAccess
	err = extract.IntoSlicePtr(raw.Body, &res, "flavor_access")
	return res, err
}
