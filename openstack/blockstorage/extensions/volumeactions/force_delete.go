package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func ForceDelete(client *golangsdk.ServiceClient, id string) (r ForceDeleteResult) {
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), map[string]interface{}{"os-force_delete": ""}, nil, nil)
	return
}
