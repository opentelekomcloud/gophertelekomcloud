package apiversions

import (
	"net/url"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(c *golangsdk.ServiceClient) pagination.Pager {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"

	return pagination.NewPager(c, u.String(), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}

type APIVersion struct {
	// unique identifier
	ID string `json:"id"`
	// current status
	Status string `json:"status"`
	// date last updated
	Updated string `json:"updated"`
}

type APIVersionPage struct {
	pagination.SinglePageBase
}

func (r APIVersionPage) IsEmpty() (bool, error) {
	is, err := ExtractAPIVersions(r)
	return len(is) == 0, err
}

func ExtractAPIVersions(r pagination.Page) ([]APIVersion, error) {
	var res struct {
		Versions []APIVersion `json:"versions"`
	}
	err := extract.Into(r.(APIVersionPage).Body, &res)
	return res.Versions, err
}
