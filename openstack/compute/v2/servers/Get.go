package servers

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get requests details on a single server, by ID.
func Get(client *golangsdk.ServiceClient, id string) (*Server, error) {
	raw, err := get(client, id)
	return ExtractSer(err, raw)
}

func GetInto(client *golangsdk.ServiceClient, id string, v interface{}) (err error) {
	raw, err := get(client, id)
	if err != nil {
		return
	}

	err = extract.IntoStructPtr(raw.Body, v, "server")
	return
}

func get(client *golangsdk.ServiceClient, id string) (*http.Response, error) {
	raw, err := client.Get(client.ServiceURL("servers", id), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return raw, err
}
