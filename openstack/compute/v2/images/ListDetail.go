package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts contain options filtering Images returned from a call to ListDetail.
type ListOpts struct {
	// ChangesSince filters Images based on the last changed status (in date-time format).
	ChangesSince string `q:"changes-since"`
	// Limit limits the number of Images to return.
	Limit int `q:"limit"`
	// Mark is an Image UUID at which to set a marker.
	Marker string `q:"marker"`
	// Name is the name of the Image.
	Name string `q:"name"`
	// Server is the name of the Server (in URL format).
	Server string `q:"server"`
	// Status is the current status of the Image.
	Status string `q:"status"`
	// Type is the type of image (e.g. BASE, SERVER, ALL).
	Type string `q:"type"`
}

// ListDetail enumerates the available images.
func ListDetail(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("images", "detail")+query.String(),
		func(r pagination.PageResult) pagination.Page {
			return ImagePage{pagination.LinkedPageBase{PageResult: r}}
		})
}

// ImagePage contains a single page of all Images return from a ListDetail
// operation. Use ExtractImages to convert it into a slice of usable structs.
type ImagePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if an ImagePage contains no Image results.
func (page ImagePage) IsEmpty() (bool, error) {
	images, err := ExtractImages(page)
	return len(images) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page ImagePage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(page.BodyReader(), &res, "images_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// ExtractImages converts a page of List results into a slice of usable Image structs.
func ExtractImages(r pagination.Page) ([]Image, error) {
	var res []Image
	err := extract.IntoSlicePtr(r.(ImagePage).BodyReader(), &res, "images")
	return res, err
}
