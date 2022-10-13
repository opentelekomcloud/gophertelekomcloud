package evacuate

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// EvacuateOpts specifies Evacuate action parameters.
type EvacuateOpts struct {
	// The name of the host to which the server is evacuated
	Host string `json:"host,omitempty"`
	// Indicates whether server is on shared storage
	OnSharedStorage bool `json:"onSharedStorage"`
	// An administrative password to access the evacuated server
	AdminPass string `json:"adminPass,omitempty"`
}

// Evacuate will Evacuate a failed instance to another host.
func Evacuate(client *golangsdk.ServiceClient, id string, opts EvacuateOpts) (string, error) {
	b, err := build.RequestBody(opts, "evacuate")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		AdminPass string `json:"adminPass"`
	}
	err = extract.Into(raw.Body, &res)
	return res.AdminPass, err
}
