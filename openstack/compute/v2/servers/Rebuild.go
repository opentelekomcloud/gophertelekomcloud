package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Rebuild will reprovision the server according to the configuration options provided in the RebuildOpts struct.
func Rebuild(client *golangsdk.ServiceClient, id string, opts RebuildOpts) (*Server, error) {
	b, err := opts.ToServerRebuildMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return ExtractSer(err, raw)
}
