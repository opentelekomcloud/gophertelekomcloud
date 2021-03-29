package testing

import (
	"fmt"
	"strings"
)

const (
	providerID          = "ACME"
	providerDescription = "Stores ACME identities"

	listURI = "/OS-FEDERATION/identity_providers"
)

var (
	providerURI = fmt.Sprintf("%s/%s", listURI, providerID)

	providerResponse = fmt.Sprintf(`
{
    "identity_provider": {
        "description": "%s",
        "enabled": true,
        "id": "%s",
        "remote_ids": [],
        "links": {
            "protocols": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME/protocols",
            "self": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME"
        }
    }
}
`, providerDescription, providerID)

	providerListResponse = fmt.Sprintf(`
{
    "identity_providers": [
        {
            "description": "%s",
            "enabled": true,
            "id": "%s",
            "remote_ids": [],
            "links": {
                "protocols": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME/protocols",
                "self": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME"
            }
        },
        {
            "description": "Stores contractor identities",
            "enabled": false,
            "remote_ids": [],
            "id": "ACME-contractors",
            "links": {
                "protocols": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME-contractors/protocols",
                "self": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME-contractors"
            }
        }
    ],
    "links": {
        "next": null,
        "previous": null,
        "self": "https://example.com/v3/OS-FEDERATION/identity_providers"
    }
}
`, providerDescription, providerID)

	updatedProviderResposnse = strings.ReplaceAll(providerResponse, "true", "false")
)
