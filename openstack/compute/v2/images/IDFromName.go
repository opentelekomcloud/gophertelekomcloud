package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// IDFromName is a convenience function that returns an image's ID given its name.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	allPages, err := ListDetail(client, ListOpts{}).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractImages(allPages)
	if err != nil {
		return "", err
	}

	for _, f := range all {
		if f.Name == name {
			count++
			id = f.ID
		}
	}

	switch count {
	case 0:
		err := &golangsdk.ErrResourceNotFound{}
		err.ResourceType = "image"
		err.Name = name
		return "", err
	case 1:
		return id, nil
	default:
		err := &golangsdk.ErrMultipleResourcesFound{}
		err.ResourceType = "image"
		err.Name = name
		err.Count = count
		return "", err
	}
}
