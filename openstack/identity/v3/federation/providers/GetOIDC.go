package providers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation"
)

func GetOIDC(c *golangsdk.ServiceClient, idpId string) (*CreateOIDCOpts, error) {
	raw, err := c.Get(c.ServiceURL(federation.BaseURL, "identity-providers", idpId, "openid-connect-config"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}

	var res CreateOIDCOpts
	err = extract.IntoStructPtr(raw.Body, &res, "openid_connect_config")
	return &res, err
}
