package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ShowConsoleOutputOpts satisfies the ShowConsoleOutputOptsBuilder
type ShowConsoleOutputOpts struct {
	// The number of lines to fetch from the end of console log.
	// All lines will be returned if this is not specified.
	Length int `json:"length,omitempty"`
}

// ShowConsoleOutput makes a request against the nova API to get console log from the server
func ShowConsoleOutput(client *golangsdk.ServiceClient, id string, opts ShowConsoleOutputOpts) (string, error) {
	b, err := build.RequestBody(opts, "os-getConsoleOutput")
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
		Output string `json:"output"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Output, err
}
