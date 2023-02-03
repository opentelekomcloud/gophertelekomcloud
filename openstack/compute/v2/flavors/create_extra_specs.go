package flavors

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateExtraSpecs will create or update the extraAcc-specs key-value pairs for the specified Flavor.
func CreateExtraSpecs(client *golangsdk.ServiceClient, flavorID string, opts map[string]string) (map[string]string, error) {
	raw, err := client.Post(client.ServiceURL("flavors", flavorID, "os-extra_specs"),
		map[string]interface{}{"extra_specs": opts}, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	return extraSpes(err, raw)
}

func extraSpes(err error, raw *http.Response) (map[string]string, error) {
	if err != nil {
		return nil, err
	}

	var res struct {
		ExtraSpecs map[string]string `json:"extra_specs"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ExtraSpecs, err
}
