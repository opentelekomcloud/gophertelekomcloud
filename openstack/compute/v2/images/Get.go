package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get returns data about a specific image by its ID.
func Get(client *golangsdk.ServiceClient, id string) (*Image, error) {
	raw, err := client.Get(client.ServiceURL("images", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Image
	err = extract.IntoStructPtr(raw.Body, &res, "image")
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
	// Metadata provides free-form key/value pairs that further describe the image.
	Metadata map[string]interface{}
}
