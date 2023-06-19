package metadata

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Create(client *golangsdk.ServiceClient, volumeId string, opts map[string]string) (map[string]string, error) {
	b, err := build.RequestBody(opts, "metadata")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/volumes/{volume_id}/metadata
	raw, err := client.Post(client.ServiceURL("volumes", volumeId, "metadata"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraMetadata(err, raw)
}

func extraMetadata(err error, raw *http.Response) (map[string]string, error) {
	if err != nil {
		return nil, err
	}

	var res struct {
		Metadata map[string]string `json:"metadata"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Metadata, err
}
