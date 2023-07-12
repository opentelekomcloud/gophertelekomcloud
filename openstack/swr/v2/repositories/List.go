package repositories

import (
	"bytes"
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// The value can only be self, indicating that the image is a self-owned image.
	Center string `q:"center"`
	// Organization (namespace) name
	Organization string `q:"namespace"`
	// Image repository name
	Name string `q:"name"`
	// Image repository type. The value can be app_server, linux, framework_app, database, lang, other, windows, or arm.
	Category string `q:"category"`
	// Sorting criteria. The value can be name, updated_time, or tag_count.
	// Ensure that the order_column and order_type parameters are used together
	OrderColumn string `q:"order_column"`
	// Sorting type. You can set this parameter to desc (descending sort) and asc (ascending sort).
	// Ensure that the order_column and order_type parameters are used together.
	OrderType string `q:"order_type"`
	// Start index.
	// Ensure that the offset and limit parameters are used together.
	Offset *int `q:"offset,omitempty"`
	// Number of returned records.
	// Ensure that the offset and limit parameters are used together.
	Limit int `q:"limit,omitempty"`
}

func (opts ListOpts) ToRepositoryListQuery() (string, error) {
	if opts.Limit == 0 && opts.Offset != nil {
		opts.Limit = 25
	}

	if opts.Limit != 0 && opts.Offset == nil {
		return "", fmt.Errorf("offset has to be defined if the limit is set")
	}

	if (opts.OrderColumn != "" && opts.OrderType == "") || (opts.OrderColumn == "" && opts.OrderType != "") {
		return "", fmt.Errorf("`OrderColumn` and `OrderType` should always be used together")
	}

	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}

	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (p pagination.Pager) {
	q, err := opts.ToRepositoryListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v2/manage/repos
	url := client.ServiceURL("manage", "repos") + q
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return RepositoryPage{pagination.OffsetPageBase{PageResult: r}}
		},
	}
}

type RepositoryPage struct {
	pagination.OffsetPageBase
}

func ExtractRepositories(p pagination.Page) ([]ImageRepository, error) {
	var res []ImageRepository
	err := extract.IntoSlicePtr(bytes.NewReader(p.(RepositoryPage).Body), &res, "")
	return res, err
}
