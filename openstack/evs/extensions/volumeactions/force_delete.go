package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func ForceDelete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Post(client.ServiceURL("volumes", id, "action"), map[string]any{"os-force_delete": ""}, nil, nil)
	return
}
