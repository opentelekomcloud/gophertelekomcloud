package images

import (
	"bytes"
	"net/url"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/utils"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List implements image list request.
func List(c *golangsdk.ServiceClient, opts images.ListImagesOpts) pagination.Pager {
	q, err := build.QueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v2/images
	return pagination.Pager{
		Client:     c,
		InitialURL: c.ServiceURL("images") + q.String(),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ImagePage{
				serviceURL:     c.ServiceURL(),
				LinkedPageBase: pagination.LinkedPageBase{PageResult: r},
			}
		},
	}
}

// ImagePage represents the results of a List request.
type ImagePage struct {
	serviceURL string
	pagination.LinkedPageBase
}

// IsEmpty returns true if an ImagePage contains no Images results.
func (r ImagePage) IsEmpty() (bool, error) {
	i, err := ExtractImages(r)
	return len(i) == 0, err
}

// ExtractImages interprets the results of a single page from a List() call,
// producing a slice of Image entities.
func ExtractImages(r pagination.Page) ([]images.ImageInfo, error) {
	var s struct {
		Images []images.ImageInfo `json:"images"`
	}

	err := extract.Into(bytes.NewReader(r.(ImagePage).Body), &s)
	return s.Images, err
}

// NextPageURL uses the response's embedded link reference to navigate to
// the next page of results.
func (r ImagePage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}

	err := extract.Into(bytes.NewReader(r.Body), &s)
	if err != nil {
		return "", err
	}

	if s.Next == "" {
		return "", nil
	}

	return nextPageURL(r.serviceURL, s.Next)
}

func nextPageURL(serviceURL, requestedNext string) (string, error) {
	base, err := utils.BaseEndpoint(serviceURL)
	if err != nil {
		return "", err
	}

	requestedNextURL, err := url.Parse(requestedNext)
	if err != nil {
		return "", err
	}

	base = golangsdk.NormalizeURL(base)
	nextPath := base + strings.TrimPrefix(requestedNextURL.Path, "/")

	nextURL, err := url.Parse(nextPath)
	if err != nil {
		return "", err
	}

	nextURL.RawQuery = requestedNextURL.RawQuery

	return nextURL.String(), nil
}
