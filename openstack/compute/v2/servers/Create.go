package servers

import (
	"encoding/base64"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/flavors"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/images"
)

// CreateOpts specifies server creation parameters.
type CreateOpts struct {
	// Name is the name to assign to the newly launched server.
	Name string `json:"name" required:"true"`
	// ImageRef [optional; required if ImageName is not provided] is the ID or
	// full URL to the image that contains the server's OS and initial state.
	// Also, optional if using the boot-from-volume extension.
	ImageRef string `json:"imageRef"`
	// ImageName [optional; required if ImageRef is not provided] is the name of
	// the image that contains the server's OS and initial state.
	// Also, optional if using the boot-from-volume extension.
	ImageName string `json:"-"`
	// FlavorRef [optional; required if FlavorName is not provided] is the ID or
	// full URL to the flavor that describes the server's specs.
	FlavorRef string `json:"flavorRef"`
	// FlavorName [optional; required if FlavorRef is not provided] is the name of
	// the flavor that describes the server's specs.
	FlavorName string `json:"-"`
	// SecurityGroups lists the names of the security groups to which this server should belong.
	SecurityGroups []string `json:"-"`
	// UserData contains configuration information or scripts to use upon launch.
	// Create will base64-encode it for you, if it isn't already.
	UserData []byte `json:"-"`
	// AvailabilityZone in which to launch the server.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// Networks dictates how this server will be attached to available networks.
	// By default, the server will be attached to all isolated networks for the tenant.
	Networks []Network `json:"-"`
	// Metadata contains key-value pairs (up to 255 bytes each) to attach to the server.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Personality includes files to inject into the server at launch.
	// Create will base64-encode file contents for you.
	Personality []*File `json:"personality,omitempty"`
	// ConfigDrive enables metadata injection through a configuration drive.
	ConfigDrive *bool `json:"config_drive,omitempty"`
	// AdminPass sets the root user password. If not set, a randomly-generated
	// password will be created and returned to the response.
	AdminPass string `json:"adminPass,omitempty"`
	// AccessIPv4 specifies an IPv4 address for the instance.
	AccessIPv4 string `json:"accessIPv4,omitempty"`
	// AccessIPv6 specifies an IPv6 address for the instance.
	AccessIPv6 string `json:"accessIPv6,omitempty"`
	// ServiceClient will allow calls to be made to retrieve an image or flavor ID by name.
	ServiceClient *golangsdk.ServiceClient `json:"-"`
}

// Network is used within CreateOpts to control a new server's network attachments.
type Network struct {
	// UUID of a network to attach to the newly provisioned server.
	// Required unless Port is provided.
	UUID string
	// Port of a neutron network to attach to the newly provisioned server.
	// Required unless UUID is provided.
	Port string
	// FixedIP specifies a fixed IPv4 address to be used on this network.
	FixedIP string
}

// CreateOptsBuilder CreateOptsWithCustomField
type CreateOptsBuilder interface {
	ToServerCreateMap() (map[string]interface{}, error)
}

// ToServerCreateMap assembles a request body based on the contents of a CreateOpts.
func (opts CreateOpts) ToServerCreateMap() (map[string]interface{}, error) {
	sc := opts.ServiceClient
	opts.ServiceClient = nil
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.UserData != nil {
		var userData string
		if _, err := base64.StdEncoding.DecodeString(string(opts.UserData)); err != nil {
			userData = base64.StdEncoding.EncodeToString(opts.UserData)
		} else {
			userData = string(opts.UserData)
		}
		b["user_data"] = &userData
	}

	if len(opts.SecurityGroups) > 0 {
		securityGroups := make([]map[string]interface{}, len(opts.SecurityGroups))
		for i, groupName := range opts.SecurityGroups {
			securityGroups[i] = map[string]interface{}{"name": groupName}
		}
		b["security_groups"] = securityGroups
	}

	if len(opts.Networks) > 0 {
		networks := make([]map[string]interface{}, len(opts.Networks))
		for i, net := range opts.Networks {
			networks[i] = make(map[string]interface{})
			if net.UUID != "" {
				networks[i]["uuid"] = net.UUID
			}
			if net.Port != "" {
				networks[i]["port"] = net.Port
			}
			if net.FixedIP != "" {
				networks[i]["fixed_ip"] = net.FixedIP
			}
		}
		b["networks"] = networks
	}

	// If ImageRef isn't provided, check if ImageName was provided to ascertain
	// the image ID.
	if opts.ImageRef == "" {
		if opts.ImageName != "" {
			if sc == nil {
				err := ErrNoClientProvidedForIDByName{}
				err.Argument = "ServiceClient"
				return nil, err
			}
			imageID, err := images.IDFromName(sc, opts.ImageName)
			if err != nil {
				return nil, err
			}
			b["imageRef"] = imageID
		}
	}

	// If FlavorRef isn't provided, use FlavorName to ascertain the flavor ID.
	if opts.FlavorRef == "" {
		if opts.FlavorName == "" {
			err := ErrNeitherFlavorIDNorFlavorNameProvided{}
			err.Argument = "FlavorRef/FlavorName"
			return nil, err
		}
		if sc == nil {
			err := ErrNoClientProvidedForIDByName{}
			err.Argument = "ServiceClient"
			return nil, err
		}
		flavorID, err := flavors.IDFromName(sc, opts.FlavorName)
		if err != nil {
			return nil, err
		}
		b["flavorRef"] = flavorID
	}

	return map[string]interface{}{"server": b}, nil
}

// Create requests a server to be provisioned to the user in the current tenant.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (*Server, error) {
	b, err := opts.ToServerCreateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("servers"), b, nil, nil)
	return ExtractSer(err, raw)
}
