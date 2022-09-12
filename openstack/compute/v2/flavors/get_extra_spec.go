package flavors

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetExtraSpec(client *golangsdk.ServiceClient, flavorID string, key string) (map[string]string, error) {
	raw, err := client.Get(client.ServiceURL("flavors", flavorID, "os-extra_specs", key), nil, nil)
	return extraSpe(err, raw)
}

func extraSpe(err error, raw *http.Response) (map[string]string, error) {
	if err != nil {
		return nil, err
	}

	var res map[string]string
	err = extract.Into(raw.Body, &res)
	return res, err
}
