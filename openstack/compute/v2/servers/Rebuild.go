package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/images"
)

// RebuildOpts represents the configuration options used in a server rebuild operation.
type RebuildOpts struct {
	// AdminPass is the server's admin password
	AdminPass string `json:"adminPass,omitempty"`
	// ImageID is the ID of the image you want your server to be provisioned on.
	ImageID string `json:"imageRef"`
	// ImageName is readable name of an image.
	ImageName string `json:"-"`
	// Name to set the server to
	Name string `json:"name,omitempty"`
	// AccessIPv4 [optional] provides a new IPv4 address for the instance.
	AccessIPv4 string `json:"accessIPv4,omitempty"`
	// AccessIPv6 [optional] provides a new IPv6 address for the instance.
	AccessIPv6 string `json:"accessIPv6,omitempty"`
	// Metadata [optional] contains key-value pairs (up to 255 bytes each)
	// to attach to the server.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Personality [optional] includes files to inject into the server at launch.
	// Rebuild will base64-encode file contents for you.
	Personality []*File `json:"personality,omitempty"`
	// ServiceClient will allow calls to be made to retrieve an image or
	// flavor ID by name.
	ServiceClient *golangsdk.ServiceClient `json:"-"`
}

// File is used within CreateOpts and RebuildOpts to inject a file into the server at launch.
// File implements the json.Marshaler interface, so when a Create or Rebuild
// operation is requested, json.Marshal will call File's MarshalJSON method.
type File struct {
	// Path of the file.
	Path string
	// Contents of the file. Maximum content size is 255 bytes.
	Contents []byte
}

// ToServerRebuildMap formats a RebuildOpts struct into a map for use in JSON
func (opts RebuildOpts) ToServerRebuildMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// If ImageRef isn't provided, check if ImageName was provided to ascertain
	// the image ID.
	if opts.ImageID == "" {
		if opts.ImageName != "" {
			if opts.ServiceClient == nil {
				err := ErrNoClientProvidedForIDByName{}
				err.Argument = "ServiceClient"
				return nil, err
			}
			imageID, err := images.IDFromName(opts.ServiceClient, opts.ImageName)
			if err != nil {
				return nil, err
			}
			b["imageRef"] = imageID
		}
	}

	return map[string]interface{}{"rebuild": b}, nil
}

// Rebuild will reprovision the server according to the configuration options provided in the RebuildOpts struct.
func Rebuild(client *golangsdk.ServiceClient, id string, opts RebuildOpts) (*Server, error) {
	b, err := opts.ToServerRebuildMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return ExtractSer(err, raw)
}
