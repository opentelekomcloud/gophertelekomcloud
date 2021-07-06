package testing

const (
	protocolID = "saml"
	providerID = "ACME"
	mappingID  = "ACME"

	listURL     = "/OS-FEDERATION/identity_providers/ACME/protocols"
	protocolURL = "/OS-FEDERATION/identity_providers/ACME/protocols/saml"

	listResponseBody = `
{
    "links": {
        "next": null,
        "previous": null,
        "self": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME/protocols"
    },
    "protocols": [
        {
            "id": "saml",
            "links": {
                "identity_provider": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME",
                "self": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME/protocols/saml"
            },
            "mapping_id": "ACME"
        }
    ]
}
`

	getResponseBody = `
{
    "protocol": {
        "id": "saml",
        "links": {
            "identity_provider": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME",
            "self": "https://example.com/v3/OS-FEDERATION/identity_providers/ACME/protocols/saml"
        },
        "mapping_id": "ACME"
    }
}
`

	createRequestBody = `
{
  "protocol": {
    "mapping_id": "ACME"
  }
}
`
)
