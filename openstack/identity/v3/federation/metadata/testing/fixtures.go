package testing

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
)

const (
	domainID    = "ed7a77d365304f458f7d0a7909c6d889"
	metadataURL = "/OS-FEDERATION/identity_providers/ACME/protocols/saml/metadata"
)

var (
	data = tools.RandomString("meta-", 500)

	getMetadataResponseBody = fmt.Sprintf(`
{
  "id": "40c174f35ff94e31b8257ad4991bce8b",
  "idp_id": "ACME",
  "entity_id": "https://idp.test.com/idp/shibboleth",
  "protocol_id": "saml",
  "domain_id": "%s",
  "xaccount_type": "",
  "update_time": "2016-10-26T09:26:23.000000",
  "data": "%s"
}
`, domainID, data)
	importMetadataBody = fmt.Sprintf(
		`
{
  "xaccount_type": "",
  "domain_id": "%s",
  "metadata": "%s"
}
`, domainID, data)
)
