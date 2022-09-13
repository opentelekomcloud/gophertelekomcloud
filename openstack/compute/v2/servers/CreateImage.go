package servers

import (
	"fmt"
	"net/url"
	"path"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateImage makes a request against the nova API to schedule an image to be created of the server
func CreateImage(client *golangsdk.ServiceClient, id string, opts CreateImageOptsBuilder) (string, error) {
	b, err := opts.ToServerCreateImageMap()
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return "", err
	}

	// Get the image id from the header
	u, err := url.ParseRequestURI(raw.Header.Get("Location"))
	if err != nil {
		return "", err
	}

	imageID := path.Base(u.Path)
	if imageID == "." || imageID == "/" {
		return "", fmt.Errorf("failed to parse the ID of newly created image: %s", u)
	}
	return imageID, nil
}
