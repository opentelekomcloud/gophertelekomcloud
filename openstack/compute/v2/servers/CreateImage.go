package servers

import (
	"fmt"
	"net/url"
	"path"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateImageOpts provides options to pass to the CreateImage request.
type CreateImageOpts struct {
	// Name of the image/snapshot.
	Name string `json:"name" required:"true"`
	// Metadata contains key-value pairs (up to 255 bytes each) to attach to the created image.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// CreateImage makes a request against the nova API to schedule an image to be created of the server
func CreateImage(client *golangsdk.ServiceClient, id string, opts CreateImageOpts) (string, error) {
	b, err := build.RequestBody(opts, "createImage")
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
