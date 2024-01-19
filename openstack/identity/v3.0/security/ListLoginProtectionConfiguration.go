package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListLoginProtectionConfigurations(client *golangsdk.ServiceClient) ([]LoginProtectionConfig, error) {
	// GET  /v3.0/OS-USER/login-protects
	raw, err := client.Get(client.ServiceURL("OS-USER", "login-protects"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []LoginProtectionConfig
	err = extract.IntoSlicePtr(raw.Body, &res, "login_protects")
	return res, err
}
