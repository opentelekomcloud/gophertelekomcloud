package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// GetResult is the response from a Get operation. Call its Extract method to
// interpret it as an Image.
type GetResult struct {
	golangsdk.Result
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// Extract interprets a GetResult as an Image.
func (raw GetResult) Extract() (*Image, error) {
	var res struct {
		Image *Image `json:"image"`
	}
	err = extract.Into(raw, &res)
	return &res, err
}

// Image represents an Image returned by the Compute API.
type Image struct {
	// ID is the unique ID of an image.
	ID string

	// Created is the date when the image was created.
	Created string

	// MinDisk is the minimum amount of disk a flavor must have to be able
	// to create a server based on the image, measured in GB.
	MinDisk int

	// MinRAM is the minimum amount of RAM a flavor must have to be able
	// to create a server based on the image, measured in MB.
	MinRAM int

	// Name provides a human-readable moniker for the OS image.
	Name string

	// The Progress and Status fields indicate image-creation status.
	Progress int

	// Status is the current status of the image.
	Status string

	// Update is the date when the image was updated.
	Updated string

	// Metadata provides free-form key/value pairs that further describe the
	// image.
	Metadata map[string]interface{}
}

// ImagePage contains a single page of all Images returne from a ListDetail
// operation. Use ExtractImages to convert it into a slice of usable structs.
type ImagePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if an ImagePage contains no Image results.
func (page ImagePage) IsEmpty() (bool, error) {
	images, err := ExtractImages(page)
	return len(images) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page ImagePage) NextPageURL() (string, error) {
	var res struct {
		Links []golangsdk.Link `json:"images_links"`
	}
	err = extract.Into(page.Result.Body, &res)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res.Links)
}

// ExtractImages converts a page of List results into a slice of usable Image
// structs.
func ExtractImages(r pagination.Page) ([]Image, error) {
	var res struct {
		Images []Image `json:"images"`
	}
	err := (r.(ImagePage)).ExtractInto(&res)
	return res, err
}
