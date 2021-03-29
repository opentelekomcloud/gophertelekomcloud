package testing

import "fmt"

const (
	providerID          = "ACME"
	providerDescription = "Stores ACME identities"
)

var (
	createURI = fmt.Sprintf("/OS-FEDERATION/identity_providers/%s", providerID)

	jsonResponse = fmt.Sprintf(`
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
)
