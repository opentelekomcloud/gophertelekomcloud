package cloudservers

import (
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Name of the ECS instance.
	// The value consists of 1 to 128 characters, including letters, digits,
	// underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`

	// Description of the ECS instance.
	// The value consists of 0 to 85 characters. Angle brackets (<>) are not allowed.
	Description *string `json:"description,omitempty"`

	// Hostname of the ECS instance.
	// The value consists of 1 to 64 characters, including letters, digits, hyphens (-), and periods (.).
	// A hostname cannot start or end with a period (.) or hyphen (-).
	// A hostname cannot contain consecutive periods (..) or hyphens (--).
	Hostname string `json:"hostname,omitempty"`

	// SecurityOptions specifies the security options of the ECS, e.g. vTPM.
	SecurityOptions *SecurityOptions `json:"security_options,omitempty"`
}

// UpdateResponse represents the response from the Update ECS API.
// Unlike CloudServer, the image field is returned as a string (empty when booted from volume).
type UpdateResponse struct {
	Status          string                     `json:"status"`
	Updated         time.Time                  `json:"updated"`
	HostID          string                     `json:"hostId"`
	Addresses       map[string][]UpdateAddress `json:"addresses"`
	ID              string                     `json:"id"`
	Name            string                     `json:"name"`
	AccessIPv4      string                     `json:"accessIPv4"`
	AccessIPv6      string                     `json:"accessIPv6"`
	Created         time.Time                  `json:"created"`
	Description     string                     `json:"description"`
	TenantID        string                     `json:"tenant_id"`
	UserID          string                     `json:"user_id"`
	Flavor          Flavor                     `json:"flavor"`
	Metadata        Metadata                   `json:"metadata"`
	Image           string                     `json:"image"`
	Progress        int                        `json:"progress"`
	DiskConfig      string                     `json:"OS-DCF:diskConfig"`
	Hostname        string                     `json:"OS-EXT-SRV-ATTR:hostname"`
	UserData        string                     `json:"OS-EXT-SRV-ATTR:user_data"`
	Links           []Link                     `json:"links"`
	SecurityOptions *SecurityOptions           `json:"security_options"`
}

type UpdateAddress struct {
	Version int    `json:"version"`
	Addr    string `json:"addr"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func Update(client *golangsdk.ServiceClient, serverID string, opts UpdateOpts) (*UpdateResponse, error) {
	b, err := build.RequestBody(opts, "server")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("cloudservers", serverID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		Server *UpdateResponse `json:"server"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Server, err
}
